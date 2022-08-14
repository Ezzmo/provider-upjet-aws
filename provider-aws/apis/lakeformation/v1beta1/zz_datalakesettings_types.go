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

type CreateDatabaseDefaultPermissionsObservation struct {
}

type CreateDatabaseDefaultPermissionsParameters struct {

	// +kubebuilder:validation:Optional
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`
}

type CreateTableDefaultPermissionsObservation struct {
}

type CreateTableDefaultPermissionsParameters struct {

	// +kubebuilder:validation:Optional
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`
}

type DataLakeSettingsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataLakeSettingsParameters struct {

	// –  Set of ARNs of AWS Lake Formation principals .
	// +kubebuilder:validation:Optional
	Admins []*string `json:"admins,omitempty" tf:"admins,omitempty"`

	// –  Identifier for the Data Catalog. By default, the account ID.
	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	// +kubebuilder:validation:Optional
	CreateDatabaseDefaultPermissions []CreateDatabaseDefaultPermissionsParameters `json:"createDatabaseDefaultPermissions,omitempty" tf:"create_database_default_permissions,omitempty"`

	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	// +kubebuilder:validation:Optional
	CreateTableDefaultPermissions []CreateTableDefaultPermissionsParameters `json:"createTableDefaultPermissions,omitempty" tf:"create_table_default_permissions,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// wning account IDs that the caller's account can use to share their user access details .
	// +kubebuilder:validation:Optional
	TrustedResourceOwners []*string `json:"trustedResourceOwners,omitempty" tf:"trusted_resource_owners,omitempty"`
}

// DataLakeSettingsSpec defines the desired state of DataLakeSettings
type DataLakeSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataLakeSettingsParameters `json:"forProvider"`
}

// DataLakeSettingsStatus defines the observed state of DataLakeSettings.
type DataLakeSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataLakeSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataLakeSettings is the Schema for the DataLakeSettingss API. Manages data lake administrators and default database and table permissions
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DataLakeSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataLakeSettingsSpec   `json:"spec"`
	Status            DataLakeSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataLakeSettingsList contains a list of DataLakeSettingss
type DataLakeSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataLakeSettings `json:"items"`
}

// Repository type metadata.
var (
	DataLakeSettings_Kind             = "DataLakeSettings"
	DataLakeSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataLakeSettings_Kind}.String()
	DataLakeSettings_KindAPIVersion   = DataLakeSettings_Kind + "." + CRDGroupVersion.String()
	DataLakeSettings_GroupVersionKind = CRDGroupVersion.WithKind(DataLakeSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&DataLakeSettings{}, &DataLakeSettingsList{})
}
