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

type AcceleratorCountObservation struct {
}

type AcceleratorCountParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type AcceleratorTotalMemoryMibObservation struct {
}

type AcceleratorTotalMemoryMibParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type BaselineEBSBandwidthMbpsObservation struct {
}

type BaselineEBSBandwidthMbpsParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type BlockDeviceMappingsObservation struct {
}

type BlockDeviceMappingsParameters struct {

	// The name of the device to mount.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Configure EBS volume properties.
	// +kubebuilder:validation:Optional
	EBS []EBSParameters `json:"ebs,omitempty" tf:"ebs,omitempty"`

	// Suppresses the specified device included in the AMI's block device mapping.
	// +kubebuilder:validation:Optional
	NoDevice *string `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// The Instance Store Device
	// Name
	// .
	// +kubebuilder:validation:Optional
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type CPUOptionsObservation struct {
}

type CPUOptionsParameters struct {

	// The number of CPU cores for the instance.
	// +kubebuilder:validation:Optional
	CoreCount *float64 `json:"coreCount,omitempty" tf:"core_count,omitempty"`

	// The number of threads per CPU core. To disable Intel Hyper-Threading Technology for the instance, specify a value of 1.
	// Otherwise, specify the default value of 2.
	// +kubebuilder:validation:Optional
	ThreadsPerCore *float64 `json:"threadsPerCore,omitempty" tf:"threads_per_core,omitempty"`
}

type CapacityReservationSpecificationCapacityReservationTargetObservation struct {
}

type CapacityReservationSpecificationCapacityReservationTargetParameters struct {

	// The ID of the Capacity Reservation in which to run the instance.
	// +kubebuilder:validation:Optional
	CapacityReservationID *string `json:"capacityReservationId,omitempty" tf:"capacity_reservation_id,omitempty"`

	// The ARN of the Capacity Reservation resource group in which to run the instance.
	// +kubebuilder:validation:Optional
	CapacityReservationResourceGroupArn *string `json:"capacityReservationResourceGroupArn,omitempty" tf:"capacity_reservation_resource_group_arn,omitempty"`
}

type EBSObservation struct {
}

type EBSParameters struct {

	// +kubebuilder:validation:Optional
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Enables EBS encryption
	// on the volume . Cannot be used with snapshot_id.
	// +kubebuilder:validation:Optional
	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The amount of provisioned
	// IOPS.
	// This must be set with a volume_type of "io1/io2".
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The ARN of the AWS Key Management Service  customer master key  to use when creating the encrypted volume.
	// encrypted must be set to true when this is set.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The Snapshot ID to mount.
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// The throughput to provision for a gp3 volume in MiB/s , with a maximum of 1,000 MiB/s.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// The size of the volume in gigabytes.
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// The volume type. Can be standard, gp2, gp3, io1, io2, sc1 or st1 .
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type ElasticGpuSpecificationsObservation struct {
}

type ElasticGpuSpecificationsParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ElasticInferenceAcceleratorObservation struct {
}

type ElasticInferenceAcceleratorParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type HibernationOptionsObservation struct {
}

type HibernationOptionsParameters struct {

	// If set to true, the launched EC2 instance will hibernation enabled.
	// +kubebuilder:validation:Required
	Configured *bool `json:"configured" tf:"configured,omitempty"`
}

type IAMInstanceProfileObservation struct {
}

type IAMInstanceProfileParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.InstanceProfile
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.InstanceProfile
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`
}

type InstanceMarketOptionsObservation struct {
}

type InstanceMarketOptionsParameters struct {

	// The market type. Can be spot.
	// +kubebuilder:validation:Optional
	MarketType *string `json:"marketType,omitempty" tf:"market_type,omitempty"`

	// The options for Spot Instance
	// +kubebuilder:validation:Optional
	SpotOptions []SpotOptionsParameters `json:"spotOptions,omitempty" tf:"spot_options,omitempty"`
}

type InstanceRequirementsObservation struct {
}

