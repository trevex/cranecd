/*
Copyright 2020 CraneCD Authors

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PipelineSpec defines the desired state of Pipeline
type PipelineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ServiceAccount which will be used for spawned deployment jobs
	ServiceAccount string `json:"serviceAccount"`

	// Git repository to deploy
	Git Git `json:"git"`

	// +optional
	// Helm configuration
	Helm *Helm `json:"helm,omitempty"`
}

type Git struct {
	//
	Repository string `json:"repository"`

	// +optional
	Branch string `json:"branch,omitempty"`

	// +optional
	SecretName string `json:"secretName,omitempty"`
}

type Helm struct {
	// +optional
	Repositories []HelmRepository `json:"repositories,omitempty"`

	//
	Chart string `json:"chart"`

	//
	Release string `json:"release"`

	//
	Values []string `json:"values"`

	// TODO: allow overrides?
}

type HelmRepository struct {
	//
	URL string `json:"url"`

	//
	Alias string `json:"alias"`

	// +optional
	SecretName string `json:"secretName,omitempty"`
}

// PipelineStatus defines the observed state of Pipeline
type PipelineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	SharedSecret *corev1.ObjectReference `json:"sharedSecret,omitempty"`
	ActiveJob    *corev1.ObjectReference `json:"activeJob,omitempty"`
}

// +kubebuilder:object:root=true

// Pipeline is the Schema for the pipelines API
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineSpec   `json:"spec,omitempty"`
	Status PipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PipelineList contains a list of Pipeline
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pipeline{}, &PipelineList{})
}
