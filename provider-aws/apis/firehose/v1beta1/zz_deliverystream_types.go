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

type CloudwatchLoggingOptionsObservation struct {
}

type CloudwatchLoggingOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	LogStreamName *string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type CommonAttributesObservation struct {
}

type CommonAttributesParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DataFormatConversionConfigurationObservation struct {
}

type DataFormatConversionConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	InputFormatConfiguration []InputFormatConfigurationParameters `json:"inputFormatConfiguration" tf:"input_format_configuration,omitempty"`

	// +kubebuilder:validation:Required
	OutputFormatConfiguration []OutputFormatConfigurationParameters `json:"outputFormatConfiguration" tf:"output_format_configuration,omitempty"`

	// +kubebuilder:validation:Required
	SchemaConfiguration []SchemaConfigurationParameters `json:"schemaConfiguration" tf:"schema_configuration,omitempty"`
}

type DeliveryStreamObservation struct {
	ElasticsearchConfiguration []ElasticsearchConfigurationObservation `json:"elasticsearchConfiguration,omitempty" tf:"elasticsearch_configuration,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DeliveryStreamParameters struct {

	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationID *string `json:"destinationId,omitempty" tf:"destination_id,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticsearchConfiguration []ElasticsearchConfigurationParameters `json:"elasticsearchConfiguration,omitempty" tf:"elasticsearch_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	ExtendedS3Configuration []ExtendedS3ConfigurationParameters `json:"extendedS3Configuration,omitempty" tf:"extended_s3_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPEndpointConfiguration []HTTPEndpointConfigurationParameters `json:"httpEndpointConfiguration,omitempty" tf:"http_endpoint_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	KinesisSourceConfiguration []KinesisSourceConfigurationParameters `json:"kinesisSourceConfiguration,omitempty" tf:"kinesis_source_configuration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RedshiftConfiguration []RedshiftConfigurationParameters `json:"redshiftConfiguration,omitempty" tf:"redshift_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	S3Configuration []S3ConfigurationParameters `json:"s3Configuration,omitempty" tf:"s3_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	ServerSideEncryption []ServerSideEncryptionParameters `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// +kubebuilder:validation:Optional
	SplunkConfiguration []SplunkConfigurationParameters `json:"splunkConfiguration,omitempty" tf:"splunk_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type DeserializerObservation struct {
}

type DeserializerParameters struct {

	// +kubebuilder:validation:Optional
	HiveJSONSerDe []HiveJSONSerDeParameters `json:"hiveJsonSerDe,omitempty" tf:"hive_json_ser_de,omitempty"`

	// +kubebuilder:validation:Optional
	OpenXJSONSerDe []OpenXJSONSerDeParameters `json:"openXJsonSerDe,omitempty" tf:"open_x_json_ser_de,omitempty"`
}

type DynamicPartitioningConfigurationObservation struct {
}

type DynamicPartitioningConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RetryDuration *float64 `json:"retryDuration,omitempty" tf:"retry_duration,omitempty"`
}

type ElasticsearchConfigurationObservation struct {
	VPCConfig []VPCConfigObservation `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ElasticsearchConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	BufferingInterval *float64 `json:"bufferingInterval,omitempty" tf:"buffering_interval,omitempty"`

	// +kubebuilder:validation:Optional
	BufferingSize *float64 `json:"bufferingSize,omitempty" tf:"buffering_size,omitempty"`

	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []CloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterEndpoint *string `json:"clusterEndpoint,omitempty" tf:"cluster_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	DomainArn *string `json:"domainArn,omitempty" tf:"domain_arn,omitempty"`

	// +kubebuilder:validation:Required
	IndexName *string `json:"indexName" tf:"index_name,omitempty"`

	// +kubebuilder:validation:Optional
	IndexRotationPeriod *string `json:"indexRotationPeriod,omitempty" tf:"index_rotation_period,omitempty"`

	// +kubebuilder:validation:Optional
	ProcessingConfiguration []ProcessingConfigurationParameters `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	RetryDuration *float64 `json:"retryDuration,omitempty" tf:"retry_duration,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	S3BackupMode *string `json:"s3BackupMode,omitempty" tf:"s3_backup_mode,omitempty"`

	// +kubebuilder:validation:Optional
	TypeName *string `json:"typeName,omitempty" tf:"type_name,omitempty"`

	// +kubebuilder:validation:Optional
	VPCConfig []VPCConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ExtendedS3ConfigurationCloudwatchLoggingOptionsObservation struct {
}

type ExtendedS3ConfigurationCloudwatchLoggingOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	LogStreamName *string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type ExtendedS3ConfigurationObservation struct {
}

type ExtendedS3ConfigurationParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	BucketArn *string `json:"bucketArn,omitempty" tf:"bucket_arn,omitempty"`

	// +kubebuilder:validation:Optional
	BucketArnRef *v1.Reference `json:"bucketArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BucketArnSelector *v1.Selector `json:"bucketArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BufferInterval *float64 `json:"bufferInterval,omitempty" tf:"buffer_interval,omitempty"`

	// +kubebuilder:validation:Optional
	BufferSize *float64 `json:"bufferSize,omitempty" tf:"buffer_size,omitempty"`

	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []ExtendedS3ConfigurationCloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// +kubebuilder:validation:Optional
	CompressionFormat *string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`

	// +kubebuilder:validation:Optional
	DataFormatConversionConfiguration []DataFormatConversionConfigurationParameters `json:"dataFormatConversionConfiguration,omitempty" tf:"data_format_conversion_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	DynamicPartitioningConfiguration []DynamicPartitioningConfigurationParameters `json:"dynamicPartitioningConfiguration,omitempty" tf:"dynamic_partitioning_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	ErrorOutputPrefix *string `json:"errorOutputPrefix,omitempty" tf:"error_output_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	ProcessingConfiguration []ExtendedS3ConfigurationProcessingConfigurationParameters `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	S3BackupConfiguration []S3BackupConfigurationParameters `json:"s3BackupConfiguration,omitempty" tf:"s3_backup_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	S3BackupMode *string `json:"s3BackupMode,omitempty" tf:"s3_backup_mode,omitempty"`
}

type ExtendedS3ConfigurationProcessingConfigurationObservation struct {
}

type ExtendedS3ConfigurationProcessingConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Processors []ProcessingConfigurationProcessorsParameters `json:"processors,omitempty" tf:"processors,omitempty"`
}

type HTTPEndpointConfigurationCloudwatchLoggingOptionsObservation struct {
}

type HTTPEndpointConfigurationCloudwatchLoggingOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	LogStreamName *string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type HTTPEndpointConfigurationObservation struct {
}

type HTTPEndpointConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AccessKeySecretRef *v1.SecretKeySelector `json:"accessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BufferingInterval *float64 `json:"bufferingInterval,omitempty" tf:"buffering_interval,omitempty"`

	// +kubebuilder:validation:Optional
	BufferingSize *float64 `json:"bufferingSize,omitempty" tf:"buffering_size,omitempty"`

	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []HTTPEndpointConfigurationCloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ProcessingConfiguration []HTTPEndpointConfigurationProcessingConfigurationParameters `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	RequestConfiguration []RequestConfigurationParameters `json:"requestConfiguration,omitempty" tf:"request_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	RetryDuration *float64 `json:"retryDuration,omitempty" tf:"retry_duration,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	S3BackupMode *string `json:"s3BackupMode,omitempty" tf:"s3_backup_mode,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type HTTPEndpointConfigurationProcessingConfigurationObservation struct {
}

type HTTPEndpointConfigurationProcessingConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Processors []HTTPEndpointConfigurationProcessingConfigurationProcessorsParameters `json:"processors,omitempty" tf:"processors,omitempty"`
}

type HTTPEndpointConfigurationProcessingConfigurationProcessorsObservation struct {
}

type HTTPEndpointConfigurationProcessingConfigurationProcessorsParameters struct {

	// +kubebuilder:validation:Optional
	Parameters []ProcessingConfigurationProcessorsParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type HiveJSONSerDeObservation struct {
}

type HiveJSONSerDeParameters struct {

	// +kubebuilder:validation:Optional
	TimestampFormats []*string `json:"timestampFormats,omitempty" tf:"timestamp_formats,omitempty"`
}

type InputFormatConfigurationObservation struct {
}

type InputFormatConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Deserializer []DeserializerParameters `json:"deserializer" tf:"deserializer,omitempty"`
}

type KinesisSourceConfigurationObservation struct {
}

type KinesisSourceConfigurationParameters struct {

	// +kubebuilder:validation:Required
	KinesisStreamArn *string `json:"kinesisStreamArn" tf:"kinesis_stream_arn,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type OpenXJSONSerDeObservation struct {
}

