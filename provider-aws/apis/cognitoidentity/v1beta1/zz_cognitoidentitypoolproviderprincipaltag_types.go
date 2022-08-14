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

type CognitoIdentityPoolProviderPrincipalTagObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CognitoIdentityPoolProviderPrincipalTagParameters struct {

	// An identity pool ID.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/cognitoidentity/v1beta1.Pool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IdentityPoolID *string `json:"identityPoolId,omitempty" tf:"identity_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityPoolIDRef *v1.Reference `json:"identityPoolIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IdentityPoolIDSelector *v1.Selector `json:"identityPoolIdSelector,omitempty" tf:"-"`

	// The name of the identity provider.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/cognitoidp/v1beta1.UserPool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("endpoint",true)
	// +kubebuilder:validation:Optional
	IdentityProviderName *string `json:"identityProviderName,omitempty" tf:"identity_provider_name,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityProviderNameRef *v1.Reference `json:"identityProviderNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IdentityProviderNameSelector *v1.Selector `json:"identityProviderNameSelector,omitempty" tf:"-"`

	// String to string map of variables.
	// +kubebuilder:validation:Optional
	PrincipalTags map[string]*string `json:"principalTags,omitempty" tf:"principal_tags,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// :  use default  attribute mappings.
	// +kubebuilder:validation:Optional
	UseDefaults *bool `json:"useDefaults,omitempty" tf:"use_defaults,omitempty"`
}

// CognitoIdentityPoolProviderPrincipalTagSpec defines the desired state of CognitoIdentityPoolProviderPrincipalTag
type CognitoIdentityPoolProviderPrincipalTagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CognitoIdentityPoolProviderPrincipalTagParameters `json:"forProvider"`
}

// CognitoIdentityPoolProviderPrincipalTagStatus defines the observed state of CognitoIdentityPoolProviderPrincipalTag.
type CognitoIdentityPoolProviderPrincipalTagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CognitoIdentityPoolProviderPrincipalTagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoIdentityPoolProviderPrincipalTag is the Schema for the CognitoIdentityPoolProviderPrincipalTags API. Provides an AWS Cognito Identity Principal Mapping.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CognitoIdentityPoolProviderPrincipalTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoIdentityPoolProviderPrincipalTagSpec   `json:"spec"`
	Status            CognitoIdentityPoolProviderPrincipalTagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoIdentityPoolProviderPrincipalTagList contains a list of CognitoIdentityPoolProviderPrincipalTags
type CognitoIdentityPoolProviderPrincipalTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoIdentityPoolProviderPrincipalTag `json:"items"`
}

// Repository type metadata.
var (
	CognitoIdentityPoolProviderPrincipalTag_Kind             = "CognitoIdentityPoolProviderPrincipalTag"
	CognitoIdentityPoolProviderPrincipalTag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CognitoIdentityPoolProviderPrincipalTag_Kind}.String()
	CognitoIdentityPoolProviderPrincipalTag_KindAPIVersion   = CognitoIdentityPoolProviderPrincipalTag_Kind + "." + CRDGroupVersion.String()
	CognitoIdentityPoolProviderPrincipalTag_GroupVersionKind = CRDGroupVersion.WithKind(CognitoIdentityPoolProviderPrincipalTag_Kind)
)

func init() {
	SchemeBuilder.Register(&CognitoIdentityPoolProviderPrincipalTag{}, &CognitoIdentityPoolProviderPrincipalTagList{})
}
