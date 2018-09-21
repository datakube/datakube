package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BackupTarget describes a BackupTarget resource
type BackupTarget struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	meta_v1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object, including
	// things like...
	//  - name
	//  - namespace
	//  - self link
	//  - labels
	//  - ... etc ...
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the custom resource spec
	Spec BackupTargetSpec `json:"spec"`
}

// MyResourceSpec is the spec for a MyResource resource
type BackupTargetSpec struct {
	// Message and SomeValue are example custom spec fields
	//
	// this is where you would put your custom resource data
	Host   		string `json:"host"`
	Type   		string `json:"type"`
	Name   		string `json:"name"`
	User   		string `json:"user"`
	Password	string `json:"password"`
	Port   		string `json:"port"`
	Interval   	string `json:"interval"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BackupTargetList is a list of MyResource resources
type BackupTargetList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []BackupTarget `json:"items"`
}
