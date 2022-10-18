/*
Copyright 2022.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MoonSpec defines the desired state of Moon
type MoonSpec struct {
	Foo string `json:"foo,omitempty"`
	Bar string `json:"bar,omitempty"`
	Qux string `json:"qux,omitempty"`
}

// MoonStatus defines the observed state of Moon
type MoonStatus struct {
	Foo string `json:"foo,omitempty"`
	Bar string `json:"bar,omitempty"`
	Qux string `json:"qux,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:JSONPath=".status.foo",name=Foo,type=string
//+kubebuilder:printcolumn:JSONPath=".status.bar",name=Bar,type=string
//+kubebuilder:printcolumn:JSONPath=".status.qux",name=Qux,type=string
//+kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="age"

// Moon is the Schema for the moons API
type Moon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MoonSpec   `json:"spec,omitempty"`
	Status MoonStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MoonList contains a list of Moon
type MoonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Moon `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Moon{}, &MoonList{})
}
