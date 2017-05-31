package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Assumptions:
//  - --addr is set
//  - --bucket is set
//  - --safile is set OR /etc/service-account/service-account.json EXISTS

var (
	// Note: we will bind to this locally, but you must set up an ingress.
	addr = flag.String("addr", "", "the address to listen to")

	// The bucket where jobs are being uploaded to.
	bucket = flag.String("bucket", "", "the bucket where the results are stored")

	// Don't mess with this unless you are running locally, and want to point to a file somewhere.
	safile = flag.String("safile", "/etc/service-account/service-account.json", "path of service account file to use")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	if *addr == "" {
		log.Fatal("--addr must be set")
	}

	if *bucket == "" {
		log.Fatal("--bucket must be set")
	}

	c, err := storage.NewClient(ctx, option.WithServiceAccountFile(*safile))
	if err != nil {
		log.Fatal(err)
	}

	s := server{b: c.Bucket(*bucket)}

	if err := http.ListenAndServe(*addr, http.HandlerFunc(s.handle)); err != nil {
		log.Fatal(err)
	}
}

type server struct {
	b *storage.BucketHandle
}

func isdir(a *storage.ObjectAttrs) bool {
	return a.Name == ""
}

func link(a *storage.ObjectAttrs) string {
	if isdir(a) {
		n := filepath.Base(a.Prefix)
		return fmt.Sprintf(`<a href="%s/">%s/</a>`, n, n)
	}

	n := filepath.Base(a.Name)
	return fmt.Sprintf(`<a href="%s">%s</a>`, n, n)
}

func (s server) handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("handle: %s", r.URL)
	ctx := r.Context()

	path := strings.TrimPrefix(r.URL.String(), "/gub/")

	iter := s.b.Objects(ctx, &storage.Query{
		Delimiter: "/",
		Prefix:    path,
	})

	listing := make([]*storage.ObjectAttrs, 0, 1)

	for {
		attrs, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}

			http.Error(w, fmt.Sprintf("iter.Next() err: %v", err), http.StatusInternalServerError)
			return
		}

		listing = append(listing, attrs)
	}

	switch len(listing) {
	case 0:
		fmt.Fprintf(w, "Path: %s\n", path)
		fmt.Fprintf(w, "No files here ! :(")
	case 1:
		if f := listing[0]; !isdir(f) {
			r, err := s.b.Object(f.Name).NewReader(ctx)
			if err != nil {
				http.Error(w, fmt.Sprintf("s.b.Object(%s).NewReader(ctx)", f.Name), http.StatusInternalServerError)
				return
			}
			defer r.Close()
			io.Copy(w, r)
			return
		}

		fallthrough
	default:
		writedirs(w, path, listing)
	}
}

func writedirs(w io.Writer, path string, listing []*storage.ObjectAttrs) {
	fmt.Fprint(w, "<html><body>")
	fmt.Fprintf(w, "<b>Path: %s</b><br>", path)
	if path != "" {
		fmt.Fprint(w, `<a href="..">..</a><br>`)
	}
	for _, a := range listing {
		fmt.Fprintf(w, "%s<br>", link(a))
	}
	fmt.Fprint(w, "</html></body>")
}
