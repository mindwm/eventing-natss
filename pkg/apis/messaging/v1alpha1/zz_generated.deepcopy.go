//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerConfigTemplate) DeepCopyInto(out *ConsumerConfigTemplate) {
	*out = *in
	if in.OptStartTime != nil {
		in, out := &in.OptStartTime, &out.OptStartTime
		*out = (*in).DeepCopy()
	}
	out.AckWait = in.AckWait
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerConfigTemplate.
func (in *ConsumerConfigTemplate) DeepCopy() *ConsumerConfigTemplate {
	if in == nil {
		return nil
	}
	out := new(ConsumerConfigTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatsJetStreamChannel) DeepCopyInto(out *NatsJetStreamChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatsJetStreamChannel.
func (in *NatsJetStreamChannel) DeepCopy() *NatsJetStreamChannel {
	if in == nil {
		return nil
	}
	out := new(NatsJetStreamChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NatsJetStreamChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatsJetStreamChannelList) DeepCopyInto(out *NatsJetStreamChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NatsJetStreamChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatsJetStreamChannelList.
func (in *NatsJetStreamChannelList) DeepCopy() *NatsJetStreamChannelList {
	if in == nil {
		return nil
	}
	out := new(NatsJetStreamChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NatsJetStreamChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatsJetStreamChannelSpec) DeepCopyInto(out *NatsJetStreamChannelSpec) {
	*out = *in
	in.ChannelableSpec.DeepCopyInto(&out.ChannelableSpec)
	in.Stream.DeepCopyInto(&out.Stream)
	if in.ConsumerConfigTemplate != nil {
		in, out := &in.ConsumerConfigTemplate, &out.ConsumerConfigTemplate
		*out = new(ConsumerConfigTemplate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatsJetStreamChannelSpec.
func (in *NatsJetStreamChannelSpec) DeepCopy() *NatsJetStreamChannelSpec {
	if in == nil {
		return nil
	}
	out := new(NatsJetStreamChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatsJetStreamChannelStatus) DeepCopyInto(out *NatsJetStreamChannelStatus) {
	*out = *in
	in.ChannelableStatus.DeepCopyInto(&out.ChannelableStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatsJetStreamChannelStatus.
func (in *NatsJetStreamChannelStatus) DeepCopy() *NatsJetStreamChannelStatus {
	if in == nil {
		return nil
	}
	out := new(NatsJetStreamChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stream) DeepCopyInto(out *Stream) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(StreamConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stream.
func (in *Stream) DeepCopy() *Stream {
	if in == nil {
		return nil
	}
	out := new(Stream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamConfig) DeepCopyInto(out *StreamConfig) {
	*out = *in
	if in.AdditionalSubjects != nil {
		in, out := &in.AdditionalSubjects, &out.AdditionalSubjects
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.MaxAge = in.MaxAge
	out.DuplicateWindow = in.DuplicateWindow
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = new(StreamPlacement)
		(*in).DeepCopyInto(*out)
	}
	if in.Mirror != nil {
		in, out := &in.Mirror, &out.Mirror
		*out = new(StreamSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]StreamSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamConfig.
func (in *StreamConfig) DeepCopy() *StreamConfig {
	if in == nil {
		return nil
	}
	out := new(StreamConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamPlacement) DeepCopyInto(out *StreamPlacement) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamPlacement.
func (in *StreamPlacement) DeepCopy() *StreamPlacement {
	if in == nil {
		return nil
	}
	out := new(StreamPlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamSource) DeepCopyInto(out *StreamSource) {
	*out = *in
	if in.OptStartTime != nil {
		in, out := &in.OptStartTime, &out.OptStartTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamSource.
func (in *StreamSource) DeepCopy() *StreamSource {
	if in == nil {
		return nil
	}
	out := new(StreamSource)
	in.DeepCopyInto(out)
	return out
}