type InstanceRequirementsParameters struct {

	// Block describing the minimum and maximum number of accelerators . Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	AcceleratorCount []AcceleratorCountParameters `json:"acceleratorCount,omitempty" tf:"accelerator_count,omitempty"`

	// List of accelerator manufacturer names. Default is any manufacturer.
	// +kubebuilder:validation:Optional
	AcceleratorManufacturers []*string `json:"acceleratorManufacturers,omitempty" tf:"accelerator_manufacturers,omitempty"`

	// List of accelerator names. Default is any acclerator.
	// +kubebuilder:validation:Optional
	AcceleratorNames []*string `json:"acceleratorNames,omitempty" tf:"accelerator_names,omitempty"`

	// Block describing the minimum and maximum total memory of the accelerators. Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	AcceleratorTotalMemoryMib []AcceleratorTotalMemoryMibParameters `json:"acceleratorTotalMemoryMib,omitempty" tf:"accelerator_total_memory_mib,omitempty"`

	// List of accelerator types. Default is any accelerator type.
	// +kubebuilder:validation:Optional
	AcceleratorTypes []*string `json:"acceleratorTypes,omitempty" tf:"accelerator_types,omitempty"`

	// Indicate whether bare metal instace types should be included, excluded, or required. Default is excluded.
	// +kubebuilder:validation:Optional
	BareMetal *string `json:"bareMetal,omitempty" tf:"bare_metal,omitempty"`

	// Block describing the minimum and maximum baseline EBS bandwidth, in Mbps. Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	BaselineEBSBandwidthMbps []BaselineEBSBandwidthMbpsParameters `json:"baselineEbsBandwidthMbps,omitempty" tf:"baseline_ebs_bandwidth_mbps,omitempty"`

	// Indicate whether burstable performance instance types should be included, excluded, or required. Default is excluded.
	// +kubebuilder:validation:Optional
	BurstablePerformance *string `json:"burstablePerformance,omitempty" tf:"burstable_performance,omitempty"`

	// List of CPU manufacturer names. Default is any manufacturer.
	// +kubebuilder:validation:Optional
	CPUManufacturers []*string `json:"cpuManufacturers,omitempty" tf:"cpu_manufacturers,omitempty"`

	// List of instance types to exclude. You can use strings with one or more wild cards, represented by an asterisk . The following are examples: c5*, m5a.*, r*, *3*. For example, if you specify c5*, you are excluding the entire C5 instance family, which includes all C5a and C5n instance types. If you specify m5a.*, you are excluding all the M5a instance types, but not the M5n instance types. Maximum of 400 entries in the list; each entry is limited to 30 characters. Default is no excluded instance types.
	// +kubebuilder:validation:Optional
	ExcludedInstanceTypes []*string `json:"excludedInstanceTypes,omitempty" tf:"excluded_instance_types,omitempty"`

	// List of instance generation names. Default is any generation.
	// +kubebuilder:validation:Optional
	InstanceGenerations []*string `json:"instanceGenerations,omitempty" tf:"instance_generations,omitempty"`

	// Indicate whether instance types with local storage volumes are included, excluded, or required. Default is included.
	// +kubebuilder:validation:Optional
	LocalStorage *string `json:"localStorage,omitempty" tf:"local_storage,omitempty"`

	// List of local storage type names. Default any storage type.
	// +kubebuilder:validation:Optional
	LocalStorageTypes []*string `json:"localStorageTypes,omitempty" tf:"local_storage_types,omitempty"`

	// Block describing the minimum and maximum amount of memory  per vCPU. Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	MemoryGibPerVcpu []MemoryGibPerVcpuParameters `json:"memoryGibPerVcpu,omitempty" tf:"memory_gib_per_vcpu,omitempty"`

	// Block describing the minimum and maximum amount of memory . Default is no maximum.
	// +kubebuilder:validation:Required
	MemoryMib []MemoryMibParameters `json:"memoryMib" tf:"memory_mib,omitempty"`

	// Block describing the minimum and maximum number of network interfaces. Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	NetworkInterfaceCount []NetworkInterfaceCountParameters `json:"networkInterfaceCount,omitempty" tf:"network_interface_count,omitempty"`

	// The price protection threshold for On-Demand Instances. This is the maximum you’ll pay for an On-Demand Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as 999999. Default is 20.
	// +kubebuilder:validation:Optional
	OnDemandMaxPricePercentageOverLowestPrice *float64 `json:"onDemandMaxPricePercentageOverLowestPrice,omitempty" tf:"on_demand_max_price_percentage_over_lowest_price,omitempty"`

	// Indicate whether instance types must support On-Demand Instance Hibernation, either true or false. Default is false.
	// +kubebuilder:validation:Optional
	RequireHibernateSupport *bool `json:"requireHibernateSupport,omitempty" tf:"require_hibernate_support,omitempty"`

	// The price protection threshold for Spot Instances. This is the maximum you’ll pay for a Spot Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as 999999. Default is 100.
	// +kubebuilder:validation:Optional
	SpotMaxPricePercentageOverLowestPrice *float64 `json:"spotMaxPricePercentageOverLowestPrice,omitempty" tf:"spot_max_price_percentage_over_lowest_price,omitempty"`

	// Block describing the minimum and maximum total local storage . Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	TotalLocalStorageGb []TotalLocalStorageGbParameters `json:"totalLocalStorageGb,omitempty" tf:"total_local_storage_gb,omitempty"`

	// Block describing the minimum and maximum number of vCPUs. Default is no maximum.
	// +kubebuilder:validation:Required
	VcpuCount []VcpuCountParameters `json:"vcpuCount" tf:"vcpu_count,omitempty"`
}

