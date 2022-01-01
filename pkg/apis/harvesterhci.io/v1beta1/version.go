package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Version struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	VersionSpec	`json:"spec"`
}

type VersionSpec struct {
	ISOURL	string	`json:"isoURL"`

	ISOChecksum	string	`json:"isoChecksum"`

	ReleaseDate	string	`json:"releaseDate"`

	MinUpgradableVersion	string	`json:"minUpgradableVersion,omitempty"`

	Tags	[]string	`json:"tags"`
}
