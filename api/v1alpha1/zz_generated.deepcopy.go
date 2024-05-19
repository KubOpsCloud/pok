//go:build !ignore_autogenerated

/*
Copyright (C) 2024  KubOps Technology

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuilder) DeepCopyInto(out *ImageBuilder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuilder.
func (in *ImageBuilder) DeepCopy() *ImageBuilder {
	if in == nil {
		return nil
	}
	out := new(ImageBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageBuilder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuilderList) DeepCopyInto(out *ImageBuilderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageBuilder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuilderList.
func (in *ImageBuilderList) DeepCopy() *ImageBuilderList {
	if in == nil {
		return nil
	}
	out := new(ImageBuilderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageBuilderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuilderSpec) DeepCopyInto(out *ImageBuilderSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Destination.DeepCopyInto(&out.Destination)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]ImageBuilderSpecRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuilderSpec.
func (in *ImageBuilderSpec) DeepCopy() *ImageBuilderSpec {
	if in == nil {
		return nil
	}
	out := new(ImageBuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuilderSpecDestination) DeepCopyInto(out *ImageBuilderSpecDestination) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuilderSpecDestination.
func (in *ImageBuilderSpecDestination) DeepCopy() *ImageBuilderSpecDestination {
	if in == nil {
		return nil
	}
	out := new(ImageBuilderSpecDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuilderSpecRule) DeepCopyInto(out *ImageBuilderSpecRule) {
	*out = *in
	out.Source = in.Source
	in.Build.DeepCopyInto(&out.Build)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuilderSpecRule.
func (in *ImageBuilderSpecRule) DeepCopy() *ImageBuilderSpecRule {
	if in == nil {
		return nil
	}
	out := new(ImageBuilderSpecRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuilderSpecRuleBuild) DeepCopyInto(out *ImageBuilderSpecRuleBuild) {
	*out = *in
	if in.Platforms != nil {
		in, out := &in.Platforms, &out.Platforms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuilderSpecRuleBuild.
func (in *ImageBuilderSpecRuleBuild) DeepCopy() *ImageBuilderSpecRuleBuild {
	if in == nil {
		return nil
	}
	out := new(ImageBuilderSpecRuleBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuilderSpecRuleSource) DeepCopyInto(out *ImageBuilderSpecRuleSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuilderSpecRuleSource.
func (in *ImageBuilderSpecRuleSource) DeepCopy() *ImageBuilderSpecRuleSource {
	if in == nil {
		return nil
	}
	out := new(ImageBuilderSpecRuleSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuilderSpecSource) DeepCopyInto(out *ImageBuilderSpecSource) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuilderSpecSource.
func (in *ImageBuilderSpecSource) DeepCopy() *ImageBuilderSpecSource {
	if in == nil {
		return nil
	}
	out := new(ImageBuilderSpecSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuilderStatus) DeepCopyInto(out *ImageBuilderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuilderStatus.
func (in *ImageBuilderStatus) DeepCopy() *ImageBuilderStatus {
	if in == nil {
		return nil
	}
	out := new(ImageBuilderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}
