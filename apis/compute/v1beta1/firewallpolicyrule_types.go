// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ComputeFirewallPolicyRuleGVK = GroupVersion.WithKind("ComputeFirewallPolicyRule")
)

// +kcc:proto=google.cloud.compute.v1.FirewallPolicyRuleMatcherLayer4Config
type FirewallPolicyRuleMatcherLayer4Config struct {

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcherLayer4Config.ip_protocol
	IPProtocol string `json:"ipProtocol"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcherLayer4Config.ports
	Ports []string `json:"ports,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.FirewallPolicyRuleMatcher
type FirewallPolicyRuleMatcher struct {

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.dest_address_groups
	DestAddressGroups []string `json:"destAddressGroups,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.dest_fqdns
	DestFqdns []string `json:"destFqdns,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.dest_ip_ranges
	DestIPRanges []string `json:"destIPRanges,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.dest_region_codes
	DestRegionCodes []string `json:"destRegionCodes,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.dest_threat_intelligences
	DestThreatIntelligences []string `json:"destThreatIntelligences,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.layer4_configs
	Layer4Configs []FirewallPolicyRuleMatcherLayer4Config `json:"layer4Configs"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_address_groups
	SrcAddressGroups []string `json:"srcAddressGroups,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_fqdns
	SrcFqdns []string `json:"srcFqdns,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_ip_ranges
	SrcIPRanges []string `json:"srcIPRanges,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_region_codes
	SrcRegionCodes []string `json:"srcRegionCodes,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleMatcher.src_threat_intelligences
	SrcThreatIntelligences []string `json:"srcThreatIntelligences,omitempty"`
}

// +kcc:spec:proto=google.cloud.compute.v1.FirewallPolicyRule
type ComputeFirewallPolicyRuleSpec struct {

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.action
	Action string `json:"action"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.direction
	Direction string `json:"direction"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.enable_logging
	EnableLogging *bool `json:"enableLogging,omitempty"`

	FirewallPolicyRef *refs.ComputeFirewallPolicyRef `json:"firewallPolicyRef"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.match
	Match *FirewallPolicyRuleMatcher `json:"match"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.priority
	Priority int64 `json:"priority"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.target_resources
	TargetResources []*ComputeNetworkRef `json:"targetResources,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.target_service_accounts
	TargetServiceAccounts []*refs.IAMServiceAccountRef `json:"targetServiceAccounts,omitempty"`
}

// +kcc:status:proto=google.cloud.compute.v1.FirewallPolicyRule
type ComputeFirewallPolicyRuleStatus struct {
	commonv1alpha1.CommonStatus `json:",inline"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.kind
	Kind *string `json:"kind,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRule.rule_tuple_count
	RuleTupleCount *int64 `json:"ruleTupleCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputefirewallpolicyrule;gcpcomputefirewallpolicyrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeFirewallPolicyRule is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeFirewallPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeFirewallPolicyRuleSpec   `json:"spec,omitempty"`
	Status ComputeFirewallPolicyRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeFirewallPolicyRuleList contains a list of ComputeFirewallPolicyRule
type ComputeFirewallPolicyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeFirewallPolicyRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeFirewallPolicyRule{}, &ComputeFirewallPolicyRuleList{})
}