type LaunchTemplateCapacityReservationSpecificationObservation struct {
}

type LaunchTemplateCapacityReservationSpecificationParameters struct {

	// Indicates the instance's Capacity Reservation preferences. Can be open or none. .
	// +kubebuilder:validation:Optional
	CapacityReservationPreference *string `json:"capacityReservationPreference,omitempty" tf:"capacity_reservation_preference,omitempty"`

	// Used to target a specific Capacity Reservation:
	// +kubebuilder:validation:Optional
	CapacityReservationTarget []CapacityReservationSpecificationCapacityReservationTargetParameters `json:"capacityReservationTarget,omitempty" tf:"capacity_reservation_target,omitempty"`
}

type LaunchTemplateCreditSpecificationObservation struct {
}

type LaunchTemplateCreditSpecificationParameters struct {

	// The credit option for CPU usage. Can be "standard" or "unlimited". T3 instances are launched as unlimited by default. T2 instances are launched as standard by default.
	// +kubebuilder:validation:Optional
	CPUCredits *string `json:"cpuCredits,omitempty" tf:"cpu_credits,omitempty"`
}

type LaunchTemplateEnclaveOptionsObservation struct {
}

type LaunchTemplateEnclaveOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LaunchTemplateMaintenanceOptionsObservation struct {
}

type LaunchTemplateMaintenanceOptionsParameters struct {

	// Disables the automatic recovery behavior of your instance or sets it to default. Can be "default" or "disabled". See Recover your instance for more details.
	// +kubebuilder:validation:Optional
	AutoRecovery *string `json:"autoRecovery,omitempty" tf:"auto_recovery,omitempty"`
}

type LaunchTemplateMetadataOptionsObservation struct {
}

type LaunchTemplateMetadataOptionsParameters struct {

	// Whether the metadata service is available. Can be "enabled" or "disabled". .
	// +kubebuilder:validation:Optional
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint,omitempty"`

	// Enables or disables the IPv6 endpoint for the instance metadata service. .
	// +kubebuilder:validation:Optional
	HTTPProtocolIPv6 *string `json:"httpProtocolIpv6,omitempty" tf:"http_protocol_ipv6,omitempty"`

	// The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel. Can be an integer from 1 to 64. .
	// +kubebuilder:validation:Optional
	HTTPPutResponseHopLimit *float64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit,omitempty"`

	// Whether or not the metadata service requires session tokens, also referred to as Instance Metadata Service Version 2 . Can be "optional" or "required". .
	// +kubebuilder:validation:Optional
	HTTPTokens *string `json:"httpTokens,omitempty" tf:"http_tokens,omitempty"`

	// Enables or disables access to instance tags from the instance metadata service. .
	// +kubebuilder:validation:Optional
	InstanceMetadataTags *string `json:"instanceMetadataTags,omitempty" tf:"instance_metadata_tags,omitempty"`
}

