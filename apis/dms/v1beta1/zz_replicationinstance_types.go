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

type ReplicationInstanceInitParameters struct {

	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage *int64 `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`

	// Indicates that major version upgrades are allowed.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The engine version number of the replication instance.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Specifies if the replication instance is a multi-az deployment. You cannot set the availability_zone parameter if the multi_az parameter is set to true.
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// The type of IP address protocol used by a replication instance. Valid values: IPV4, DUAL.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// The compute and memory capacity of the replication instance as specified by the replication instance class. See AWS DMS User Guide for available instance sizes and advice on which one to choose.
	ReplicationInstanceClass *string `json:"replicationInstanceClass,omitempty" tf:"replication_instance_class,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ReplicationInstanceObservation struct {

	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage *int64 `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`

	// Indicates that major version upgrades are allowed.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The engine version number of the replication instance.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for kms_key_arn, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Specifies if the replication instance is a multi-az deployment. You cannot set the availability_zone parameter if the multi_az parameter is set to true.
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// The type of IP address protocol used by a replication instance. Valid values: IPV4, DUAL.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn *string `json:"replicationInstanceArn,omitempty" tf:"replication_instance_arn,omitempty"`

	// The compute and memory capacity of the replication instance as specified by the replication instance class. See AWS DMS User Guide for available instance sizes and advice on which one to choose.
	ReplicationInstanceClass *string `json:"replicationInstanceClass,omitempty" tf:"replication_instance_class,omitempty"`

	// A list of the private IP addresses of the replication instance.
	ReplicationInstancePrivateIps []*string `json:"replicationInstancePrivateIps,omitempty" tf:"replication_instance_private_ips,omitempty"`

	// A list of the public IP addresses of the replication instance.
	ReplicationInstancePublicIps []*string `json:"replicationInstancePublicIps,omitempty" tf:"replication_instance_public_ips,omitempty"`

	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupID *string `json:"replicationSubnetGroupId,omitempty" tf:"replication_subnet_group_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type ReplicationInstanceParameters struct {

	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	// +kubebuilder:validation:Optional
	AllocatedStorage *int64 `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`

	// Indicates that major version upgrades are allowed.
	// +kubebuilder:validation:Optional
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// The EC2 Availability Zone that the replication instance will be created in.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The engine version number of the replication instance.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for kms_key_arn, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// Specifies if the replication instance is a multi-az deployment. You cannot set the availability_zone parameter if the multi_az parameter is set to true.
	// +kubebuilder:validation:Optional
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// The type of IP address protocol used by a replication instance. Valid values: IPV4, DUAL.
	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	// +kubebuilder:validation:Optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The compute and memory capacity of the replication instance as specified by the replication instance class. See AWS DMS User Guide for available instance sizes and advice on which one to choose.
	// +kubebuilder:validation:Optional
	ReplicationInstanceClass *string `json:"replicationInstanceClass,omitempty" tf:"replication_instance_class,omitempty"`

	// A subnet group to associate with the replication instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/dms/v1beta1.ReplicationSubnetGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ReplicationSubnetGroupID *string `json:"replicationSubnetGroupId,omitempty" tf:"replication_subnet_group_id,omitempty"`

	// Reference to a ReplicationSubnetGroup in dms to populate replicationSubnetGroupId.
	// +kubebuilder:validation:Optional
	ReplicationSubnetGroupIDRef *v1.Reference `json:"replicationSubnetGroupIdRef,omitempty" tf:"-"`

	// Selector for a ReplicationSubnetGroup in dms to populate replicationSubnetGroupId.
	// +kubebuilder:validation:Optional
	ReplicationSubnetGroupIDSelector *v1.Selector `json:"replicationSubnetGroupIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

// ReplicationInstanceSpec defines the desired state of ReplicationInstance
type ReplicationInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicationInstanceParameters `json:"forProvider"`
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
	InitProvider ReplicationInstanceInitParameters `json:"initProvider,omitempty"`
}

// ReplicationInstanceStatus defines the observed state of ReplicationInstance.
type ReplicationInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicationInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationInstance is the Schema for the ReplicationInstances API. Provides a DMS (Data Migration Service) replication instance resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReplicationInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.replicationInstanceClass) || (has(self.initProvider) && has(self.initProvider.replicationInstanceClass))",message="spec.forProvider.replicationInstanceClass is a required parameter"
	Spec   ReplicationInstanceSpec   `json:"spec"`
	Status ReplicationInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationInstanceList contains a list of ReplicationInstances
type ReplicationInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationInstance `json:"items"`
}

// Repository type metadata.
var (
	ReplicationInstance_Kind             = "ReplicationInstance"
	ReplicationInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReplicationInstance_Kind}.String()
	ReplicationInstance_KindAPIVersion   = ReplicationInstance_Kind + "." + CRDGroupVersion.String()
	ReplicationInstance_GroupVersionKind = CRDGroupVersion.WithKind(ReplicationInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&ReplicationInstance{}, &ReplicationInstanceList{})
}
