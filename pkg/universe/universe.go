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
	"fmt"
	"os"
	"strings"
)

const (
	DefaultUniverseDomain         = "googleapis.com"
	DefaultWorkloadIdentityDomain = "svc.id.goog"

	EnvUniverseDomain         = "GOOGLE_CLOUD_UNIVERSE_DOMAIN"
	EnvWorkloadIdentityDomain = "WORKLOAD_IDENTITY_DOMAIN"
)

// GetUniverseDomain returns the configured TPC Universe Domain.
// Checks GOOGLE_CLOUD_UNIVERSE_DOMAIN env var, falling back to DefaultUniverseDomain ("googleapis.com").
func GetUniverseDomain() string {
	if domain := os.Getenv(EnvUniverseDomain); domain != "" {
		return strings.TrimSpace(domain)
	}
	return DefaultUniverseDomain
}

// GetWorkloadIdentityDomain returns the configured Workload Identity domain suffix.
// Checks WORKLOAD_IDENTITY_DOMAIN env var, falling back to DefaultWorkloadIdentityDomain ("svc.id.goog").
func GetWorkloadIdentityDomain() string {
	if domain := os.Getenv(EnvWorkloadIdentityDomain); domain != "" {
		return strings.TrimSpace(domain)
	}
	return DefaultWorkloadIdentityDomain
}

// FormatEndpoint returns the service endpoint for a given GCP service prefix (e.g. "storage", "compute").
func FormatEndpoint(servicePrefix string) string {
	domain := GetUniverseDomain()
	return fmt.Sprintf("https://%s.%s", servicePrefix, domain)
}

// FormatWorkloadIdentityPrincipal builds the federated Workload Identity principal URI dynamically.
func FormatWorkloadIdentityPrincipal(projectNumber, projectID, namespace, saName string) string {
	wiDomain := GetWorkloadIdentityDomain()
	return fmt.Sprintf("principal://iam.googleapis.com/projects/%s/locations/global/workloadIdentityPools/%s.%s/subject/ns/%s/sa/%s",
		projectNumber, projectID, wiDomain, namespace, saName)
}

// IsDefaultUniverse returns true if running in standard global GCP (googleapis.com).
func IsDefaultUniverse() bool {
	return GetUniverseDomain() == DefaultUniverseDomain
}
