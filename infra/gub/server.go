package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/storage"
)

type server struct {
	b *storage.BucketHandle
}

func (s server) handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	path := strings.TrimPrefix(r.URL.String(), "/gub/")
	log.Printf("handle: %s -> %s", r.URL, path)

	listing, err := list(ctx, s.b, path)
	if err != nil {
		http.Error(w, fmt.Sprintf("list err: %v", err), http.StatusInternalServerError)
		return
	}

	switch len(listing) {
	case 0:
		fmt.Fprintf(w, "Path: %s\nNo files here ! :(", path)
	case 1:
		var f *storage.ObjectAttrs
		for _, v := range listing {
			f = v
			break
		}

		if !isdir(f) {
			// this is a file (i.e., a leaf in the tree induced by paths)

			r, err := s.b.Object(f.Name).NewReader(ctx)
			if err != nil {
				http.Error(w, fmt.Sprintf("s.b.Object(%s).NewReader(ctx) err: %s", f.Name, err), http.StatusInternalServerError)
				return
			}
			defer r.Close()
			io.Copy(w, r)
		}

		fallthrough
	default:
		if _, ok := listing[".job.json"]; ok {
			if err := s.writeJobDir(ctx, w, path, listing); err != nil {
				http.Error(w, fmt.Sprintf("s.writejobdir err: %s", err), http.StatusInternalServerError)
			}
			return
		}

		s.writedirs(w, path, listing)
	}
}

func (s server) writeJobDir(ctx context.Context, w http.ResponseWriter, path string, listing map[string]*storage.ObjectAttrs) error {
	d := jobTempData{}

	if f, ok := listing["stdout.txt"]; ok {
		html, err := fileAsHTML(ctx, s.b, f)
		if err != nil {
			return fmt.Errorf("fileAsHTML err: %s", err)
		}
		d.Stdout = html
	}

	if f, ok := listing["stderr.txt"]; ok {
		html, err := fileAsHTML(ctx, s.b, f)
		if err != nil {
			return fmt.Errorf("fileAsHTML err: %s", err)
		}
		d.Stderr = html
	}

	if f, ok := listing["artifacts"]; ok {
		log.Printf("artifacts listing: %s", f.Prefix)

		artListing, err := list(ctx, s.b, f.Prefix)
		if err != nil {
			return fmt.Errorf("list artifacts/ err: %s", err)
		}

		for _, a := range artListing {
			d.ArtifactLinks = append(d.ArtifactLinks, template.HTML(link(a)))
		}
	}

	if err := jobDirTemplate.Execute(w, d); err != nil {
		return fmt.Errorf("jobDirTemplate.Execute err: %s", err)
	}

	return nil
}

type jobTempData struct {
	Stdout        template.HTML
	StdoutLines   int
	Stderr        template.HTML
	StderrLines   int
	ArtifactLinks []template.HTML
}

var jobDirTemplate = template.Must(template.New("").Parse(`
<html>
	<body>
		<h3>stdout ({{ printf "%d" .StdoutLines }} Lines)</h3>
		<hr>
		<div id="stdout-body">
			{{ .Stdout }}
		</div>
		<h3>stderr ({{ printf "%d" .StderrLines }} Lines)</h3>
		<hr>
		<div id="stdout-err">
			{{ .Stderr }}
		</div>
		<h3>artifacts</h3>
		<hr>
		{{ range .ArtifactLinks }}
			{{ . }} <br>
		{{ else }}
			No artifacts !
		{{ end }}
	</body>
</html>
`))

func (s server) writedirs(w io.Writer, path string, listing map[string]*storage.ObjectAttrs) error {
	log.Printf("listing, %v", listing)
	d := dirsTempData{
		Path: path,
	}

	dirs := make([]string, 0)
	files := make([]string, 0)

	for n, a := range listing {
		if isdir(a) {
			dirs = append(dirs, a)
			continue
		}

		files = append(files, a)
	}

	for _, n := range dirs {
		d.Links = append(d.Links, template.HTML(link(a)))
	}

	for _, n := range files {
		d.Links = append(d.Links, template.HTML(link(a)))
	}

	if err := dirsTemplate.Execute(w, d); err != nil {
		return fmt.Errorf("dirsTemplate.Execute err: %s", err)
	}

	return nil
}

type dirsTempData struct {
	Path  string
	Links []template.HTML
}

var dirsTemplate = template.Must(template.New("").Parse(`
<html>
	<body>
		<b>Path: {{ printf "%s" .Path }}</b>
		<br>
		<a href="..">..</a>
		<br>
		{{ range .Links }}
			{{ . }} <br>
		{{ end }}
	</body>
</html>
`))
