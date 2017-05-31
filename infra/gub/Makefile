PROJECT ?= isito-prow
VERSION = 0.2

dev:
	echo "Running on :9001, bucket precise-firmament-3016, using ~/Downloads/myprojectkey.json"
	go run main.go --addr :9001 --bucket precise-firmament-3016 --safile ~/Downloads/myprojectkey.json

image:
	CGO_ENABLED=0 go build -o gub ./
	docker build -t "gcr.io/$(PROJECT)/gub:$(VERSION)" ./
	gcloud docker -- push "gcr.io/$(PROJECT)/gub:$(VERSION)"
