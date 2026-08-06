#!/usr/bin/env bash
set -euo pipefail

CRD_DIR="config/crds/resources"

# Allowlist of S3NS supported CRD prefixes
ALLOWED_PATTERNS=(
  "containerclusters"
  "containernodepools"
  "computeinstances"
  "computedisks"
  "computenetworks"
  "computesubnetworks"
  "computefirewalls"
  "computeaddresses"
  "computerouters"
  "computerouternats"
  "computeforwardingrules"
  "computebackendservices"
  "storagebuckets"
  "sqlinstances"
  "sqldatabases"
  "sqlusers"
  "bigquerydatasets"
  "bigquerytables"
  "projects"
  "folders"
  "iampolicies"
  "iampartialpolicies"
  "iampolicymembers"
  "iamserviceaccounts"
  "secretmanagersecrets"
  "secretmanagersecretversions"
  "computeserviceattachments"
  "computetargetgrpcproxies"
  "vertexaifeaturestores"
  "aiplatformmodels"
)

echo "Starting CRD pruning in $CRD_DIR..."
initial_count=$(ls -1 "$CRD_DIR"/*.yaml | wc -l)
echo "Initial CRDs: $initial_count"

for file in "$CRD_DIR"/*.yaml; do
  filename=$(basename "$file")
  keep=false
  for pattern in "${ALLOWED_PATTERNS[@]}"; do
    if [[ "$filename" == *"$pattern"* ]]; then
      keep=true
      break
    fi
  done
  if [ "$keep" = false ]; then
    rm -f "$file"
  fi
done

final_count=$(ls -1 "$CRD_DIR"/*.yaml | wc -l)
echo "Pruning complete!"
echo "Remaining S3NS CRDs: $final_count"
