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

type AttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AttachmentParameters struct {

	// +crossplane:generate:reference:type=ELB
	// +kubebuilder:validation:Optional
	ELB *string `json:"elb,omitempty" tf:"elb,omitempty"`

	// +kubebuilder:validation:Optional
	ELBRef *v1.Reference `json:"elbRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ELBSelector *v1.Selector `json:"elbSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Instance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// AttachmentSpec defines the desired state of Attachment
type AttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttachmentParameters `json:"forProvider"`
}

// AttachmentStatus defines the observed state of Attachment.
type AttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Attachment is the Schema for the Attachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Attachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AttachmentSpec   `json:"spec"`
	Status            AttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttachmentList contains a list of Attachments
type AttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Attachment `json:"items"`
}

// Repository type metadata.
var (
	Attachment_Kind             = "Attachment"
	Attachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Attachment_Kind}.String()
	Attachment_KindAPIVersion   = Attachment_Kind + "." + CRDGroupVersion.String()
	Attachment_GroupVersionKind = CRDGroupVersion.WithKind(Attachment_Kind)
)

func init() {
	SchemeBuilder.Register(&Attachment{}, &AttachmentList{})
}
