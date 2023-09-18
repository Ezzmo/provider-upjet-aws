/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DomainInitParameters struct {

	// Domain endpoint options. Documented below.
	EndpointOptions []EndpointOptionsInitParameters `json:"endpointOptions,omitempty" tf:"endpoint_options,omitempty"`

	// The index fields for documents added to the domain. Documented below.
	IndexField []IndexFieldInitParameters `json:"indexField,omitempty" tf:"index_field,omitempty"`

	// Whether or not to maintain extra instances for the domain in a second Availability Zone to ensure high availability.
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// Domain scaling parameters. Documented below.
	ScalingParameters []ScalingParametersInitParameters `json:"scalingParameters,omitempty" tf:"scaling_parameters,omitempty"`
}

type DomainObservation struct {

	// The domain's ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The service endpoint for updating documents in a search domain.
	DocumentServiceEndpoint *string `json:"documentServiceEndpoint,omitempty" tf:"document_service_endpoint,omitempty"`

	// An internally generated unique identifier for the domain.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// Domain endpoint options. Documented below.
	EndpointOptions []EndpointOptionsObservation `json:"endpointOptions,omitempty" tf:"endpoint_options,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The index fields for documents added to the domain. Documented below.
	IndexField []IndexFieldObservation `json:"indexField,omitempty" tf:"index_field,omitempty"`

	// Whether or not to maintain extra instances for the domain in a second Availability Zone to ensure high availability.
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// Domain scaling parameters. Documented below.
	ScalingParameters []ScalingParametersObservation `json:"scalingParameters,omitempty" tf:"scaling_parameters,omitempty"`

	// The service endpoint for requesting search results from a search domain.
	SearchServiceEndpoint *string `json:"searchServiceEndpoint,omitempty" tf:"search_service_endpoint,omitempty"`
}

type DomainParameters struct {

	// Domain endpoint options. Documented below.
	// +kubebuilder:validation:Optional
	EndpointOptions []EndpointOptionsParameters `json:"endpointOptions,omitempty" tf:"endpoint_options,omitempty"`

	// The index fields for documents added to the domain. Documented below.
	// +kubebuilder:validation:Optional
	IndexField []IndexFieldParameters `json:"indexField,omitempty" tf:"index_field,omitempty"`

	// Whether or not to maintain extra instances for the domain in a second Availability Zone to ensure high availability.
	// +kubebuilder:validation:Optional
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Domain scaling parameters. Documented below.
	// +kubebuilder:validation:Optional
	ScalingParameters []ScalingParametersParameters `json:"scalingParameters,omitempty" tf:"scaling_parameters,omitempty"`
}

type EndpointOptionsInitParameters struct {

	// Enables or disables the requirement that all requests to the domain arrive over HTTPS.
	EnforceHTTPS *bool `json:"enforceHttps,omitempty" tf:"enforce_https,omitempty"`

	// The minimum required TLS version. See the AWS documentation for valid values.
	TLSSecurityPolicy *string `json:"tlsSecurityPolicy,omitempty" tf:"tls_security_policy,omitempty"`
}

type EndpointOptionsObservation struct {

	// Enables or disables the requirement that all requests to the domain arrive over HTTPS.
	EnforceHTTPS *bool `json:"enforceHttps,omitempty" tf:"enforce_https,omitempty"`

	// The minimum required TLS version. See the AWS documentation for valid values.
	TLSSecurityPolicy *string `json:"tlsSecurityPolicy,omitempty" tf:"tls_security_policy,omitempty"`
}

type EndpointOptionsParameters struct {

	// Enables or disables the requirement that all requests to the domain arrive over HTTPS.
	// +kubebuilder:validation:Optional
	EnforceHTTPS *bool `json:"enforceHttps,omitempty" tf:"enforce_https,omitempty"`

	// The minimum required TLS version. See the AWS documentation for valid values.
	// +kubebuilder:validation:Optional
	TLSSecurityPolicy *string `json:"tlsSecurityPolicy,omitempty" tf:"tls_security_policy,omitempty"`
}

type IndexFieldInitParameters struct {

	// The analysis scheme you want to use for a text field. The analysis scheme specifies the language-specific text processing options that are used during indexing.
	AnalysisScheme *string `json:"analysisScheme,omitempty" tf:"analysis_scheme,omitempty"`

	// The default value for the field. This value is used when no value is specified for the field in the document data.
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// You can get facet information by enabling this.
	Facet *bool `json:"facet,omitempty" tf:"facet,omitempty"`

	// You can highlight information.
	Highlight *bool `json:"highlight,omitempty" tf:"highlight,omitempty"`

	// The name of the CloudSearch domain.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// You can enable returning the value of all searchable fields.
	Return *bool `json:"return,omitempty" tf:"return,omitempty"`

	// You can set whether this index should be searchable or not.
	Search *bool `json:"search,omitempty" tf:"search,omitempty"`

	// You can enable the property to be sortable.
	Sort *bool `json:"sort,omitempty" tf:"sort,omitempty"`

	// A comma-separated list of source fields to map to the field. Specifying a source field copies data from one field to another, enabling you to use the same source data in different ways by configuring different options for the fields.
	SourceFields *string `json:"sourceFields,omitempty" tf:"source_fields,omitempty"`

	// The field type. Valid values: date, date-array, double, double-array, int, int-array, literal, literal-array, text, text-array.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IndexFieldObservation struct {

	// The analysis scheme you want to use for a text field. The analysis scheme specifies the language-specific text processing options that are used during indexing.
	AnalysisScheme *string `json:"analysisScheme,omitempty" tf:"analysis_scheme,omitempty"`

	// The default value for the field. This value is used when no value is specified for the field in the document data.
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// You can get facet information by enabling this.
	Facet *bool `json:"facet,omitempty" tf:"facet,omitempty"`

	// You can highlight information.
	Highlight *bool `json:"highlight,omitempty" tf:"highlight,omitempty"`

	// The name of the CloudSearch domain.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// You can enable returning the value of all searchable fields.
	Return *bool `json:"return,omitempty" tf:"return,omitempty"`

	// You can set whether this index should be searchable or not.
	Search *bool `json:"search,omitempty" tf:"search,omitempty"`

	// You can enable the property to be sortable.
	Sort *bool `json:"sort,omitempty" tf:"sort,omitempty"`

	// A comma-separated list of source fields to map to the field. Specifying a source field copies data from one field to another, enabling you to use the same source data in different ways by configuring different options for the fields.
	SourceFields *string `json:"sourceFields,omitempty" tf:"source_fields,omitempty"`

	// The field type. Valid values: date, date-array, double, double-array, int, int-array, literal, literal-array, text, text-array.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IndexFieldParameters struct {

	// The analysis scheme you want to use for a text field. The analysis scheme specifies the language-specific text processing options that are used during indexing.
	// +kubebuilder:validation:Optional
	AnalysisScheme *string `json:"analysisScheme,omitempty" tf:"analysis_scheme,omitempty"`

	// The default value for the field. This value is used when no value is specified for the field in the document data.
	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// You can get facet information by enabling this.
	// +kubebuilder:validation:Optional
	Facet *bool `json:"facet,omitempty" tf:"facet,omitempty"`

	// You can highlight information.
	// +kubebuilder:validation:Optional
	Highlight *bool `json:"highlight,omitempty" tf:"highlight,omitempty"`

	// The name of the CloudSearch domain.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// You can enable returning the value of all searchable fields.
	// +kubebuilder:validation:Optional
	Return *bool `json:"return,omitempty" tf:"return,omitempty"`

	// You can set whether this index should be searchable or not.
	// +kubebuilder:validation:Optional
	Search *bool `json:"search,omitempty" tf:"search,omitempty"`

	// You can enable the property to be sortable.
	// +kubebuilder:validation:Optional
	Sort *bool `json:"sort,omitempty" tf:"sort,omitempty"`

	// A comma-separated list of source fields to map to the field. Specifying a source field copies data from one field to another, enabling you to use the same source data in different ways by configuring different options for the fields.
	// +kubebuilder:validation:Optional
	SourceFields *string `json:"sourceFields,omitempty" tf:"source_fields,omitempty"`

	// The field type. Valid values: date, date-array, double, double-array, int, int-array, literal, literal-array, text, text-array.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ScalingParametersInitParameters struct {

	// The instance type that you want to preconfigure for your domain. See the AWS documentation for valid values.
	DesiredInstanceType *string `json:"desiredInstanceType,omitempty" tf:"desired_instance_type,omitempty"`

	// The number of partitions you want to preconfigure for your domain. Only valid when you select search.2xlarge as the instance type.
	DesiredPartitionCount *int64 `json:"desiredPartitionCount,omitempty" tf:"desired_partition_count,omitempty"`

	// The number of replicas you want to preconfigure for each index partition.
	DesiredReplicationCount *int64 `json:"desiredReplicationCount,omitempty" tf:"desired_replication_count,omitempty"`
}

type ScalingParametersObservation struct {

	// The instance type that you want to preconfigure for your domain. See the AWS documentation for valid values.
	DesiredInstanceType *string `json:"desiredInstanceType,omitempty" tf:"desired_instance_type,omitempty"`

	// The number of partitions you want to preconfigure for your domain. Only valid when you select search.2xlarge as the instance type.
	DesiredPartitionCount *int64 `json:"desiredPartitionCount,omitempty" tf:"desired_partition_count,omitempty"`

	// The number of replicas you want to preconfigure for each index partition.
	DesiredReplicationCount *int64 `json:"desiredReplicationCount,omitempty" tf:"desired_replication_count,omitempty"`
}

type ScalingParametersParameters struct {

	// The instance type that you want to preconfigure for your domain. See the AWS documentation for valid values.
	// +kubebuilder:validation:Optional
	DesiredInstanceType *string `json:"desiredInstanceType,omitempty" tf:"desired_instance_type,omitempty"`

	// The number of partitions you want to preconfigure for your domain. Only valid when you select search.2xlarge as the instance type.
	// +kubebuilder:validation:Optional
	DesiredPartitionCount *int64 `json:"desiredPartitionCount,omitempty" tf:"desired_partition_count,omitempty"`

	// The number of replicas you want to preconfigure for each index partition.
	// +kubebuilder:validation:Optional
	DesiredReplicationCount *int64 `json:"desiredReplicationCount,omitempty" tf:"desired_replication_count,omitempty"`
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DomainInitParameters `json:"initProvider,omitempty"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Domain is the Schema for the Domains API. Provides an CloudSearch domain resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec"`
	Status            DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	Domain_Kind             = "Domain"
	Domain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Domain_Kind}.String()
	Domain_KindAPIVersion   = Domain_Kind + "." + CRDGroupVersion.String()
	Domain_GroupVersionKind = CRDGroupVersion.WithKind(Domain_Kind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