type OpenXJSONSerDeParameters struct {

	// +kubebuilder:validation:Optional
	CaseInsensitive *bool `json:"caseInsensitive,omitempty" tf:"case_insensitive,omitempty"`

	// +kubebuilder:validation:Optional
	ColumnToJSONKeyMappings map[string]*string `json:"columnToJsonKeyMappings,omitempty" tf:"column_to_json_key_mappings,omitempty"`

	// +kubebuilder:validation:Optional
	ConvertDotsInJSONKeysToUnderscores *bool `json:"convertDotsInJsonKeysToUnderscores,omitempty" tf:"convert_dots_in_json_keys_to_underscores,omitempty"`
}

type OrcSerDeObservation struct {
}

type OrcSerDeParameters struct {

	// +kubebuilder:validation:Optional
	BlockSizeBytes *float64 `json:"blockSizeBytes,omitempty" tf:"block_size_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	BloomFilterColumns []*string `json:"bloomFilterColumns,omitempty" tf:"bloom_filter_columns,omitempty"`

	// +kubebuilder:validation:Optional
	BloomFilterFalsePositiveProbability *float64 `json:"bloomFilterFalsePositiveProbability,omitempty" tf:"bloom_filter_false_positive_probability,omitempty"`

	// +kubebuilder:validation:Optional
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// +kubebuilder:validation:Optional
	DictionaryKeyThreshold *float64 `json:"dictionaryKeyThreshold,omitempty" tf:"dictionary_key_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	EnablePadding *bool `json:"enablePadding,omitempty" tf:"enable_padding,omitempty"`

	// +kubebuilder:validation:Optional
	FormatVersion *string `json:"formatVersion,omitempty" tf:"format_version,omitempty"`

	// +kubebuilder:validation:Optional
	PaddingTolerance *float64 `json:"paddingTolerance,omitempty" tf:"padding_tolerance,omitempty"`

	// +kubebuilder:validation:Optional
	RowIndexStride *float64 `json:"rowIndexStride,omitempty" tf:"row_index_stride,omitempty"`

	// +kubebuilder:validation:Optional
	StripeSizeBytes *float64 `json:"stripeSizeBytes,omitempty" tf:"stripe_size_bytes,omitempty"`
}

