// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package runtime

import (
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func DeepCopy_runtime_RawExtension(in RawExtension, out *RawExtension, c *conversion.Cloner) error {
	if in.Raw != nil {
		in, out := in.Raw, &out.Raw
		*out = make([]byte, len(in))
		copy(*out, in)
	} else {
		out.Raw = nil
	}
	if in.Object == nil {
		out.Object = nil
	} else if newVal, err := c.DeepCopy(in.Object); err != nil {
		return err
	} else {
		out.Object = newVal.(Object)
	}
	return nil
}

func DeepCopy_runtime_TypeMeta(in TypeMeta, out *TypeMeta, c *conversion.Cloner) error {
	out.APIVersion = in.APIVersion
	out.Kind = in.Kind
	return nil
}

func DeepCopy_runtime_Unknown(in Unknown, out *Unknown, c *conversion.Cloner) error {
	if err := DeepCopy_runtime_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if in.Raw != nil {
		in, out := in.Raw, &out.Raw
		*out = make([]byte, len(in))
		copy(*out, in)
	} else {
		out.Raw = nil
	}
	out.ContentEncoding = in.ContentEncoding
	out.ContentType = in.ContentType
	return nil
}
