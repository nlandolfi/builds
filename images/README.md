This directory contains standard images maintained by Istio.

* bazelbuild - gives an environment with git, gcloud, python, node, and bazel TODO: has go?

Each subdirectory has a Makefile with `image` and `push` commands which should be used to build and push the image to the istio gcloud project.
