# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# These are the usual GKE variables.
PROJECT ?= isito-prow
ZONE ?= us-central1-b
CLUSTER ?= prow

get-cluster-credentials:
	gcloud container clusters get-credentials "$(CLUSTER)" --project="$(PROJECT)" --zone="$(ZONE)"

update-config: get-cluster-credentials
	kubectl create configmap config --from-file=config=./cluster/configmaps/config.yaml --dry-run -o yaml | kubectl replace configmap config -f -

update-plugins: get-cluster-credentials
	kubectl create configmap plugins --from-file=plugins=./cluster/configmaps/plugins.yaml --dry-run -o yaml | kubectl replace configmap plugins -f -

.PHONY: update-config update-plugins get-cluster-credentials

hook-deployment: get-cluster-credentials
	kubectl apply -f cluster/deployments/hook_deployment.yaml

hook-service: get-cluster-credentials
	kubectl create -f cluster/services/hook_service.yaml

sinker-deployment: get-cluster-credentials
	kubectl apply -f cluster/deployments/sinker_deployment.yaml

deck-deployment: get-cluster-credentials
	kubectl apply -f cluster/deployments/deck_deployment.yaml

deck-service: get-cluster-credentials
	kubectl create -f cluster/services/deck_service.yaml

splice-deployment: get-cluster-credentials
	kubectl apply -f cluster/deployments/splice_deployment.yaml

tot-deployment: get-cluster-credentials
	kubectl apply -f cluster/deployments/tot_deployment.yaml

tot-service: get-cluster-credentials
	kubectl apply -f cluster/services/tot_service.yaml

crier-deployment: get-cluster-credentials
	kubectl apply -f cluster/deployment/crier_deployment.yaml

crier-service: get-cluster-credentials
	kubectl apply -f cluster/services/crier_service.yaml

horologium-deployment: get-cluster-credentials
	kubectl apply -f cluster/deployments/horologium_deployment.yaml

plank-deployment: get-cluster-credentials
	kubectl apply -f cluster/deployments/plank_deployment.yaml

.PHONY: hook-deployment hook-service sinker-deployment deck-deployment deck-service splice-deployment tot-service tot-deployment crier-service crier-deployment horologium-deployment plank-deployment
