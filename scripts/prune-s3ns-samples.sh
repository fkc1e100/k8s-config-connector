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

SAMPLES_DIR="config/samples/resources"

S3NS_SAMPLE_ALLOWLIST=(
  "containercluster"
  "containernodepool"
  "computeinstance"
  "computedisk"
  "computenetwork"
  "computesubnetwork"
  "computefirewall"
  "computeaddress"
  "computerouter"
  "computerouternat"
  "computeforwardingrule"
  "computebackendservice"
  "computeserviceattachment"
  "storagebucket"
  "sqlinstance"
  "sqldatabase"
  "sqluser"
  "bigquerydataset"
  "bigquerytable"
  "project"
  "folder"
  "iampolicy"
  "iampartialpolicy"
  "iampolicymember"
  "iamserviceaccount"
  "secretmanagersecret"
  "secretmanagersecretversion"
  "vertexaifeaturestore"
  "aiplatformmodel"
)

echo "Pruning config/samples/resources to only retain S3NS supported samples..."
count_removed=0
count_retained=0

for dir in ${SAMPLES_DIR}/*; do
  if [ -d "${dir}" ]; then
    dirname=$(basename "${dir}")
    retained=0
    for allow in "${S3NS_SAMPLE_ALLOWLIST[@]}"; do
      if [ "${dirname}" = "${allow}" ]; then
        retained=1
        break
      fi
    done

    if [ "${retained}" -eq 0 ]; then
      rm -rf "${dir}"
      count_removed=$((count_removed + 1))
    else
      count_retained=$((count_retained + 1))
    fi
  fi
done

echo "Done! Removed ${count_removed} unsupported sample directories. Retained ${count_retained} S3NS sample directories."
