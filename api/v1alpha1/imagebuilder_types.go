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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LocalObjectReference struct {
	// Name of the referent.
	// +required
	Name string `json:"name"`
}

// ImageBuilderSpecSource defines the repository source
type ImageBuilderSpecSource struct {
	// URL of the Git repository in http, https.
	// +kubebuilder:validation:Pattern="^(http|https)://.*$"
	// +required
	URL string `json:"url"`
	// SecretRef is a reference to a secret containing the credentials needed to access the repository
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty"`
}

// ImageBuilderSpecSource defines the desired state of ImageBuilder
type ImageBuilderSpecDestination struct {
	// Image name must be in the form [hostname[:port]/][namespace/]repository
	//   - hostname is registry hostname (default: index.docker.io)
	//   - port is the port of the registry (default: 443)
	//   - namespace is the path within the registry (default: library)
	//   - repository is the name of the image
	// +required
	Image string `json:"image"`
	// SecretRef is a reference to a secret containing the credentials needed to scan and push the images
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty"`
}

// ImageBuilderSpecSource defines the desired state of ImageBuilder
type ImageBuilderSpecRuleSource struct {
	// Type of the source (e.g. branch, tag)
	// +kubebuilder:validation:Enum=branch;tag
	// +kubebuilder:default:="branch"
	// +required
	Type string `json:"type,omitempty"`
	// Pattern to match the source (e.g. main, v1.0)
	// +kubebuilder:default:="main"
	// +optional
	Pattern string `json:"pattern,omitempty"`
}

// ImageBuilderSpecRuleBuild defines the desired state of ImageBuilder
type ImageBuilderSpecRuleBuild struct {
	// Dockerfile to use for building the image default is Dockerfile
	// +kubebuilder:default:="Dockerfile"
	// +required
	File string `json:"file"`
	// Context to use for building the image default is .
	// +kubebuilder:default:="."
	// +required
	Context string `json:"context"`
	// Platforms to use for building the image default is linux/amd64
	// +kubebuilder:validation:MinItems:=1
	// +kubebuilder:default:={"linux/amd64"}
	// +required
	Platforms []string `json:"platforms"`
	// Target to use for building the image
	// +optional
	Target string `json:"target,omitempty"`
}

// ImageBuilderSpecRule defines the desired state of ImageBuilder
type ImageBuilderSpecRule struct {
	// Source of the rule
	// +required
	Source ImageBuilderSpecRuleSource `json:"source"`
	// Build configuration of the rule
	// TODO: kubebuilder:default:={} does not work
	// +kubebuilder:default:={file: "Dockerfile", context: ".", platforms: {"linux/amd64"}}
	// +required
	Build ImageBuilderSpecRuleBuild `json:"build"`
	// Tags to apply to the image
	// +kubebuilder:validation:MinItems:=1
	// +kubebuilder:default:={latest}
	// +required
	Tags []string `json:"tags"`
}

// ImageBuilderSpec defines the desired state of ImageBuilder
type ImageBuilderSpec struct {
	// Source of the repository
	// +required
	Source ImageBuilderSpecSource `json:"source"`
	// Destination of the built image
	// +required
	Destination ImageBuilderSpecDestination `json:"destination"`
	// Build rules
	// TODO: kubebuilder:default:={} does not work
	// +kubebuilder:default:={{source:{type: "branch", pattern: "main"}, build: {file: "Dockerfile", context: ".", platforms: {"linux/amd64"} }, tags: {"latest"}}}
	// +required
	Rules []ImageBuilderSpecRule `json:"rules"`
}

// ImageBuilderStatus defines the observed state of ImageBuilder
type ImageBuilderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ImageBuilder is the Schema for the imagebuilders API
type ImageBuilder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImageBuilderSpec   `json:"spec,omitempty"`
	Status ImageBuilderStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ImageBuilderList contains a list of ImageBuilder
type ImageBuilderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageBuilder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImageBuilder{}, &ImageBuilderList{})
}
