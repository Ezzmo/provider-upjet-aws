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

type ClusterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CapacityProviders []*string `json:"capacityProviders,omitempty" tf:"capacity_providers,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultCapacityProviderStrategy []DefaultCapacityProviderStrategyParameters `json:"defaultCapacityProviderStrategy,omitempty" tf:"default_capacity_provider_strategy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Setting []SettingParameters `json:"setting,omitempty" tf:"setting,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConfigurationObservation struct {
}

type ConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ExecuteCommandConfiguration []ExecuteCommandConfigurationParameters `json:"executeCommandConfiguration,omitempty" tf:"execute_command_configuration,omitempty"`
}

type DefaultCapacityProviderStrategyObservation struct {
}

type DefaultCapacityProviderStrategyParameters struct {

	// +kubebuilder:validation:Optional
	Base *float64 `json:"base,omitempty" tf:"base,omitempty"`

	// +kubebuilder:validation:Required
	CapacityProvider *string `json:"capacityProvider" tf:"capacity_provider,omitempty"`

	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type ExecuteCommandConfigurationObservation struct {
}

type ExecuteCommandConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	LogConfiguration []LogConfigurationParameters `json:"logConfiguration,omitempty" tf:"log_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Logging *string `json:"logging,omitempty" tf:"logging,omitempty"`
}

type LogConfigurationObservation struct {
}

type LogConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CloudWatchEncryptionEnabled *bool `json:"cloudWatchEncryptionEnabled,omitempty" tf:"cloud_watch_encryption_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	CloudWatchLogGroupName *string `json:"cloudWatchLogGroupName,omitempty" tf:"cloud_watch_log_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	S3BucketEncryptionEnabled *bool `json:"s3BucketEncryptionEnabled,omitempty" tf:"s3_bucket_encryption_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// +kubebuilder:validation:Optional
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type SettingObservation struct {
}

type SettingParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
