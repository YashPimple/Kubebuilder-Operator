/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DemoVolumesSpec defines the desired state of DemoVolumes
type DemoVolumesSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of DemoVolumes. Edit demovolumes_types.go to remove/update
	Name string `json:"name,omitempty"`
	Size int    `json:"size,omitempty"`
}

// DemoVolumesStatus defines the observed state of DemoVolumes
type DemoVolumesStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Name string `json:"name,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DemoVolumes is the Schema for the demovolumes API
type DemoVolumes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoVolumesSpec   `json:"spec,omitempty"`
	Status DemoVolumesStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DemoVolumesList contains a list of DemoVolumes
type DemoVolumesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DemoVolumes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DemoVolumes{}, &DemoVolumesList{})
}