type LaunchTemplateObservation_2 struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the launch template.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The latest version of the launch template.
	LatestVersion *float64 `json:"latestVersion,omitempty" tf:"latest_version,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LaunchTemplateParameters_2 struct {

	// Specify volumes to attach to the instance besides the volumes specified by the AMI.
	// See Block Devices below for details.
	// +kubebuilder:validation:Optional
	BlockDeviceMappings []BlockDeviceMappingsParameters `json:"blockDeviceMappings,omitempty" tf:"block_device_mappings,omitempty"`

	// The CPU options for the instance. See CPU Options below for more details.
	// +kubebuilder:validation:Optional
	CPUOptions []CPUOptionsParameters `json:"cpuOptions,omitempty" tf:"cpu_options,omitempty"`

	// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
	// +kubebuilder:validation:Optional
	CapacityReservationSpecification []LaunchTemplateCapacityReservationSpecificationParameters `json:"capacityReservationSpecification,omitempty" tf:"capacity_reservation_specification,omitempty"`

	// Customize the credit specification of the instance. See Credit
	// Specification below for more details.
	// +kubebuilder:validation:Optional
	CreditSpecification []LaunchTemplateCreditSpecificationParameters `json:"creditSpecification,omitempty" tf:"credit_specification,omitempty"`

	// Default Version of the launch template.
	// +kubebuilder:validation:Optional
	DefaultVersion *float64 `json:"defaultVersion,omitempty" tf:"default_version,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If true, enables EC2 Instance
	// Termination Protection
	// +kubebuilder:validation:Optional
	DisableAPITermination *bool `json:"disableApiTermination,omitempty" tf:"disable_api_termination,omitempty"`

	// If true, the launched EC2 instance will be EBS-optimized.
	// +kubebuilder:validation:Optional
	EBSOptimized *string `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`

	// The elastic GPU to attach to the instance. See Elastic GPU
	// below for more details.
	// +kubebuilder:validation:Optional
	ElasticGpuSpecifications []ElasticGpuSpecificationsParameters `json:"elasticGpuSpecifications,omitempty" tf:"elastic_gpu_specifications,omitempty"`

	// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
	// +kubebuilder:validation:Optional
	ElasticInferenceAccelerator []ElasticInferenceAcceleratorParameters `json:"elasticInferenceAccelerator,omitempty" tf:"elastic_inference_accelerator,omitempty"`

	// Enable Nitro Enclaves on launched instances. See Enclave Options below for more details.
	// +kubebuilder:validation:Optional
	EnclaveOptions []LaunchTemplateEnclaveOptionsParameters `json:"enclaveOptions,omitempty" tf:"enclave_options,omitempty"`

	// The hibernation options for the instance. See Hibernation Options below for more details.
	// +kubebuilder:validation:Optional
	HibernationOptions []HibernationOptionsParameters `json:"hibernationOptions,omitempty" tf:"hibernation_options,omitempty"`

	// The IAM Instance Profile to launch the instance with. See Instance Profile
	// below for more details.
	// +kubebuilder:validation:Optional
	IAMInstanceProfile []IAMInstanceProfileParameters `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`

	// The AMI from which to launch the instance.
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// Shutdown behavior for the instance. Can be stop or terminate.
	// .
	// +kubebuilder:validation:Optional
	InstanceInitiatedShutdownBehavior *string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior,omitempty"`

	// The market  option for the instance. See Market Options
	// below for details.
	// +kubebuilder:validation:Optional
	InstanceMarketOptions []InstanceMarketOptionsParameters `json:"instanceMarketOptions,omitempty" tf:"instance_market_options,omitempty"`

	// The attribute requirements for the type of instance. If present then instance_type cannot be present.
	// +kubebuilder:validation:Optional
	InstanceRequirements []InstanceRequirementsParameters `json:"instanceRequirements,omitempty" tf:"instance_requirements,omitempty"`

	// The type of the instance. If present then instance_requirements cannot be present.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The kernel ID.
	// +kubebuilder:validation:Optional
	KernelID *string `json:"kernelId,omitempty" tf:"kernel_id,omitempty"`

	// The key name to use for the instance.
	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// A list of license specifications to associate with. See License Specification below for more details.
	// +kubebuilder:validation:Optional
	LicenseSpecification []LicenseSpecificationParameters `json:"licenseSpecification,omitempty" tf:"license_specification,omitempty"`

	// The maintenance options for the instance. See Maintenance Options below for more details.
	// +kubebuilder:validation:Optional
	MaintenanceOptions []LaunchTemplateMaintenanceOptionsParameters `json:"maintenanceOptions,omitempty" tf:"maintenance_options,omitempty"`

	// Customize the metadata options for the instance. See Metadata Options below for more details.
	// +kubebuilder:validation:Optional
	MetadataOptions []LaunchTemplateMetadataOptionsParameters `json:"metadataOptions,omitempty" tf:"metadata_options,omitempty"`

	// The monitoring option for the instance. See Monitoring below for more details.
	// +kubebuilder:validation:Optional
	Monitoring []MonitoringParameters `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// Customize network interfaces to be attached at instance boot time. See Network
	// Interfaces below for more details.
	// +kubebuilder:validation:Optional
	NetworkInterfaces []NetworkInterfacesParameters `json:"networkInterfaces,omitempty" tf:"network_interfaces,omitempty"`

	// The placement of the instance. See Placement below for more details.
	// +kubebuilder:validation:Optional
	Placement []PlacementParameters `json:"placement,omitempty" tf:"placement,omitempty"`

	// The options for the instance hostname. The default values are inherited from the subnet. See Private DNS Name Options below for more details.
	// +kubebuilder:validation:Optional
	PrivateDNSNameOptions []PrivateDNSNameOptionsParameters `json:"privateDnsNameOptions,omitempty" tf:"private_dns_name_options,omitempty"`

	// The ID of the RAM disk.
	// +kubebuilder:validation:Optional
	RAMDiskID *string `json:"ramDiskId,omitempty" tf:"ram_disk_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupNameRefs []v1.Reference `json:"securityGroupNameRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupNameSelector *v1.Selector `json:"securityGroupNameSelector,omitempty" tf:"-"`

	// A list of security group names to associate with. If you are creating Instances in a VPC, use
	// vpc_security_group_ids instead.
	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupNameRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupNameSelector
	// +kubebuilder:validation:Optional
	SecurityGroupNames []*string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`

	// The tags to apply to the resources during launch. See Tag Specifications below for more details.
	// +kubebuilder:validation:Optional
	TagSpecifications []TagSpecificationsParameters `json:"tagSpecifications,omitempty" tf:"tag_specifications,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether to update Default Version each update. Conflicts with default_version.
	// +kubebuilder:validation:Optional
	UpdateDefaultVersion *bool `json:"updateDefaultVersion,omitempty" tf:"update_default_version,omitempty"`

	// The base64-encoded user data to provide when launching the instance.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// A list of security group IDs to associate with. Conflicts with network_interfaces.security_groups
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type LicenseSpecificationObservation struct {
}

