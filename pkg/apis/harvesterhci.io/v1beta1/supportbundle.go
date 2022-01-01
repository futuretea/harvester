package v1beta1

import (
	"github.com/rancher/wrangler/pkg/condition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	SupportBundleInitialized condition.Cond = "Initialized"
)

type SupportBundle struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	SupportBundleSpec	`json:"spec"`
	Status	SupportBundleStatus	`json:"status,omitempty"`
}

type SupportBundleSpec struct {
	IssueURL	string	`json:"issueURL"`

	Description	string	`json:"description"`
}

type SupportBundleStatus struct {
	State	string	`json:"state,omitempty"`

	Progress	int	`json:"progress,omitempty"`

	Filename	string	`json:"filename,omitempty"`

	Filesize	int64	`json:"filesize,omitempty"`

	Conditions	[]Condition	`json:"conditions,omitempty"`
}