type OutputFormatConfigurationObservation struct {
}

type OutputFormatConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Serializer []SerializerParameters `json:"serializer" tf:"serializer,omitempty"`
}

type ParametersObservation struct {
}

type ParametersParameters struct {

	// +kubebuilder:validation:Required
	ParameterName *string `json:"parameterName" tf:"parameter_name,omitempty"`

	// +kubebuilder:validation:Required
	ParameterValue *string `json:"parameterValue" tf:"parameter_value,omitempty"`
}

type ParquetSerDeObservation struct {
}

type ParquetSerDeParameters struct {

	// +kubebuilder:validation:Optional
	BlockSizeBytes *float64 `json:"blockSizeBytes,omitempty" tf:"block_size_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// +kubebuilder:validation:Optional
	EnableDictionaryCompression *bool `json:"enableDictionaryCompression,omitempty" tf:"enable_dictionary_compression,omitempty"`

	// +kubebuilder:validation:Optional
	MaxPaddingBytes *float64 `json:"maxPaddingBytes,omitempty" tf:"max_padding_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	PageSizeBytes *float64 `json:"pageSizeBytes,omitempty" tf:"page_size_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	WriterVersion *string `json:"writerVersion,omitempty" tf:"writer_version,omitempty"`
}

type ProcessingConfigurationObservation struct {
}

type ProcessingConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Processors []ProcessorsParameters `json:"processors,omitempty" tf:"processors,omitempty"`
}

type ProcessingConfigurationProcessorsObservation struct {
}

