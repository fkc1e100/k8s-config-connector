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
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	ComputeForwardingRuleGVK = schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: GroupVersion.Version,
		Kind:    "ComputeForwardingRule",
	}
)

// +kcc:proto=google.cloud.compute.v1.MetadataFilterLabelMatch
type MetadataFilterLabelMatch struct {
	
	// +kcc:proto:field=google.cloud.compute.v1.MetadataFilterLabelMatch.name
	Name *string `json:"name"`

	
	// +kcc:proto:field=google.cloud.compute.v1.MetadataFilterLabelMatch.value
	Value *string `json:"value"`
}

type IpAddress struct {
	// +optional
	AddressRef *refs.ComputeAddressRef `json:"addressRef,omitempty"`

	// +optional
	Ip *string `json:"ip,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.MetadataFilter
type MetadataFilter struct {
	
	// +kcc:proto:field=google.cloud.compute.v1.MetadataFilter.filter_labels
	FilterLabels []MetadataFilterLabelMatch `json:"filterLabels"`

	
	// +kcc:proto:field=google.cloud.compute.v1.MetadataFilter.filter_match_criteria
	FilterMatchCriteria *string `json:"filterMatchCriteria"`
}

// +kcc:proto=google.cloud.compute.v1.ForwardingRuleServiceDirectoryRegistration
type ForwardingruleServiceDirectoryRegistrations struct {
	
	// +optional
	Namespace *string `json:"namespace,omitempty"`

	
	// +optional
	Service *string `json:"service,omitempty"`
}

type Target struct {
	// +optional
	GoogleAPIsBundle *string `json:"googleAPIsBundle,omitempty"`

	// +optional
	ServiceAttachmentRef *refs.ComputeServiceAttachmentRef `json:"serviceAttachmentRef,omitempty"`

	// +optional
	TargetGRPCProxyRef *refs.ComputeTargetGrpcProxyRef `json:"targetGRPCProxyRef,omitempty"`

	// +optional
	TargetHTTPProxyRef *refs.ComputeTargetHTTPProxyRef `json:"targetHTTPProxyRef,omitempty"`

	// +optional
	TargetHTTPSProxyRef *refs.ComputeTargetHTTPSProxyRef `json:"targetHTTPSProxyRef,omitempty"`

	// +optional
	TargetSSLProxyRef *refs.ComputeTargetSSLProxyRef `json:"targetSSLProxyRef,omitempty"`

	// +optional
	TargetTCPProxyRef *refs.ComputeTargetTCPProxyRef `json:"targetTCPProxyRef,omitempty"`

	// +optional
	TargetVPNGatewayRef *refs.ComputeTargetVPNGatewayRef `json:"targetVPNGatewayRef,omitempty"`
}

// +kcc:spec:proto=google.cloud.compute.v1.ForwardingRule
type ComputeForwardingRuleSpec struct {
	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.all_ports
	AllPorts *bool `json:"allPorts,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.allow_global_access
	AllowGlobalAccess *bool `json:"allowGlobalAccess,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.allow_psc_global_access
	AllowPscGlobalAccess *bool `json:"allowPscGlobalAccess,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.backend_service
	BackendServiceRef *ComputeBackendServiceRef `json:"backendServiceRef,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.description
	Description *string `json:"description,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.IPAddress
	IpAddress *IpAddress `json:"ipAddress,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.IPProtocol
	IpProtocol *string `json:"ipProtocol,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.ip_version
	IpVersion *string `json:"ipVersion,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.is_mirroring_collector
	IsMirroringCollector *bool `json:"isMirroringCollector,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.load_balancing_scheme
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	
	Location *string `json:"location"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.metadata_filters
	MetadataFilters []MetadataFilter `json:"metadataFilters,omitempty"`

	
	NetworkRef *ComputeNetworkRef `json:"networkRef,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.network_tier
	NetworkTier *string `json:"networkTier,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.no_automate_dns_zone
	NoAutomateDnsZone *bool `json:"noAutomateDnsZone,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.port_range
	PortRange *string `json:"portRange,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.ports
	Ports []string `json:"ports,omitempty"`

	
	ResourceID *string `json:"resourceID,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.service_directory_registrations
	ServiceDirectoryRegistrations []ForwardingruleServiceDirectoryRegistrations `json:"serviceDirectoryRegistrations,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.service_label
	ServiceLabel *string `json:"serviceLabel,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.source_ip_ranges
	SourceIpRanges []string `json:"sourceIpRanges,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.subnetwork
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.target
	Target *Target `json:"target,omitempty"`
}

// +kcc:status:proto=google.cloud.compute.v1.ForwardingRule
type ComputeForwardingRuleStatus struct {
	commonv1alpha1.CommonStatus `json:",inline"`
	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.base_forwarding_rule
	BaseForwardingRule *string `json:"baseForwardingRule,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.psc_connection_id
	PscConnectionId *string `json:"pscConnectionId,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.psc_connection_status
	PscConnectionStatus *string `json:"pscConnectionStatus,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.service_name
	ServiceName *string `json:"serviceName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeforwardingrule;gcpcomputeforwardingrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeForwardingRule is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeForwardingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeForwardingRuleSpec   `json:"spec,omitempty"`
	Status ComputeForwardingRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeForwardingRuleList contains a list of ComputeForwardingRule
type ComputeForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeForwardingRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeForwardingRule{}, &ComputeForwardingRuleList{})
}
