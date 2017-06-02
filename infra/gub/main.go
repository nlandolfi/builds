package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
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
