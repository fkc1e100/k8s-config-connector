// Copyright 2025 Google LLC
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

// +tool:fuzz-gen
// proto.message: google.cloud.compute.v1.UrlMap
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeURLMapFuzzer())
}

func computeURLMapFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.UrlMap{},
		ComputeURLMapSpec_v1beta1_FromProto, ComputeURLMapSpec_v1beta1_ToProto,
		ComputeURLMapStatus_v1beta1_FromProto, ComputeURLMapStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecFields.Insert(".default_route_action")
	f.SpecFields.Insert(".default_service")
	f.SpecFields.Insert(".default_url_redirect")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".header_action")
	f.SpecFields.Insert(".host_rules")
	f.SpecFields.Insert(".path_matchers")
	f.SpecFields.Insert(".name")
	f.SpecFields.Insert(".tests")

	// Status fields
	f.StatusFields.Insert(".creation_timestamp")
	f.StatusFields.Insert(".fingerprint")
	f.StatusFields.Insert(".id")
	f.StatusFields.Insert(".self_link")

	// Unimplemented fields
	f.UnimplementedFields.Insert(".kind")
	f.UnimplementedFields.Insert(".region")
	f.UnimplementedFields.Insert(".default_custom_error_response_policy")

	return f
}
