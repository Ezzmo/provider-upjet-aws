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

type ClusterInitParameters struct {

	// If true , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default is true.
	AllowVersionUpgrade *bool `json:"allowVersionUpgrade,omitempty" tf:"allow_version_upgrade,omitempty"`

	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is false.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored.
	// No longer supported by the AWS API.
	// Always returns auto.
	AquaConfigurationStatus *string `json:"aquaConfigurationStatus,omitempty" tf:"aqua_configuration_status,omitempty"`

	// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with create-cluster-snapshot. Default is 1.
	AutomatedSnapshotRetentionPeriod *int64 `json:"automatedSnapshotRetentionPeriod,omitempty" tf:"automated_snapshot_retention_period,omitempty"`

	// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency. Can only be changed if availability_zone_relocation_enabled is true.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// If true, the cluster can be relocated to another availabity zone, either automatically by AWS or when requested. Default is false. Available for use on clusters from the RA3 instance family.
	AvailabilityZoneRelocationEnabled *bool `json:"availabilityZoneRelocationEnabled,omitempty" tf:"availability_zone_relocation_enabled,omitempty"`

	// The name of the parameter group to be associated with this cluster.
	ClusterParameterGroupName *string `json:"clusterParameterGroupName,omitempty" tf:"cluster_parameter_group_name,omitempty"`

	// The public key for the cluster
	ClusterPublicKey *string `json:"clusterPublicKey,omitempty" tf:"cluster_public_key,omitempty"`

	// The specific revision number of the database in the cluster
	ClusterRevisionNumber *string `json:"clusterRevisionNumber,omitempty" tf:"cluster_revision_number,omitempty"`

	// The name of a cluster subnet group to be associated with this cluster. If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
	ClusterSubnetGroupName *string `json:"clusterSubnetGroupName,omitempty" tf:"cluster_subnet_group_name,omitempty"`

	// The cluster type to use. Either single-node or multi-node.
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// The version of the Amazon Redshift engine software that you want to deploy on the cluster.
	// The version selected runs on all the nodes in the cluster.
	ClusterVersion *string `json:"clusterVersion,omitempty" tf:"cluster_version,omitempty"`

	// The name of the first database to be created when the cluster is created.
	// If you do not provide a name, Amazon Redshift will create a default database called dev.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// The Elastic IP (EIP) address for the cluster.
	ElasticIP *string `json:"elasticIp,omitempty" tf:"elastic_ip,omitempty"`

	// If true , the data in the cluster is encrypted at rest.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The connection endpoint
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// If true , enhanced VPC routing is enabled.
	EnhancedVPCRouting *bool `json:"enhancedVpcRouting,omitempty" tf:"enhanced_vpc_routing,omitempty"`

	// The identifier of the final snapshot that is to be created immediately before deleting the cluster. If this parameter is provided, skip_final_snapshot must be false.
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// Logging, documented below.
	Logging []LoggingInitParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// The name of the maintenance track for the restored cluster. When you take a snapshot, the snapshot inherits the MaintenanceTrack value from the cluster. The snapshot might be on a different track than the cluster that was the source for the snapshot. For example, suppose that you take a snapshot of  a cluster that is on the current track and then change the cluster to be on the trailing track. In this case, the snapshot and the source cluster are on different tracks. Default value is current.
	MaintenanceTrackName *string `json:"maintenanceTrackName,omitempty" tf:"maintenance_track_name,omitempty"`

	// The default number of days to retain a manual snapshot. If the value is -1, the snapshot is retained indefinitely. This setting doesn't change the retention period of existing snapshots. Valid values are between -1 and 3653. Default value is -1.
	ManualSnapshotRetentionPeriod *int64 `json:"manualSnapshotRetentionPeriod,omitempty" tf:"manual_snapshot_retention_period,omitempty"`

	// Username for the master DB user.
	MasterUsername *string `json:"masterUsername,omitempty" tf:"master_username,omitempty"`

	// The node type to be provisioned for the cluster.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node. Default is 1.
	NumberOfNodes *int64 `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`

	// The AWS customer account used to create or copy the snapshot. Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
	OwnerAccount *string `json:"ownerAccount,omitempty" tf:"owner_account,omitempty"`

	// The port number on which the cluster accepts incoming connections. Valid values are between 1115 and 65535.
	// The cluster is accessible only via the JDBC and ODBC connection strings.
	// Part of the connection string requires the port on which the cluster will listen for incoming connections.
	// Default port is 5439.
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// The weekly time range (in UTC) during which automated cluster maintenance can occur.
	// Format: ddd:hh24:mi-ddd:hh24:mi
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// If true, the cluster can be accessed from a public network. Default is true.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster. If true , a final cluster snapshot is not created. If false , a final cluster snapshot is created before the cluster is deleted. Default is false.
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// The name of the cluster the source snapshot was created from.
	SnapshotClusterIdentifier *string `json:"snapshotClusterIdentifier,omitempty" tf:"snapshot_cluster_identifier,omitempty"`

	// Configuration of automatic copy of snapshots from one region to another. Documented below.
	SnapshotCopy []SnapshotCopyInitParameters `json:"snapshotCopy,omitempty" tf:"snapshot_copy,omitempty"`

	// The name of the snapshot from which to create the new cluster.
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ClusterNodesInitParameters struct {
}

type ClusterNodesObservation struct {

	// Whether the node is a leader node or a compute node
	NodeRole *string `json:"nodeRole,omitempty" tf:"node_role,omitempty"`

	// The private IP address of a node within a cluster
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The public IP address of a node within a cluster
	PublicIPAddress *string `json:"publicIpAddress,omitempty" tf:"public_ip_address,omitempty"`
}

type ClusterNodesParameters struct {
}

type ClusterObservation struct {

	// If true , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default is true.
	AllowVersionUpgrade *bool `json:"allowVersionUpgrade,omitempty" tf:"allow_version_upgrade,omitempty"`

	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is false.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored.
	// No longer supported by the AWS API.
	// Always returns auto.
	AquaConfigurationStatus *string `json:"aquaConfigurationStatus,omitempty" tf:"aqua_configuration_status,omitempty"`

	// Amazon Resource Name (ARN) of cluster
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with create-cluster-snapshot. Default is 1.
	AutomatedSnapshotRetentionPeriod *int64 `json:"automatedSnapshotRetentionPeriod,omitempty" tf:"automated_snapshot_retention_period,omitempty"`

	// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency. Can only be changed if availability_zone_relocation_enabled is true.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// If true, the cluster can be relocated to another availabity zone, either automatically by AWS or when requested. Default is false. Available for use on clusters from the RA3 instance family.
	AvailabilityZoneRelocationEnabled *bool `json:"availabilityZoneRelocationEnabled,omitempty" tf:"availability_zone_relocation_enabled,omitempty"`

	// The namespace Amazon Resource Name (ARN) of the cluster
	ClusterNamespaceArn *string `json:"clusterNamespaceArn,omitempty" tf:"cluster_namespace_arn,omitempty"`

	// The nodes in the cluster. Cluster node blocks are documented below
	ClusterNodes []ClusterNodesObservation `json:"clusterNodes,omitempty" tf:"cluster_nodes,omitempty"`

	// The name of the parameter group to be associated with this cluster.
	ClusterParameterGroupName *string `json:"clusterParameterGroupName,omitempty" tf:"cluster_parameter_group_name,omitempty"`

	// The public key for the cluster
	ClusterPublicKey *string `json:"clusterPublicKey,omitempty" tf:"cluster_public_key,omitempty"`

	// The specific revision number of the database in the cluster
	ClusterRevisionNumber *string `json:"clusterRevisionNumber,omitempty" tf:"cluster_revision_number,omitempty"`

	// The name of a cluster subnet group to be associated with this cluster. If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
	ClusterSubnetGroupName *string `json:"clusterSubnetGroupName,omitempty" tf:"cluster_subnet_group_name,omitempty"`

	// The cluster type to use. Either single-node or multi-node.
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// The version of the Amazon Redshift engine software that you want to deploy on the cluster.
	// The version selected runs on all the nodes in the cluster.
	ClusterVersion *string `json:"clusterVersion,omitempty" tf:"cluster_version,omitempty"`

	// The DNS name of the cluster
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// The name of the first database to be created when the cluster is created.
	// If you do not provide a name, Amazon Redshift will create a default database called dev.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// The Amazon Resource Name (ARN) for the IAM role that was set as default for the cluster when the cluster was created.
	DefaultIAMRoleArn *string `json:"defaultIamRoleArn,omitempty" tf:"default_iam_role_arn,omitempty"`

	// The Elastic IP (EIP) address for the cluster.
	ElasticIP *string `json:"elasticIp,omitempty" tf:"elastic_ip,omitempty"`

	// If true , the data in the cluster is encrypted at rest.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The connection endpoint
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// If true , enhanced VPC routing is enabled.
	EnhancedVPCRouting *bool `json:"enhancedVpcRouting,omitempty" tf:"enhanced_vpc_routing,omitempty"`

	// The identifier of the final snapshot that is to be created immediately before deleting the cluster. If this parameter is provided, skip_final_snapshot must be false.
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
	IAMRoles []*string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`

	// The Redshift Cluster ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Logging, documented below.
	Logging []LoggingObservation `json:"logging,omitempty" tf:"logging,omitempty"`

	// The name of the maintenance track for the restored cluster. When you take a snapshot, the snapshot inherits the MaintenanceTrack value from the cluster. The snapshot might be on a different track than the cluster that was the source for the snapshot. For example, suppose that you take a snapshot of  a cluster that is on the current track and then change the cluster to be on the trailing track. In this case, the snapshot and the source cluster are on different tracks. Default value is current.
	MaintenanceTrackName *string `json:"maintenanceTrackName,omitempty" tf:"maintenance_track_name,omitempty"`

	// The default number of days to retain a manual snapshot. If the value is -1, the snapshot is retained indefinitely. This setting doesn't change the retention period of existing snapshots. Valid values are between -1 and 3653. Default value is -1.
	ManualSnapshotRetentionPeriod *int64 `json:"manualSnapshotRetentionPeriod,omitempty" tf:"manual_snapshot_retention_period,omitempty"`

	// Username for the master DB user.
	MasterUsername *string `json:"masterUsername,omitempty" tf:"master_username,omitempty"`

	// The node type to be provisioned for the cluster.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node. Default is 1.
	NumberOfNodes *int64 `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`

	// The AWS customer account used to create or copy the snapshot. Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
	OwnerAccount *string `json:"ownerAccount,omitempty" tf:"owner_account,omitempty"`

	// The port number on which the cluster accepts incoming connections. Valid values are between 1115 and 65535.
	// The cluster is accessible only via the JDBC and ODBC connection strings.
	// Part of the connection string requires the port on which the cluster will listen for incoming connections.
	// Default port is 5439.
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// The weekly time range (in UTC) during which automated cluster maintenance can occur.
	// Format: ddd:hh24:mi-ddd:hh24:mi
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// If true, the cluster can be accessed from a public network. Default is true.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster. If true , a final cluster snapshot is not created. If false , a final cluster snapshot is created before the cluster is deleted. Default is false.
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// The name of the cluster the source snapshot was created from.
	SnapshotClusterIdentifier *string `json:"snapshotClusterIdentifier,omitempty" tf:"snapshot_cluster_identifier,omitempty"`

	// Configuration of automatic copy of snapshots from one region to another. Documented below.
	SnapshotCopy []SnapshotCopyObservation `json:"snapshotCopy,omitempty" tf:"snapshot_copy,omitempty"`

	// The name of the snapshot from which to create the new cluster.
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type ClusterParameters struct {

	// If true , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default is true.
	// +kubebuilder:validation:Optional
	AllowVersionUpgrade *bool `json:"allowVersionUpgrade,omitempty" tf:"allow_version_upgrade,omitempty"`

	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is false.
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored.
	// No longer supported by the AWS API.
	// Always returns auto.
	// +kubebuilder:validation:Optional
	AquaConfigurationStatus *string `json:"aquaConfigurationStatus,omitempty" tf:"aqua_configuration_status,omitempty"`

	// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with create-cluster-snapshot. Default is 1.
	// +kubebuilder:validation:Optional
	AutomatedSnapshotRetentionPeriod *int64 `json:"automatedSnapshotRetentionPeriod,omitempty" tf:"automated_snapshot_retention_period,omitempty"`

	// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency. Can only be changed if availability_zone_relocation_enabled is true.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// If true, the cluster can be relocated to another availabity zone, either automatically by AWS or when requested. Default is false. Available for use on clusters from the RA3 instance family.
	// +kubebuilder:validation:Optional
	AvailabilityZoneRelocationEnabled *bool `json:"availabilityZoneRelocationEnabled,omitempty" tf:"availability_zone_relocation_enabled,omitempty"`

	// The name of the parameter group to be associated with this cluster.
	// +kubebuilder:validation:Optional
	ClusterParameterGroupName *string `json:"clusterParameterGroupName,omitempty" tf:"cluster_parameter_group_name,omitempty"`

	// The public key for the cluster
	// +kubebuilder:validation:Optional
	ClusterPublicKey *string `json:"clusterPublicKey,omitempty" tf:"cluster_public_key,omitempty"`

	// The specific revision number of the database in the cluster
	// +kubebuilder:validation:Optional
	ClusterRevisionNumber *string `json:"clusterRevisionNumber,omitempty" tf:"cluster_revision_number,omitempty"`

	// The name of a cluster subnet group to be associated with this cluster. If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
	// +kubebuilder:validation:Optional
	ClusterSubnetGroupName *string `json:"clusterSubnetGroupName,omitempty" tf:"cluster_subnet_group_name,omitempty"`

	// The cluster type to use. Either single-node or multi-node.
	// +kubebuilder:validation:Optional
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// The version of the Amazon Redshift engine software that you want to deploy on the cluster.
	// The version selected runs on all the nodes in the cluster.
	// +kubebuilder:validation:Optional
	ClusterVersion *string `json:"clusterVersion,omitempty" tf:"cluster_version,omitempty"`

	// The name of the first database to be created when the cluster is created.
	// If you do not provide a name, Amazon Redshift will create a default database called dev.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// The Amazon Resource Name (ARN) for the IAM role that was set as default for the cluster when the cluster was created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	DefaultIAMRoleArn *string `json:"defaultIamRoleArn,omitempty" tf:"default_iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate defaultIamRoleArn.
	// +kubebuilder:validation:Optional
	DefaultIAMRoleArnRef *v1.Reference `json:"defaultIamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate defaultIamRoleArn.
	// +kubebuilder:validation:Optional
	DefaultIAMRoleArnSelector *v1.Selector `json:"defaultIamRoleArnSelector,omitempty" tf:"-"`

	// The Elastic IP (EIP) address for the cluster.
	// +kubebuilder:validation:Optional
	ElasticIP *string `json:"elasticIp,omitempty" tf:"elastic_ip,omitempty"`

	// If true , the data in the cluster is encrypted at rest.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The connection endpoint
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// If true , enhanced VPC routing is enabled.
	// +kubebuilder:validation:Optional
	EnhancedVPCRouting *bool `json:"enhancedVpcRouting,omitempty" tf:"enhanced_vpc_routing,omitempty"`

	// The identifier of the final snapshot that is to be created immediately before deleting the cluster. If this parameter is provided, skip_final_snapshot must be false.
	// +kubebuilder:validation:Optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// References to Role in iam to populate iamRoles.
	// +kubebuilder:validation:Optional
	IAMRoleRefs []v1.Reference `json:"iamRoleRefs,omitempty" tf:"-"`

	// Selector for a list of Role in iam to populate iamRoles.
	// +kubebuilder:validation:Optional
	IAMRoleSelector *v1.Selector `json:"iamRoleSelector,omitempty" tf:"-"`

	// A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:refFieldName=IAMRoleRefs
	// +crossplane:generate:reference:selectorFieldName=IAMRoleSelector
	// +kubebuilder:validation:Optional
	IAMRoles []*string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`

	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Logging, documented below.
	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// The name of the maintenance track for the restored cluster. When you take a snapshot, the snapshot inherits the MaintenanceTrack value from the cluster. The snapshot might be on a different track than the cluster that was the source for the snapshot. For example, suppose that you take a snapshot of  a cluster that is on the current track and then change the cluster to be on the trailing track. In this case, the snapshot and the source cluster are on different tracks. Default value is current.
	// +kubebuilder:validation:Optional
	MaintenanceTrackName *string `json:"maintenanceTrackName,omitempty" tf:"maintenance_track_name,omitempty"`

	// The default number of days to retain a manual snapshot. If the value is -1, the snapshot is retained indefinitely. This setting doesn't change the retention period of existing snapshots. Valid values are between -1 and 3653. Default value is -1.
	// +kubebuilder:validation:Optional
	ManualSnapshotRetentionPeriod *int64 `json:"manualSnapshotRetentionPeriod,omitempty" tf:"manual_snapshot_retention_period,omitempty"`

	// Password for the master DB user.
	// Note that this may show up in logs, and it will be stored in the state file. Password must contain at least 8 chars and
	// contain at least one uppercase letter, one lowercase letter, and one number.
	// +kubebuilder:validation:Optional
	MasterPasswordSecretRef *v1.SecretKeySelector `json:"masterPasswordSecretRef,omitempty" tf:"-"`

	// Username for the master DB user.
	// +kubebuilder:validation:Optional
	MasterUsername *string `json:"masterUsername,omitempty" tf:"master_username,omitempty"`

	// The node type to be provisioned for the cluster.
	// +kubebuilder:validation:Optional
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node. Default is 1.
	// +kubebuilder:validation:Optional
	NumberOfNodes *int64 `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`

	// The AWS customer account used to create or copy the snapshot. Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
	// +kubebuilder:validation:Optional
	OwnerAccount *string `json:"ownerAccount,omitempty" tf:"owner_account,omitempty"`

	// The port number on which the cluster accepts incoming connections. Valid values are between 1115 and 65535.
	// The cluster is accessible only via the JDBC and ODBC connection strings.
	// Part of the connection string requires the port on which the cluster will listen for incoming connections.
	// Default port is 5439.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// The weekly time range (in UTC) during which automated cluster maintenance can occur.
	// Format: ddd:hh24:mi-ddd:hh24:mi
	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// If true, the cluster can be accessed from a public network. Default is true.
	// +kubebuilder:validation:Optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster. If true , a final cluster snapshot is not created. If false , a final cluster snapshot is created before the cluster is deleted. Default is false.
	// +kubebuilder:validation:Optional
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// The name of the cluster the source snapshot was created from.
	// +kubebuilder:validation:Optional
	SnapshotClusterIdentifier *string `json:"snapshotClusterIdentifier,omitempty" tf:"snapshot_cluster_identifier,omitempty"`

	// Configuration of automatic copy of snapshots from one region to another. Documented below.
	// +kubebuilder:validation:Optional
	SnapshotCopy []SnapshotCopyParameters `json:"snapshotCopy,omitempty" tf:"snapshot_copy,omitempty"`

	// The name of the snapshot from which to create the new cluster.
	// +kubebuilder:validation:Optional
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type LoggingInitParameters struct {

	// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
	// For more information on the permissions required for the bucket, please read the AWS documentation
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// The log destination type. An enum with possible values of s3 and cloudwatch.
	LogDestinationType *string `json:"logDestinationType,omitempty" tf:"log_destination_type,omitempty"`

	// The collection of exported log types. Log types include the connection log, user log and user activity log. Required when log_destination_type is cloudwatch. Valid log types are connectionlog, userlog, and useractivitylog.
	LogExports []*string `json:"logExports,omitempty" tf:"log_exports,omitempty"`

	// The prefix applied to the log file names.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type LoggingObservation struct {

	// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
	// For more information on the permissions required for the bucket, please read the AWS documentation
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// The log destination type. An enum with possible values of s3 and cloudwatch.
	LogDestinationType *string `json:"logDestinationType,omitempty" tf:"log_destination_type,omitempty"`

	// The collection of exported log types. Log types include the connection log, user log and user activity log. Required when log_destination_type is cloudwatch. Valid log types are connectionlog, userlog, and useractivitylog.
	LogExports []*string `json:"logExports,omitempty" tf:"log_exports,omitempty"`

	// The prefix applied to the log file names.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type LoggingParameters struct {

	// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
	// For more information on the permissions required for the bucket, please read the AWS documentation
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable" tf:"enable,omitempty"`

	// The log destination type. An enum with possible values of s3 and cloudwatch.
	// +kubebuilder:validation:Optional
	LogDestinationType *string `json:"logDestinationType,omitempty" tf:"log_destination_type,omitempty"`

	// The collection of exported log types. Log types include the connection log, user log and user activity log. Required when log_destination_type is cloudwatch. Valid log types are connectionlog, userlog, and useractivitylog.
	// +kubebuilder:validation:Optional
	LogExports []*string `json:"logExports,omitempty" tf:"log_exports,omitempty"`

	// The prefix applied to the log file names.
	// +kubebuilder:validation:Optional
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type SnapshotCopyInitParameters struct {

	// The destination region that you want to copy snapshots to.
	DestinationRegion *string `json:"destinationRegion,omitempty" tf:"destination_region,omitempty"`

	// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
	GrantName *string `json:"grantName,omitempty" tf:"grant_name,omitempty"`

	// The number of days to retain automated snapshots in the destination region after they are copied from the source region. Defaults to 7.
	RetentionPeriod *int64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

type SnapshotCopyObservation struct {

	// The destination region that you want to copy snapshots to.
	DestinationRegion *string `json:"destinationRegion,omitempty" tf:"destination_region,omitempty"`

	// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
	GrantName *string `json:"grantName,omitempty" tf:"grant_name,omitempty"`

	// The number of days to retain automated snapshots in the destination region after they are copied from the source region. Defaults to 7.
	RetentionPeriod *int64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

type SnapshotCopyParameters struct {

	// The destination region that you want to copy snapshots to.
	// +kubebuilder:validation:Optional
	DestinationRegion *string `json:"destinationRegion" tf:"destination_region,omitempty"`

	// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
	// +kubebuilder:validation:Optional
	GrantName *string `json:"grantName,omitempty" tf:"grant_name,omitempty"`

	// The number of days to retain automated snapshots in the destination region after they are copied from the source region. Defaults to 7.
	// +kubebuilder:validation:Optional
	RetentionPeriod *int64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
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
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. Provides a Redshift Cluster resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodeType) || (has(self.initProvider) && has(self.initProvider.nodeType))",message="spec.forProvider.nodeType is a required parameter"
	Spec   ClusterSpec   `json:"spec"`
	Status ClusterStatus `json:"status,omitempty"`
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
