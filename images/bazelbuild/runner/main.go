package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// can override in tests
var actualLog = log.Printf

// to be uploaded
var buildLog bytes.Buffer

func blogf(s string, v ...interface{}) {
	fmt.Fprintf(&buildLog, s+"\n", v...)
	actualLog(s, v...)
}

func sh(name string, arg ...string) (output []byte, error) {
	blogf("%s %v", name, arg)
	cmd := exec.Command(name, arg...)
	output, err = cmd.CombinedOutput()
	blogf(string(stdoutStderr))
	if err != nil {
		blogf("exec.Command(%s, %v).CombinedOutput error: %v", name, arg, err)
	}
}

// Assumptions:
//  - --bucket is set

var (
	// The bucket where jobs are being uploaded to.
	bucket = flag.String("bucket", "", "the bucket where the results are stored")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	if *bucket == "" {
		log.Fatal("--bucket must be set")
	}

	c, err := storage.NewClient(ctx, option.WithServiceAccountFile(*safile))
	if err != nil {
		log.Fatal(err)
	}

	e, err := getenv();

	if err == nil {
		if err := run(ctx, e, c.Bucket(*bucket)); err != nil {
			blogf("run err: %v", err)
			log.Print(err)
		}
	} else {
		blogf("getenv err: %v", err)
	}

	if err := writeObject(jobPath(e), ; err != nil {
		log.Fatal(err)
	}
}

type context struct {
	err error
}

func (c context) blogf(name string, v ...interface{}) {
	if c.err != nil {
		return
	}

	blogf(name, arg...)
}

func (c context) sh(name string, arg ...string) (out []byte) {
	if c.err != nil {
		return
	}

	out, c.err = sh(name, arg...)
}

func jobPath(e *env) string {
	return fmt.Sprintf("jobs/%s/%s/%s", e.RepoOwner, e.RepoName, e.JobName)
}

func writeObject(b *storage.BucketHandle, path string, in io.Reader) (err error) {
	w, err := b.Object(path).NewWriter()
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, in)
}

// assumes a valid env
func run(ctx context.Context, e *env, c storage.BucketHandle) error {
	c := context{}
	s := gub.JobStatus{}

	c.blogf("> Cloning: %s/%s", e.RepoOwner, e.RepoName)
	c.sh("git", "clone", fmt.Sprintf("https://github.com/%s/%s.git", e.RepoOwner, e.RepoName))
	c.sh("cd", e.RepoName)
	c.blogf("git fetch origin pull/%s/head:prow", e.PullNumber)
	c.sh("git", "checkout", "prow")
	c.sh("mkdir", "/artifacts")
	out := c.sh(fmt.Sprintf("./jobs/%s.sh", e.JobName));
	if err := writeObject(b, jobPath(e), bytes.NewReader(out)); err != nil {
		return fmt.Errorf("writeObject error: %v", err)
	}



}
