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

type ReplicaKeyInitParameters struct {

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the Default Key Policy section in the AWS Key Management Service Developer Guide.
	// The default value is false.
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check,omitempty"`

	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between 7 and 30, inclusive. If you do not specify a value, it defaults to 30.
	DeletionWindowInDays *int64 `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`

	// A description of the KMS key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. The default value is true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the default key policy to the KMS key.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ReplicaKeyObservation struct {

	// The Amazon Resource Name (ARN) of the replica key. The key ARNs of related multi-Region keys differ only in the Region value.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the Default Key Policy section in the AWS Key Management Service Developer Guide.
	// The default value is false.
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check,omitempty"`

	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between 7 and 30, inclusive. If you do not specify a value, it defaults to 30.
	DeletionWindowInDays *int64 `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`

	// A description of the KMS key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. The default value is true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The key ID of the replica key. Related multi-Region keys have the same key ID.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// A Boolean value that specifies whether key rotation is enabled. This is a shared property of multi-Region keys.
	KeyRotationEnabled *bool `json:"keyRotationEnabled,omitempty" tf:"key_rotation_enabled,omitempty"`

	// The type of key material in the KMS key. This is a shared property of multi-Region keys.
	KeySpec *string `json:"keySpec,omitempty" tf:"key_spec,omitempty"`

	// The cryptographic operations for which you can use the KMS key. This is a shared property of multi-Region keys.
	KeyUsage *string `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`

	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the default key policy to the KMS key.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn *string `json:"primaryKeyArn,omitempty" tf:"primary_key_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ReplicaKeyParameters struct {

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the Default Key Policy section in the AWS Key Management Service Developer Guide.
	// The default value is false.
	// +kubebuilder:validation:Optional
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check,omitempty"`

	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between 7 and 30, inclusive. If you do not specify a value, it defaults to 30.
	// +kubebuilder:validation:Optional
	DeletionWindowInDays *int64 `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`

	// A description of the KMS key.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. The default value is true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the default key policy to the KMS key.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	// +crossplane:generate:reference:type=Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	PrimaryKeyArn *string `json:"primaryKeyArn,omitempty" tf:"primary_key_arn,omitempty"`

	// Reference to a Key to populate primaryKeyArn.
	// +kubebuilder:validation:Optional
	PrimaryKeyArnRef *v1.Reference `json:"primaryKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key to populate primaryKeyArn.
	// +kubebuilder:validation:Optional
	PrimaryKeyArnSelector *v1.Selector `json:"primaryKeyArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ReplicaKeySpec defines the desired state of ReplicaKey
type ReplicaKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicaKeyParameters `json:"forProvider"`
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
	InitProvider ReplicaKeyInitParameters `json:"initProvider,omitempty"`
}

// ReplicaKeyStatus defines the observed state of ReplicaKey.
type ReplicaKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicaKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicaKey is the Schema for the ReplicaKeys API. Manages a KMS multi-Region replica key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReplicaKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicaKeySpec   `json:"spec"`
	Status            ReplicaKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicaKeyList contains a list of ReplicaKeys
type ReplicaKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicaKey `json:"items"`
}

// Repository type metadata.
var (
	ReplicaKey_Kind             = "ReplicaKey"
	ReplicaKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReplicaKey_Kind}.String()
	ReplicaKey_KindAPIVersion   = ReplicaKey_Kind + "." + CRDGroupVersion.String()
	ReplicaKey_GroupVersionKind = CRDGroupVersion.WithKind(ReplicaKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ReplicaKey{}, &ReplicaKeyList{})
}