type ProcessingConfigurationProcessorsParameters struct {

	// +kubebuilder:validation:Optional
	Parameters []ProcessorsParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ProcessingConfigurationProcessorsParametersObservation struct {
}

type ProcessingConfigurationProcessorsParametersParameters struct {

	// +kubebuilder:validation:Required
	ParameterName *string `json:"parameterName" tf:"parameter_name,omitempty"`

	// +kubebuilder:validation:Required
	ParameterValue *string `json:"parameterValue" tf:"parameter_value,omitempty"`
}

type ProcessorsObservation struct {
}

type ProcessorsParameters struct {

	// +kubebuilder:validation:Optional
	Parameters []ParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ProcessorsParametersObservation struct {
}

type ProcessorsParametersParameters struct {

	// +kubebuilder:validation:Required
	ParameterName *string `json:"parameterName" tf:"parameter_name,omitempty"`

	// +kubebuilder:validation:Required
	ParameterValue *string `json:"parameterValue" tf:"parameter_value,omitempty"`
}

type RedshiftConfigurationCloudwatchLoggingOptionsObservation struct {
}

type RedshiftConfigurationCloudwatchLoggingOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	LogStreamName *string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type RedshiftConfigurationObservation struct {
}

type RedshiftConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []RedshiftConfigurationCloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// +kubebuilder:validation:Required
	ClusterJdbcurl *string `json:"clusterJdbcurl" tf:"cluster_jdbcurl,omitempty"`

	// +kubebuilder:validation:Optional
	CopyOptions *string `json:"copyOptions,omitempty" tf:"copy_options,omitempty"`

	// +kubebuilder:validation:Optional
	DataTableColumns *string `json:"dataTableColumns,omitempty" tf:"data_table_columns,omitempty"`

	// +kubebuilder:validation:Required
	DataTableName *string `json:"dataTableName" tf:"data_table_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	ProcessingConfiguration []RedshiftConfigurationProcessingConfigurationParameters `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	RetryDuration *float64 `json:"retryDuration,omitempty" tf:"retry_duration,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	S3BackupConfiguration []RedshiftConfigurationS3BackupConfigurationParameters `json:"s3BackupConfiguration,omitempty" tf:"s3_backup_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	S3BackupMode *string `json:"s3BackupMode,omitempty" tf:"s3_backup_mode,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type RedshiftConfigurationProcessingConfigurationObservation struct {
}

type RedshiftConfigurationProcessingConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Processors []RedshiftConfigurationProcessingConfigurationProcessorsParameters `json:"processors,omitempty" tf:"processors,omitempty"`
}

type RedshiftConfigurationProcessingConfigurationProcessorsObservation struct {
}

type RedshiftConfigurationProcessingConfigurationProcessorsParameters struct {

	// +kubebuilder:validation:Optional
	Parameters []RedshiftConfigurationProcessingConfigurationProcessorsParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RedshiftConfigurationProcessingConfigurationProcessorsParametersObservation struct {
}

type RedshiftConfigurationProcessingConfigurationProcessorsParametersParameters struct {

	// +kubebuilder:validation:Required
	ParameterName *string `json:"parameterName" tf:"parameter_name,omitempty"`

	// +kubebuilder:validation:Required
	ParameterValue *string `json:"parameterValue" tf:"parameter_value,omitempty"`
}

type RedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptionsObservation struct {
}

type RedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	LogStreamName *string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type RedshiftConfigurationS3BackupConfigurationObservation struct {
}

type RedshiftConfigurationS3BackupConfigurationParameters struct {

	// +kubebuilder:validation:Required
	BucketArn *string `json:"bucketArn" tf:"bucket_arn,omitempty"`

	// +kubebuilder:validation:Optional
	BufferInterval *float64 `json:"bufferInterval,omitempty" tf:"buffer_interval,omitempty"`

	// +kubebuilder:validation:Optional
	BufferSize *float64 `json:"bufferSize,omitempty" tf:"buffer_size,omitempty"`

	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []RedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// +kubebuilder:validation:Optional
	CompressionFormat *string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`

	// +kubebuilder:validation:Optional
	ErrorOutputPrefix *string `json:"errorOutputPrefix,omitempty" tf:"error_output_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type RequestConfigurationObservation struct {
}

type RequestConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CommonAttributes []CommonAttributesParameters `json:"commonAttributes,omitempty" tf:"common_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`
}

type S3BackupConfigurationCloudwatchLoggingOptionsObservation struct {
}

type S3BackupConfigurationCloudwatchLoggingOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	LogStreamName *string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type S3BackupConfigurationObservation struct {
}

type S3BackupConfigurationParameters struct {

	// +kubebuilder:validation:Required
	BucketArn *string `json:"bucketArn" tf:"bucket_arn,omitempty"`

	// +kubebuilder:validation:Optional
	BufferInterval *float64 `json:"bufferInterval,omitempty" tf:"buffer_interval,omitempty"`

	// +kubebuilder:validation:Optional
	BufferSize *float64 `json:"bufferSize,omitempty" tf:"buffer_size,omitempty"`

	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []S3BackupConfigurationCloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// +kubebuilder:validation:Optional
	CompressionFormat *string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`

	// +kubebuilder:validation:Optional
	ErrorOutputPrefix *string `json:"errorOutputPrefix,omitempty" tf:"error_output_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type S3ConfigurationCloudwatchLoggingOptionsObservation struct {
}

type S3ConfigurationCloudwatchLoggingOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	LogStreamName *string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type S3ConfigurationObservation struct {
}

type S3ConfigurationParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	BucketArn *string `json:"bucketArn,omitempty" tf:"bucket_arn,omitempty"`

	// +kubebuilder:validation:Optional
	BucketArnRef *v1.Reference `json:"bucketArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BucketArnSelector *v1.Selector `json:"bucketArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BufferInterval *float64 `json:"bufferInterval,omitempty" tf:"buffer_interval,omitempty"`

	// +kubebuilder:validation:Optional
	BufferSize *float64 `json:"bufferSize,omitempty" tf:"buffer_size,omitempty"`

	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []S3ConfigurationCloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// +kubebuilder:validation:Optional
	CompressionFormat *string `json:"compressionFormat,omitempty" tf:"compression_format,omitempty"`

	// +kubebuilder:validation:Optional
	ErrorOutputPrefix *string `json:"errorOutputPrefix,omitempty" tf:"error_output_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

type SchemaConfigurationObservation struct {
}

type SchemaConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`

	// +kubebuilder:validation:Optional
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type SerializerObservation struct {
}

type SerializerParameters struct {

	// +kubebuilder:validation:Optional
	OrcSerDe []OrcSerDeParameters `json:"orcSerDe,omitempty" tf:"orc_ser_de,omitempty"`

	// +kubebuilder:validation:Optional
	ParquetSerDe []ParquetSerDeParameters `json:"parquetSerDe,omitempty" tf:"parquet_ser_de,omitempty"`
}

type ServerSideEncryptionObservation struct {
}

type ServerSideEncryptionParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	KeyArn *string `json:"keyArn,omitempty" tf:"key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`
}

