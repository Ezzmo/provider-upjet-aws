//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaluateOnExitInitParameters) DeepCopyInto(out *EvaluateOnExitInitParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.OnExitCode != nil {
		in, out := &in.OnExitCode, &out.OnExitCode
		*out = new(string)
		**out = **in
	}
	if in.OnReason != nil {
		in, out := &in.OnReason, &out.OnReason
		*out = new(string)
		**out = **in
	}
	if in.OnStatusReason != nil {
		in, out := &in.OnStatusReason, &out.OnStatusReason
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaluateOnExitInitParameters.
func (in *EvaluateOnExitInitParameters) DeepCopy() *EvaluateOnExitInitParameters {
	if in == nil {
		return nil
	}
	out := new(EvaluateOnExitInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaluateOnExitObservation) DeepCopyInto(out *EvaluateOnExitObservation) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.OnExitCode != nil {
		in, out := &in.OnExitCode, &out.OnExitCode
		*out = new(string)
		**out = **in
	}
	if in.OnReason != nil {
		in, out := &in.OnReason, &out.OnReason
		*out = new(string)
		**out = **in
	}
	if in.OnStatusReason != nil {
		in, out := &in.OnStatusReason, &out.OnStatusReason
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaluateOnExitObservation.
func (in *EvaluateOnExitObservation) DeepCopy() *EvaluateOnExitObservation {
	if in == nil {
		return nil
	}
	out := new(EvaluateOnExitObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaluateOnExitParameters) DeepCopyInto(out *EvaluateOnExitParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.OnExitCode != nil {
		in, out := &in.OnExitCode, &out.OnExitCode
		*out = new(string)
		**out = **in
	}
	if in.OnReason != nil {
		in, out := &in.OnReason, &out.OnReason
		*out = new(string)
		**out = **in
	}
	if in.OnStatusReason != nil {
		in, out := &in.OnStatusReason, &out.OnStatusReason
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaluateOnExitParameters.
func (in *EvaluateOnExitParameters) DeepCopy() *EvaluateOnExitParameters {
	if in == nil {
		return nil
	}
	out := new(EvaluateOnExitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FairSharePolicyInitParameters) DeepCopyInto(out *FairSharePolicyInitParameters) {
	*out = *in
	if in.ComputeReservation != nil {
		in, out := &in.ComputeReservation, &out.ComputeReservation
		*out = new(int64)
		**out = **in
	}
	if in.ShareDecaySeconds != nil {
		in, out := &in.ShareDecaySeconds, &out.ShareDecaySeconds
		*out = new(int64)
		**out = **in
	}
	if in.ShareDistribution != nil {
		in, out := &in.ShareDistribution, &out.ShareDistribution
		*out = make([]ShareDistributionInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FairSharePolicyInitParameters.
func (in *FairSharePolicyInitParameters) DeepCopy() *FairSharePolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(FairSharePolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FairSharePolicyObservation) DeepCopyInto(out *FairSharePolicyObservation) {
	*out = *in
	if in.ComputeReservation != nil {
		in, out := &in.ComputeReservation, &out.ComputeReservation
		*out = new(int64)
		**out = **in
	}
	if in.ShareDecaySeconds != nil {
		in, out := &in.ShareDecaySeconds, &out.ShareDecaySeconds
		*out = new(int64)
		**out = **in
	}
	if in.ShareDistribution != nil {
		in, out := &in.ShareDistribution, &out.ShareDistribution
		*out = make([]ShareDistributionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FairSharePolicyObservation.
func (in *FairSharePolicyObservation) DeepCopy() *FairSharePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(FairSharePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FairSharePolicyParameters) DeepCopyInto(out *FairSharePolicyParameters) {
	*out = *in
	if in.ComputeReservation != nil {
		in, out := &in.ComputeReservation, &out.ComputeReservation
		*out = new(int64)
		**out = **in
	}
	if in.ShareDecaySeconds != nil {
		in, out := &in.ShareDecaySeconds, &out.ShareDecaySeconds
		*out = new(int64)
		**out = **in
	}
	if in.ShareDistribution != nil {
		in, out := &in.ShareDistribution, &out.ShareDistribution
		*out = make([]ShareDistributionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FairSharePolicyParameters.
func (in *FairSharePolicyParameters) DeepCopy() *FairSharePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(FairSharePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinition) DeepCopyInto(out *JobDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinition.
func (in *JobDefinition) DeepCopy() *JobDefinition {
	if in == nil {
		return nil
	}
	out := new(JobDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinitionInitParameters) DeepCopyInto(out *JobDefinitionInitParameters) {
	*out = *in
	if in.ContainerProperties != nil {
		in, out := &in.ContainerProperties, &out.ContainerProperties
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PlatformCapabilities != nil {
		in, out := &in.PlatformCapabilities, &out.PlatformCapabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PropagateTags != nil {
		in, out := &in.PropagateTags, &out.PropagateTags
		*out = new(bool)
		**out = **in
	}
	if in.RetryStrategy != nil {
		in, out := &in.RetryStrategy, &out.RetryStrategy
		*out = make([]RetryStrategyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = make([]TimeoutInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinitionInitParameters.
func (in *JobDefinitionInitParameters) DeepCopy() *JobDefinitionInitParameters {
	if in == nil {
		return nil
	}
	out := new(JobDefinitionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinitionList) DeepCopyInto(out *JobDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JobDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinitionList.
func (in *JobDefinitionList) DeepCopy() *JobDefinitionList {
	if in == nil {
		return nil
	}
	out := new(JobDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinitionObservation) DeepCopyInto(out *JobDefinitionObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ContainerProperties != nil {
		in, out := &in.ContainerProperties, &out.ContainerProperties
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PlatformCapabilities != nil {
		in, out := &in.PlatformCapabilities, &out.PlatformCapabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PropagateTags != nil {
		in, out := &in.PropagateTags, &out.PropagateTags
		*out = new(bool)
		**out = **in
	}
	if in.RetryStrategy != nil {
		in, out := &in.RetryStrategy, &out.RetryStrategy
		*out = make([]RetryStrategyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = make([]TimeoutObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinitionObservation.
func (in *JobDefinitionObservation) DeepCopy() *JobDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(JobDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinitionParameters) DeepCopyInto(out *JobDefinitionParameters) {
	*out = *in
	if in.ContainerProperties != nil {
		in, out := &in.ContainerProperties, &out.ContainerProperties
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PlatformCapabilities != nil {
		in, out := &in.PlatformCapabilities, &out.PlatformCapabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PropagateTags != nil {
		in, out := &in.PropagateTags, &out.PropagateTags
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RetryStrategy != nil {
		in, out := &in.RetryStrategy, &out.RetryStrategy
		*out = make([]RetryStrategyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = make([]TimeoutParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinitionParameters.
func (in *JobDefinitionParameters) DeepCopy() *JobDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(JobDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinitionSpec) DeepCopyInto(out *JobDefinitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinitionSpec.
func (in *JobDefinitionSpec) DeepCopy() *JobDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(JobDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefinitionStatus) DeepCopyInto(out *JobDefinitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefinitionStatus.
func (in *JobDefinitionStatus) DeepCopy() *JobDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(JobDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryStrategyInitParameters) DeepCopyInto(out *RetryStrategyInitParameters) {
	*out = *in
	if in.Attempts != nil {
		in, out := &in.Attempts, &out.Attempts
		*out = new(float64)
		**out = **in
	}
	if in.EvaluateOnExit != nil {
		in, out := &in.EvaluateOnExit, &out.EvaluateOnExit
		*out = make([]EvaluateOnExitInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryStrategyInitParameters.
func (in *RetryStrategyInitParameters) DeepCopy() *RetryStrategyInitParameters {
	if in == nil {
		return nil
	}
	out := new(RetryStrategyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryStrategyObservation) DeepCopyInto(out *RetryStrategyObservation) {
	*out = *in
	if in.Attempts != nil {
		in, out := &in.Attempts, &out.Attempts
		*out = new(float64)
		**out = **in
	}
	if in.EvaluateOnExit != nil {
		in, out := &in.EvaluateOnExit, &out.EvaluateOnExit
		*out = make([]EvaluateOnExitObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryStrategyObservation.
func (in *RetryStrategyObservation) DeepCopy() *RetryStrategyObservation {
	if in == nil {
		return nil
	}
	out := new(RetryStrategyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryStrategyParameters) DeepCopyInto(out *RetryStrategyParameters) {
	*out = *in
	if in.Attempts != nil {
		in, out := &in.Attempts, &out.Attempts
		*out = new(float64)
		**out = **in
	}
	if in.EvaluateOnExit != nil {
		in, out := &in.EvaluateOnExit, &out.EvaluateOnExit
		*out = make([]EvaluateOnExitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryStrategyParameters.
func (in *RetryStrategyParameters) DeepCopy() *RetryStrategyParameters {
	if in == nil {
		return nil
	}
	out := new(RetryStrategyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicy) DeepCopyInto(out *SchedulingPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicy.
func (in *SchedulingPolicy) DeepCopy() *SchedulingPolicy {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SchedulingPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicyInitParameters) DeepCopyInto(out *SchedulingPolicyInitParameters) {
	*out = *in
	if in.FairSharePolicy != nil {
		in, out := &in.FairSharePolicy, &out.FairSharePolicy
		*out = make([]FairSharePolicyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicyInitParameters.
func (in *SchedulingPolicyInitParameters) DeepCopy() *SchedulingPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicyList) DeepCopyInto(out *SchedulingPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SchedulingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicyList.
func (in *SchedulingPolicyList) DeepCopy() *SchedulingPolicyList {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SchedulingPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicyObservation) DeepCopyInto(out *SchedulingPolicyObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.FairSharePolicy != nil {
		in, out := &in.FairSharePolicy, &out.FairSharePolicy
		*out = make([]FairSharePolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicyObservation.
func (in *SchedulingPolicyObservation) DeepCopy() *SchedulingPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicyParameters) DeepCopyInto(out *SchedulingPolicyParameters) {
	*out = *in
	if in.FairSharePolicy != nil {
		in, out := &in.FairSharePolicy, &out.FairSharePolicy
		*out = make([]FairSharePolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicyParameters.
func (in *SchedulingPolicyParameters) DeepCopy() *SchedulingPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicySpec) DeepCopyInto(out *SchedulingPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicySpec.
func (in *SchedulingPolicySpec) DeepCopy() *SchedulingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicyStatus) DeepCopyInto(out *SchedulingPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicyStatus.
func (in *SchedulingPolicyStatus) DeepCopy() *SchedulingPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareDistributionInitParameters) DeepCopyInto(out *ShareDistributionInitParameters) {
	*out = *in
	if in.ShareIdentifier != nil {
		in, out := &in.ShareIdentifier, &out.ShareIdentifier
		*out = new(string)
		**out = **in
	}
	if in.WeightFactor != nil {
		in, out := &in.WeightFactor, &out.WeightFactor
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareDistributionInitParameters.
func (in *ShareDistributionInitParameters) DeepCopy() *ShareDistributionInitParameters {
	if in == nil {
		return nil
	}
	out := new(ShareDistributionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareDistributionObservation) DeepCopyInto(out *ShareDistributionObservation) {
	*out = *in
	if in.ShareIdentifier != nil {
		in, out := &in.ShareIdentifier, &out.ShareIdentifier
		*out = new(string)
		**out = **in
	}
	if in.WeightFactor != nil {
		in, out := &in.WeightFactor, &out.WeightFactor
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareDistributionObservation.
func (in *ShareDistributionObservation) DeepCopy() *ShareDistributionObservation {
	if in == nil {
		return nil
	}
	out := new(ShareDistributionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareDistributionParameters) DeepCopyInto(out *ShareDistributionParameters) {
	*out = *in
	if in.ShareIdentifier != nil {
		in, out := &in.ShareIdentifier, &out.ShareIdentifier
		*out = new(string)
		**out = **in
	}
	if in.WeightFactor != nil {
		in, out := &in.WeightFactor, &out.WeightFactor
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareDistributionParameters.
func (in *ShareDistributionParameters) DeepCopy() *ShareDistributionParameters {
	if in == nil {
		return nil
	}
	out := new(ShareDistributionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeoutInitParameters) DeepCopyInto(out *TimeoutInitParameters) {
	*out = *in
	if in.AttemptDurationSeconds != nil {
		in, out := &in.AttemptDurationSeconds, &out.AttemptDurationSeconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeoutInitParameters.
func (in *TimeoutInitParameters) DeepCopy() *TimeoutInitParameters {
	if in == nil {
		return nil
	}
	out := new(TimeoutInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeoutObservation) DeepCopyInto(out *TimeoutObservation) {
	*out = *in
	if in.AttemptDurationSeconds != nil {
		in, out := &in.AttemptDurationSeconds, &out.AttemptDurationSeconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeoutObservation.
func (in *TimeoutObservation) DeepCopy() *TimeoutObservation {
	if in == nil {
		return nil
	}
	out := new(TimeoutObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeoutParameters) DeepCopyInto(out *TimeoutParameters) {
	*out = *in
	if in.AttemptDurationSeconds != nil {
		in, out := &in.AttemptDurationSeconds, &out.AttemptDurationSeconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeoutParameters.
func (in *TimeoutParameters) DeepCopy() *TimeoutParameters {
	if in == nil {
		return nil
	}
	out := new(TimeoutParameters)
	in.DeepCopyInto(out)
	return out
}
