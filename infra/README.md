Infra
-----

This directory contains a Makefile and other resources for managing the Istio CI infrastructure. this infrastructure consists of a variety of services, largely based off the prow work by k8s test infra.

### Managing a Cluster

The infrastructure runs, of course, on a k8s cluster. All of our tools are set up to handle running on GCE. The variables for project/zone/cluster are in the Makefile. The Makefile also contains commands for almost everything you will need to do to manage the infrastructure.
