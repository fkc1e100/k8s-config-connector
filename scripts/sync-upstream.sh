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

UPSTREAM_TAG="${1:-}"

if [ -z "${UPSTREAM_TAG}" ]; then
  echo "Usage: ./scripts/sync-upstream.sh <UPSTREAM_TAG> (e.g. v1.155.0)"
  exit 1
fi

echo "===================================================="
echo "  Syncing upstream KCC tag: ${UPSTREAM_TAG}"
echo "===================================================="

# 1. Ensure upstream remote exists
git remote add upstream https://github.com/GoogleCloudPlatform/k8s-config-connector.git 2>/dev/null || true
git fetch upstream --tags

# 2. Re-run S3NS Pruning Tools
echo "[1/3] Running S3NS CRD and sample pruning tools..."
./scripts/prune-s3ns-crds.sh
./scripts/prune-s3ns-samples.sh
./scripts/prune-s3ns-mockgcp.sh

# 3. Run S3NS Unit Test Validation
echo "[2/3] Running S3NS universe unit tests..."
go test -v ./pkg/universe/...

echo "===================================================="
echo "  Upstream sync prep complete for ${UPSTREAM_TAG}!"
echo "===================================================="
