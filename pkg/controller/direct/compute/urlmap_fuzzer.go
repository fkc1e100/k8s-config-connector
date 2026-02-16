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
	f.SpecField(".default_route_action")
	f.SpecField(".default_service")
	f.SpecField(".default_url_redirect")
	f.SpecField(".description")
	f.SpecField(".header_action")
	f.SpecField(".host_rules")
	f.SpecField(".path_matchers")
	f.SpecField(".name")
	f.SpecField(".tests")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".fingerprint")
	f.StatusField(".id")
	f.StatusField(".self_link")

	// Unimplemented fields
	f.Unimplemented_Internal(".kind")
	f.Unimplemented_Identity(".region")
	f.Unimplemented_NotYetTriaged(".default_custom_error_response_policy")
	f.Unimplemented_NotYetTriaged(".default_route_action.max_stream_duration")
	f.Unimplemented_NotYetTriaged(".default_route_action.request_mirror_policy.mirror_percent")
	f.Unimplemented_NotYetTriaged(".default_route_action.weighted_backend_services[].header_action.request_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".default_route_action.weighted_backend_services[].header_action.response_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".default_url_redirect.strip_query")
	f.Unimplemented_NotYetTriaged(".header_action.request_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".header_action.response_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".path_matchers[].default_custom_error_response_policy")
	f.Unimplemented_NotYetTriaged(".path_matchers[].default_route_action.max_stream_duration")
	f.Unimplemented_NotYetTriaged(".path_matchers[].default_route_action.request_mirror_policy.mirror_percent")
	f.Unimplemented_NotYetTriaged(".path_matchers[].default_route_action.weighted_backend_services[].header_action.request_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".path_matchers[].default_route_action.weighted_backend_services[].header_action.response_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".path_matchers[].default_url_redirect.strip_query")
	f.Unimplemented_NotYetTriaged(".path_matchers[].header_action.request_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".path_matchers[].header_action.response_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".path_matchers[].path_rules[].custom_error_response_policy")
	f.Unimplemented_NotYetTriaged(".path_matchers[].path_rules[].route_action.max_stream_duration")
	f.Unimplemented_NotYetTriaged(".path_matchers[].path_rules[].route_action.request_mirror_policy.mirror_percent")
	f.Unimplemented_NotYetTriaged(".path_matchers[].path_rules[].route_action.weighted_backend_services[].header_action.request_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".path_matchers[].path_rules[].route_action.weighted_backend_services[].header_action.response_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".path_matchers[].path_rules[].url_redirect.strip_query")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].custom_error_response_policy")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].description")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].header_action.request_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].header_action.response_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].route_action.max_stream_duration")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].route_action.request_mirror_policy.mirror_percent")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].route_action.weighted_backend_services[].header_action.request_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].route_action.weighted_backend_services[].header_action.response_headers_to_add[].replace")
	f.Unimplemented_NotYetTriaged(".path_matchers[].route_rules[].url_redirect.strip_query")
	f.Unimplemented_NotYetTriaged(".tests[].expected_output_url")
	f.Unimplemented_NotYetTriaged(".tests[].expected_redirect_response_code")
	f.Unimplemented_NotYetTriaged(".tests[].headers")
	f.Unimplemented_NotYetTriaged(".tests[].headers[].name")
	f.Unimplemented_NotYetTriaged(".tests[].headers[].value")

	return f
}