type LicenseSpecificationParameters struct {

	// ARN of the license configuration.
	// +kubebuilder:validation:Required
	LicenseConfigurationArn *string `json:"licenseConfigurationArn" tf:"license_configuration_arn,omitempty"`
}

type MemoryGibPerVcpuObservation struct {
}

type MemoryGibPerVcpuParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type MemoryMibObservation struct {
}

type MemoryMibParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type MonitoringObservation struct {
}

type MonitoringParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type NetworkInterfaceCountObservation struct {
}

type NetworkInterfaceCountParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type NetworkInterfacesObservation struct {
}

type NetworkInterfacesParameters struct {

	// Associate a Carrier IP address with eth0 for a new network interface. Use this option when you launch an instance in a Wavelength Zone and want to associate a Carrier IP address with the network interface. Boolean value.
	// +kubebuilder:validation:Optional
	AssociateCarrierIPAddress *string `json:"associateCarrierIpAddress,omitempty" tf:"associate_carrier_ip_address,omitempty"`

	// Associate a public ip address with the network interface.  Boolean value.
	// +kubebuilder:validation:Optional
	AssociatePublicIPAddress *string `json:"associatePublicIpAddress,omitempty" tf:"associate_public_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integer index of the network interface attachment.
	// +kubebuilder:validation:Optional
	DeviceIndex *float64 `json:"deviceIndex,omitempty" tf:"device_index,omitempty"`

	// The number of secondary private IPv4 addresses to assign to a network interface. Conflicts with ipv4_addresses
	// +kubebuilder:validation:Optional
	IPv4AddressCount *float64 `json:"ipv4AddressCount,omitempty" tf:"ipv4_address_count,omitempty"`

	// One or more private IPv4 addresses to associate. Conflicts with ipv4_address_count
	// +kubebuilder:validation:Optional
	IPv4Addresses []*string `json:"ipv4Addresses,omitempty" tf:"ipv4_addresses,omitempty"`

	// The number of IPv4 prefixes to be automatically assigned to the network interface. Conflicts with ipv4_prefixes
	// +kubebuilder:validation:Optional
	IPv4PrefixCount *float64 `json:"ipv4PrefixCount,omitempty" tf:"ipv4_prefix_count,omitempty"`

	// One or more IPv4 prefixes to be assigned to the network interface. Conflicts with ipv4_prefix_count
	// +kubebuilder:validation:Optional
	IPv4Prefixes []*string `json:"ipv4Prefixes,omitempty" tf:"ipv4_prefixes,omitempty"`

	// The number of IPv6 addresses to assign to a network interface. Conflicts with ipv6_addresses
	// +kubebuilder:validation:Optional
	IPv6AddressCount *float64 `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count,omitempty"`

	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Conflicts with ipv6_address_count
	// +kubebuilder:validation:Optional
	IPv6Addresses []*string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses,omitempty"`

	// The number of IPv6 prefixes to be automatically assigned to the network interface. Conflicts with ipv6_prefixes
	// +kubebuilder:validation:Optional
	IPv6PrefixCount *float64 `json:"ipv6PrefixCount,omitempty" tf:"ipv6_prefix_count,omitempty"`

	// One or more IPv6 prefixes to be assigned to the network interface. Conflicts with ipv6_prefix_count
	// +kubebuilder:validation:Optional
	IPv6Prefixes []*string `json:"ipv6Prefixes,omitempty" tf:"ipv6_prefixes,omitempty"`

	// The type of network interface. To create an Elastic Fabric Adapter , specify efa.
	// +kubebuilder:validation:Optional
	InterfaceType *string `json:"interfaceType,omitempty" tf:"interface_type,omitempty"`

	// The index of the network card. Some instance types support multiple network cards. The primary network interface must be assigned to network card index 0. The default is network card index 0.
	// +kubebuilder:validation:Optional
	NetworkCardIndex *float64 `json:"networkCardIndex,omitempty" tf:"network_card_index,omitempty"`

	// The ID of the network interface to attach.
	// +crossplane:generate:reference:type=NetworkInterface
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`

	// The primary private IPv4 address.
	// +kubebuilder:validation:Optional
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// A list of security group IDs to associate.
	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The VPC Subnet ID to associate.
	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type PlacementObservation struct {
}

type PlacementParameters struct {

	// The affinity setting for an instance on a Dedicated Host.
	// +kubebuilder:validation:Optional
	Affinity *string `json:"affinity,omitempty" tf:"affinity,omitempty"`

	// The Availability Zone for the instance.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The name of the placement group for the instance.
	// +kubebuilder:validation:Optional
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// The ID of the Dedicated Host for the instance.
	// +kubebuilder:validation:Optional
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// The ARN of the Host Resource Group in which to launch instances.
	// +kubebuilder:validation:Optional
	HostResourceGroupArn *string `json:"hostResourceGroupArn,omitempty" tf:"host_resource_group_arn,omitempty"`

	// The number of the partition the instance should launch in. Valid only if the placement group strategy is set to partition.
	// +kubebuilder:validation:Optional
	PartitionNumber *float64 `json:"partitionNumber,omitempty" tf:"partition_number,omitempty"`

	// Reserved for future use.
	// +kubebuilder:validation:Optional
	SpreadDomain *string `json:"spreadDomain,omitempty" tf:"spread_domain,omitempty"`

	// The tenancy of the instance . Can be default, dedicated, or host.
	// +kubebuilder:validation:Optional
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`
}

