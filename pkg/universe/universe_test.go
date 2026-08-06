// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package universe

import (
	"os"
	"testing"
)

func TestGetUniverseDomain(t *testing.T) {
	// Test default
	os.Unsetenv(EnvUniverseDomain)
	if got := GetUniverseDomain(); got != "googleapis.com" {
		t.Errorf("GetUniverseDomain() = %v, want googleapis.com", got)
	}

	// Test S3NS France override
	os.Setenv(EnvUniverseDomain, "s3nsapis.fr")
	if got := GetUniverseDomain(); got != "s3nsapis.fr" {
		t.Errorf("GetUniverseDomain() = %v, want s3nsapis.fr", got)
	}

	// Test German TPC override
	os.Setenv(EnvUniverseDomain, "tpc.de")
	if got := GetUniverseDomain(); got != "tpc.de" {
		t.Errorf("GetUniverseDomain() = %v, want tpc.de", got)
	}
	os.Unsetenv(EnvUniverseDomain)
}

func TestGetWorkloadIdentityDomain(t *testing.T) {
	// Test default
	os.Unsetenv(EnvWorkloadIdentityDomain)
	if got := GetWorkloadIdentityDomain(); got != "svc.id.goog" {
		t.Errorf("GetWorkloadIdentityDomain() = %v, want svc.id.goog", got)
	}

	// Test S3NS France override
	os.Setenv(EnvWorkloadIdentityDomain, "s3ns.svc.id.goog")
	if got := GetWorkloadIdentityDomain(); got != "s3ns.svc.id.goog" {
		t.Errorf("GetWorkloadIdentityDomain() = %v, want s3ns.svc.id.goog", got)
	}
	os.Unsetenv(EnvWorkloadIdentityDomain)
}

func TestFormatWorkloadIdentityPrincipal(t *testing.T) {
	os.Setenv(EnvWorkloadIdentityDomain, "s3ns.svc.id.goog")
	defer os.Unsetenv(EnvWorkloadIdentityDomain)

	expected := "principal://iam.googleapis.com/projects/123456/locations/global/workloadIdentityPools/my-project.s3ns.svc.id.goog/subject/ns/kcc-system/sa/cnrm-controller"
	got := FormatWorkloadIdentityPrincipal("123456", "my-project", "kcc-system", "cnrm-controller")
	if got != expected {
		t.Errorf("FormatWorkloadIdentityPrincipal() = %v, want %v", got, expected)
	}
}