type SplunkConfigurationCloudwatchLoggingOptionsObservation struct {
}

type SplunkConfigurationCloudwatchLoggingOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	LogStreamName *string `json:"logStreamName,omitempty" tf:"log_stream_name,omitempty"`
}

type SplunkConfigurationObservation struct {
}

type SplunkConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []SplunkConfigurationCloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// +kubebuilder:validation:Optional
	HecAcknowledgmentTimeout *float64 `json:"hecAcknowledgmentTimeout,omitempty" tf:"hec_acknowledgment_timeout,omitempty"`

	// +kubebuilder:validation:Required
	HecEndpoint *string `json:"hecEndpoint" tf:"hec_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	HecEndpointType *string `json:"hecEndpointType,omitempty" tf:"hec_endpoint_type,omitempty"`

	// +kubebuilder:validation:Required
	HecToken *string `json:"hecToken" tf:"hec_token,omitempty"`

	// +kubebuilder:validation:Optional
	ProcessingConfiguration []SplunkConfigurationProcessingConfigurationParameters `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	RetryDuration *float64 `json:"retryDuration,omitempty" tf:"retry_duration,omitempty"`

	// +kubebuilder:validation:Optional
	S3BackupMode *string `json:"s3BackupMode,omitempty" tf:"s3_backup_mode,omitempty"`
}

type SplunkConfigurationProcessingConfigurationObservation struct {
}

type SplunkConfigurationProcessingConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Processors []SplunkConfigurationProcessingConfigurationProcessorsParameters `json:"processors,omitempty" tf:"processors,omitempty"`
}

type SplunkConfigurationProcessingConfigurationProcessorsObservation struct {
}

type SplunkConfigurationProcessingConfigurationProcessorsParameters struct {

	// +kubebuilder:validation:Optional
	Parameters []SplunkConfigurationProcessingConfigurationProcessorsParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SplunkConfigurationProcessingConfigurationProcessorsParametersObservation struct {
}

type SplunkConfigurationProcessingConfigurationProcessorsParametersParameters struct {

	// +kubebuilder:validation:Required
	ParameterName *string `json:"parameterName" tf:"parameter_name,omitempty"`

	// +kubebuilder:validation:Required
	ParameterValue *string `json:"parameterValue" tf:"parameter_value,omitempty"`
}

type VPCConfigObservation struct {
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCConfigParameters struct {

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Required
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`
}

// DeliveryStreamSpec defines the desired state of DeliveryStream
type DeliveryStreamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeliveryStreamParameters `json:"forProvider"`
}

// DeliveryStreamStatus defines the observed state of DeliveryStream.
type DeliveryStreamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeliveryStreamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DeliveryStream is the Schema for the DeliveryStreams API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DeliveryStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeliveryStreamSpec   `json:"spec"`
	Status            DeliveryStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeliveryStreamList contains a list of DeliveryStreams
type DeliveryStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeliveryStream `json:"items"`
}

// Repository type metadata.
var (
	DeliveryStream_Kind             = "DeliveryStream"
	DeliveryStream_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeliveryStream_Kind}.String()
	DeliveryStream_KindAPIVersion   = DeliveryStream_Kind + "." + CRDGroupVersion.String()
	DeliveryStream_GroupVersionKind = CRDGroupVersion.WithKind(DeliveryStream_Kind)
)

func init() {
	SchemeBuilder.Register(&DeliveryStream{}, &DeliveryStreamList{})
}
