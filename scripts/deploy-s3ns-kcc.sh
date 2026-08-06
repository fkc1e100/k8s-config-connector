#!/usr/bin/env bash
# Copyright 2026 S3NS / Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

UNIVERSE_DOMAIN="${GOOGLE_CLOUD_UNIVERSE_DOMAIN:-s3nsapis.fr}"
WI_DOMAIN="${WORKLOAD_IDENTITY_DOMAIN:-s3ns.svc.id.goog}"
NAMESPACE="kcc-system"

echo "===================================================="
echo "  Deploying Slimmed S3NS KCC Bundle"
echo "  Universe Domain: ${UNIVERSE_DOMAIN}"
echo "  Workload Identity Domain: ${WI_DOMAIN}"
echo "===================================================="

# 1. Apply Pruned CRD Bundle (40 CRDs)
echo "[1/3] Applying pruned S3NS CRD manifests..."
kubectl apply -f config/crds/resources/

# 2. Create kcc-system Namespace if not exists
echo "[2/3] Preparing ${NAMESPACE} namespace..."
kubectl create namespace "${NAMESPACE}" --dry-run=client -o yaml | kubectl apply -f -

# 3. Deploy Controller Manager with TPC Environment Variables
echo "[3/3] Deploying cnrm-controller-manager..."
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cnrm-controller-manager
  namespace: ${NAMESPACE}
  labels:
    cnrm.cloud.google.com/system: "true"
spec:
  replicas: 1
  selector:
    matchLabels:
      cnrm.cloud.google.com/system: "true"
  template:
    metadata:
      labels:
        cnrm.cloud.google.com/system: "true"
    spec:
      containers:
      - name: manager
        image: gcr.io/gke-release/cnrm/controller:v1.154.1
        env:
        - name: GOOGLE_CLOUD_UNIVERSE_DOMAIN
          value: "${UNIVERSE_DOMAIN}"
        - name: WORKLOAD_IDENTITY_DOMAIN
          value: "${WI_DOMAIN}"
        resources:
          limits:
            cpu: 500m
            memory: 512Mi
          requests:
            cpu: 100m
            memory: 128Mi
EOF

echo "===================================================="
echo "  S3NS KCC deployment complete!"
echo "  Verify health: kubectl get pods -n ${NAMESPACE}"
echo "===================================================="