type PrivateDNSNameOptionsObservation struct {
}

type PrivateDNSNameOptionsParameters struct {

	// Indicates whether to respond to DNS queries for instance hostnames with DNS A records.
	// +kubebuilder:validation:Optional
	EnableResourceNameDNSARecord *bool `json:"enableResourceNameDnsARecord,omitempty" tf:"enable_resource_name_dns_a_record,omitempty"`

	// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.
	// +kubebuilder:validation:Optional
	EnableResourceNameDNSAaaaRecord *bool `json:"enableResourceNameDnsAaaaRecord,omitempty" tf:"enable_resource_name_dns_aaaa_record,omitempty"`

	// The type of hostname for Amazon EC2 instances. For IPv4 only subnets, an instance DNS name must be based on the instance IPv4 address. For IPv6 native subnets, an instance DNS name must be based on the instance ID. For dual-stack subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: ip-name and resource-name.
	// +kubebuilder:validation:Optional
	HostnameType *string `json:"hostnameType,omitempty" tf:"hostname_type,omitempty"`
}

type SpotOptionsObservation struct {
}

type SpotOptionsParameters struct {

	// The required duration in minutes. This value must be a multiple of 60.
	// +kubebuilder:validation:Optional
	BlockDurationMinutes *float64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes,omitempty"`

	// The behavior when a Spot Instance is interrupted. Can be hibernate,
	// stop, or terminate. .
	// +kubebuilder:validation:Optional
	InstanceInterruptionBehavior *string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior,omitempty"`

	// The maximum hourly price you're willing to pay for the Spot Instances.
	// +kubebuilder:validation:Optional
	MaxPrice *string `json:"maxPrice,omitempty" tf:"max_price,omitempty"`

	// The Spot Instance request type. Can be one-time, or persistent.
	// +kubebuilder:validation:Optional
	SpotInstanceType *string `json:"spotInstanceType,omitempty" tf:"spot_instance_type,omitempty"`

	// The end date of the request.
	// +kubebuilder:validation:Optional
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
}

