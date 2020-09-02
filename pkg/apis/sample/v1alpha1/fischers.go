/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	"github.com/pwittrock/apiserver-runtime/pkg/builder/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var _ resource.Object = &Fischer{}
var _ resource.ListObject = &FischerList{}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Fischer defines the schema for the "fischers" resource.
type Fischer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// DisallowedFlunders holds a list of Flunder.Names that are disallowed.
	DisallowedFlunders []string `json:"disallowedFlunders,omitempty" protobuf:"bytes,2,rep,name=disallowedFlunders"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FischerList is a list of Fischer objects.
type FischerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Fischer `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// GetGroupVersionResource returns a GroupVersionResource with "fischers" as the resource.
// GetGroupVersionResource implements resource.Object
func (Fischer) GetGroupVersionResource() schema.GroupVersionResource {
	return SchemeGroupVersion.WithResource("fischers")
}

// GetObjectMeta implements resource.Object
func (f *Fischer) GetObjectMeta() *metav1.ObjectMeta {
	return &f.ObjectMeta
}

// New implements resource.Object
func (Fischer) New() runtime.Object {
	return &Fischer{}
}

// NamespaceScoped returns false to indicate Fischer is NOT a namespaced resource.
// NamespaceScoped implements resource.Object.
func (Fischer) NamespaceScoped() bool {
	return false
}

// NewList implements resource.Object
func (Fischer) NewList() runtime.Object {
	return &FischerList{}
}

// GetListMeta implements resource.Object
func (c *FischerList) GetListMeta() *metav1.ListMeta {
	return &c.ListMeta
}
