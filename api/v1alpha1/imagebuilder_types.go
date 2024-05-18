package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ImageBuilderSpec defines the desired state of ImageBuilder
type ImageBuilderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ImageBuilder. Edit imagebuilder_types.go to remove/update
	Foo string `json:"foo,omitempty"`
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