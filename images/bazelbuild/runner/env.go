package main

import (
	"fmt"
	"os"
)

const (
	repoOwner  = "REPO_OWNER"
	repoName   = "REPO_NAME"
	pullNumber = "PULL_NUMBER"
	jobName    = "JOB_NAME"
)

type env struct {
	RepoOwner  string
	RepoName   string
	PullNumber string
	JobName    string
}

func noVarErr(name string) error {
	return fmt.Errorf("missing environment variable: %q", name)
}

func getenv() (e env, err error) {
	if s := os.Getenv(repoOwner); s == "" {
		return e, noVarErr(repoOwner)
	} else {
		e.RepoOwner = s
	}

	if s := os.Getenv(repoName); s == "" {
		return e, noVarErr(repoName)
	} else {
		e.RepoName = s
	}

	if s := os.Getenv(pullNumber); s == "" {
		return e, noVarErr(pullNumber)
	} else {
		e.PullNumber = s
	}

	if s := os.Getenv(jobName); s == "" {
		return e, noVarErr(jobName)
	} else {
		e.JobName = s
	}

	return e, nil
}
