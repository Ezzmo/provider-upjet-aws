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

type ConnectionPasswordEncryptionObservation struct {
}

type ConnectionPasswordEncryptionParameters struct {

	// A KMS key ARN that is used to encrypt the connection password. If connection password protection is enabled, the caller of CreateConnection and UpdateConnection needs at least kms:Encrypt permission on the specified AWS KMS key, to encrypt passwords before storing them in the Data Catalog.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	AwsKMSKeyID *string `json:"awsKmsKeyId,omitempty" tf:"aws_kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	AwsKMSKeyIDRef *v1.Reference `json:"awsKmsKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AwsKMSKeyIDSelector *v1.Selector `json:"awsKmsKeyIdSelector,omitempty" tf:"-"`

	// When set to true, passwords remain encrypted in the responses of GetConnection and GetConnections. This encryption takes effect independently of the catalog encryption.
	// +kubebuilder:validation:Required
	ReturnConnectionPasswordEncrypted *bool `json:"returnConnectionPasswordEncrypted" tf:"return_connection_password_encrypted,omitempty"`
}

type DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsObservation struct {
}

type DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsParameters struct {

	// When connection password protection is enabled, the Data Catalog uses a customer-provided key to encrypt the password as part of CreateConnection or UpdateConnection and store it in the ENCRYPTED_PASSWORD field in the connection properties. You can enable catalog encryption or only password encryption. see Connection Password Encryption.
	// +kubebuilder:validation:Required
	ConnectionPasswordEncryption []ConnectionPasswordEncryptionParameters `json:"connectionPasswordEncryption" tf:"connection_password_encryption,omitempty"`

	// Specifies the encryption-at-rest configuration for the Data Catalog. see Encryption At Rest.
	// +kubebuilder:validation:Required
	EncryptionAtRest []EncryptionAtRestParameters `json:"encryptionAtRest" tf:"encryption_at_rest,omitempty"`
}

type DataCatalogEncryptionSettingsObservation struct {

	// The ID of the Data Catalog to set the security configuration for.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataCatalogEncryptionSettingsParameters struct {

	// –  The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// –  The security configuration to set. see Data Catalog Encryption Settings.
	// +kubebuilder:validation:Required
	DataCatalogEncryptionSettings []DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsParameters `json:"dataCatalogEncryptionSettings" tf:"data_catalog_encryption_settings,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type EncryptionAtRestObservation struct {
}

type EncryptionAtRestParameters struct {

	// The encryption-at-rest mode for encrypting Data Catalog data. Valid values are DISABLED and SSE-KMS.
	// +kubebuilder:validation:Required
	CatalogEncryptionMode *string `json:"catalogEncryptionMode" tf:"catalog_encryption_mode,omitempty"`

	// The ARN of the AWS KMS key to use for encryption at rest.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	SseAwsKMSKeyID *string `json:"sseAwsKmsKeyId,omitempty" tf:"sse_aws_kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	SseAwsKMSKeyIDRef *v1.Reference `json:"sseAwsKmsKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SseAwsKMSKeyIDSelector *v1.Selector `json:"sseAwsKmsKeyIdSelector,omitempty" tf:"-"`
}

// DataCatalogEncryptionSettingsSpec defines the desired state of DataCatalogEncryptionSettings
type DataCatalogEncryptionSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataCatalogEncryptionSettingsParameters `json:"forProvider"`
}

// DataCatalogEncryptionSettingsStatus defines the observed state of DataCatalogEncryptionSettings.
type DataCatalogEncryptionSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataCatalogEncryptionSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataCatalogEncryptionSettings is the Schema for the DataCatalogEncryptionSettingss API. Provides a Glue Data Catalog Encryption Settings resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DataCatalogEncryptionSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataCatalogEncryptionSettingsSpec   `json:"spec"`
	Status            DataCatalogEncryptionSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataCatalogEncryptionSettingsList contains a list of DataCatalogEncryptionSettingss
type DataCatalogEncryptionSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalogEncryptionSettings `json:"items"`
}

// Repository type metadata.
var (
	DataCatalogEncryptionSettings_Kind             = "DataCatalogEncryptionSettings"
	DataCatalogEncryptionSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataCatalogEncryptionSettings_Kind}.String()
	DataCatalogEncryptionSettings_KindAPIVersion   = DataCatalogEncryptionSettings_Kind + "." + CRDGroupVersion.String()
	DataCatalogEncryptionSettings_GroupVersionKind = CRDGroupVersion.WithKind(DataCatalogEncryptionSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&DataCatalogEncryptionSettings{}, &DataCatalogEncryptionSettingsList{})
}
