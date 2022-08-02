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

type AutoscalingGroupsObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AutoscalingGroupsParameters struct {
}

type LaunchTemplateObservation struct {
}

type LaunchTemplateParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type NodeGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type NodeGroupParameters struct {

	// +kubebuilder:validation:Optional
	AMIType *string `json:"amiType,omitempty" tf:"ami_type,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityType *string `json:"capacityType,omitempty" tf:"capacity_type,omitempty"`

	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Optional
	ForceUpdateVersion *bool `json:"forceUpdateVersion,omitempty" tf:"force_update_version,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	NodeRoleArn *string `json:"nodeRoleArn,omitempty" tf:"node_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	NodeRoleArnRef *v1.Reference `json:"nodeRoleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NodeRoleArnSelector *v1.Selector `json:"nodeRoleArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReleaseVersion *string `json:"releaseVersion,omitempty" tf:"release_version,omitempty"`

	// +kubebuilder:validation:Optional
	RemoteAccess []RemoteAccessParameters `json:"remoteAccess,omitempty" tf:"remote_access,omitempty"`

	// +kubebuilder:validation:Required
	ScalingConfig []ScalingConfigParameters `json:"scalingConfig" tf:"scaling_config,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Taint []TaintParameters `json:"taint,omitempty" tf:"taint,omitempty"`

	// +kubebuilder:validation:Optional
	UpdateConfig []UpdateConfigParameters `json:"updateConfig,omitempty" tf:"update_config,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type RemoteAccessObservation struct {
}

type RemoteAccessParameters struct {

	// +kubebuilder:validation:Optional
	EC2SSHKey *string `json:"ec2SshKey,omitempty" tf:"ec2_ssh_key,omitempty"`

	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDRefs []v1.Reference `json:"sourceSecurityGroupIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDSelector *v1.Selector `json:"sourceSecurityGroupIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SourceSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SourceSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SourceSecurityGroupIds []*string `json:"sourceSecurityGroupIds,omitempty" tf:"source_security_group_ids,omitempty"`
}

type ResourcesObservation struct {
	AutoscalingGroups []AutoscalingGroupsObservation `json:"autoscalingGroups,omitempty" tf:"autoscaling_groups,omitempty"`

	RemoteAccessSecurityGroupID *string `json:"remoteAccessSecurityGroupId,omitempty" tf:"remote_access_security_group_id,omitempty"`
}

type ResourcesParameters struct {
}

type ScalingConfigObservation struct {
}

type ScalingConfigParameters struct {

	// +kubebuilder:validation:Required
	DesiredSize *float64 `json:"desiredSize" tf:"desired_size,omitempty"`

	// +kubebuilder:validation:Required
	MaxSize *float64 `json:"maxSize" tf:"max_size,omitempty"`

	// +kubebuilder:validation:Required
	MinSize *float64 `json:"minSize" tf:"min_size,omitempty"`
}

type TaintObservation struct {
}

type TaintParameters struct {

	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type UpdateConfigObservation struct {
}

type UpdateConfigParameters struct {

	// +kubebuilder:validation:Optional
	MaxUnavailable *float64 `json:"maxUnavailable,omitempty" tf:"max_unavailable,omitempty"`

	// +kubebuilder:validation:Optional
	MaxUnavailablePercentage *float64 `json:"maxUnavailablePercentage,omitempty" tf:"max_unavailable_percentage,omitempty"`
}

// NodeGroupSpec defines the desired state of NodeGroup
type NodeGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeGroupParameters `json:"forProvider"`
}

// NodeGroupStatus defines the observed state of NodeGroup.
type NodeGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroup is the Schema for the NodeGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeGroupSpec   `json:"spec"`
	Status            NodeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroupList contains a list of NodeGroups
type NodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeGroup `json:"items"`
}

// Repository type metadata.
var (
	NodeGroup_Kind             = "NodeGroup"
	NodeGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeGroup_Kind}.String()
	NodeGroup_KindAPIVersion   = NodeGroup_Kind + "." + CRDGroupVersion.String()
	NodeGroup_GroupVersionKind = CRDGroupVersion.WithKind(NodeGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeGroup{}, &NodeGroupList{})
}
