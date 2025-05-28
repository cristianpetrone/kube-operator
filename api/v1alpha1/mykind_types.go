package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Action defines the possible actions the Cleaner can take on resources.
type Action string

const (
	// Delete is the default action to remove unused resources.
	Delete Action = "Delete"
	// Update is an optional action to patch or label resources.
	Update Action = "Update"
)

// MyKindSpec defines the desired state of MyKind
type MyKindSpec struct {
	// Criteria defines the conditions under which a resource is considered unused or unhealthy.
	Criteria string `json:"criteria"`

	// Schedule defines how often the controller should scan the cluster for matching resources.
	Schedule string `json:"schedule"`

	// +kubebuilder:default:=Delete
	// Action to take on the resources matching the criteria. Defaults to Delete.
	Action Action `json:"action,omitempty"`

	// +optional
	// Transform specifies an optional mutation to apply (e.g., label, annotation).
	Transform string `json:"transform,omitempty"`

	// +optional
	// Namespaces allows filtering the scan to specific namespaces.
	Namespaces []string `json:"namespaces,omitempty"`
}

// MyKindStatus defines the observed state of MyKind
type MyKindStatus struct {
	// +optional
	// LastRunTime is the timestamp of the most recent cleaner run.
	LastRunTime *metav1.Time `json:"lastRunTime,omitempty"`

	// +optional
	// NextScheduleTime is when the next run is scheduled.
	NextScheduleTime *metav1.Time `json:"nextScheduleTime,omitempty"`

	// +optional
	// FailureMessage captures any error from the last run.
	FailureMessage string `json:"failureMessage,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=mykinds,scope=Cluster

// MyKind is the Schema for the mykinds API
type MyKind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyKindSpec   `json:"spec,omitempty"`
	Status MyKindStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MyKindList contains a list of MyKind
type MyKindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyKind `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MyKind{}, &MyKindList{})
}