type TagSpecificationsObservation struct {
}

type TagSpecificationsParameters struct {

	// The type of resource to tag.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TotalLocalStorageGbObservation struct {
}

type TotalLocalStorageGbParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type VcpuCountObservation struct {
}

type VcpuCountParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

// LaunchTemplateSpec defines the desired state of LaunchTemplate
type LaunchTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LaunchTemplateParameters_2 `json:"forProvider"`
}

// LaunchTemplateStatus defines the observed state of LaunchTemplate.
type LaunchTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LaunchTemplateObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LaunchTemplate is the Schema for the LaunchTemplates API. Provides an EC2 launch template resource. Can be used to create instances or auto scaling groups.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LaunchTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LaunchTemplateSpec   `json:"spec"`
	Status            LaunchTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LaunchTemplateList contains a list of LaunchTemplates
type LaunchTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LaunchTemplate `json:"items"`
}

// Repository type metadata.
var (
	LaunchTemplate_Kind             = "LaunchTemplate"
	LaunchTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LaunchTemplate_Kind}.String()
	LaunchTemplate_KindAPIVersion   = LaunchTemplate_Kind + "." + CRDGroupVersion.String()
	LaunchTemplate_GroupVersionKind = CRDGroupVersion.WithKind(LaunchTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&LaunchTemplate{}, &LaunchTemplateList{})
}
