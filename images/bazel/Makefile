PROJECT = isito-prow
VERSION = 0.24

image:
	docker build -t "gcr.io/$(PROJECT)/bazel:$(VERSION)" .
	gcloud docker -- push "gcr.io/$(PROJECT)/bazel:$(VERSION)"

.PHONY: image
