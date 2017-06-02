package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// Use isdir to determine if the file represents a directory.
//
// Note: there are no directories in GCS. Consider dir/file.txt
// where the file, "file.txt", is in a dir, "dir". This file
// would have the name: "dir/file.txt".
func isdir(a *storage.ObjectAttrs) bool {
	return a.Name == ""
}

// Use to get the base name for display of a file.
func name(f *storage.ObjectAttrs) string {
	s := f.Name
	if s == "" {
		s = f.Prefix
	}

	return filepath.Base(s)
}

func list(ctx context.Context, b *storage.BucketHandle, path string) (map[string]*storage.ObjectAttrs, error) {
	iter := b.Objects(ctx, &storage.Query{
		Delimiter: "/",
		Prefix:    path,
	})

	listing := make(map[string]*storage.ObjectAttrs)

	for {
		attrs, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}

			return nil, err
		}

		listing[name(attrs)] = attrs
	}

	return listing, nil
}

func link(a *storage.ObjectAttrs) string {
	if isdir(a) {
		n := filepath.Base(a.Prefix)
		return fmt.Sprintf(`<a href="%s/">%s/</a>`, n, n)
	}

	n := filepath.Base(a.Name)
	return fmt.Sprintf(`<a href="%s">%s</a>`, n, n)
}

func fileAsHTML(ctx context.Context, b *storage.BucketHandle, f *storage.ObjectAttrs) (html template.HTML, err error) {
	var r *storage.Reader
	r, err = b.Object(f.Name).NewReader(ctx)
	if err != nil {
		return
	}
	defer r.Close()

	var bs []byte

	bs, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}

	html = template.HTML(bytes.Replace(bs, []byte("\n"), []byte("<br>"), -1))
	return
}
