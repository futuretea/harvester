package v1beta1

import (
	"github.com/rancher/wrangler/pkg/condition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kv1 "kubevirt.io/client-go/api/v1"
)

var (
	VersionAssigned condition.Cond = "assigned"
)

type VirtualMachineTemplate struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	VirtualMachineTemplateSpec	`json:"spec,omitempty"`
	Status	VirtualMachineTemplateStatus	`json:"status,omitempty"`
}

type VirtualMachineTemplateSpec struct {
	DefaultVersionID	string	`json:"defaultVersionId"`

	Description	string	`json:"description,omitempty"`
}

type VirtualMachineTemplateStatus struct {
	DefaultVersion	int	`json:"defaultVersion,omitempty"`

	LatestVersion	int	`json:"latestVersion,omitempty"`
}

type VirtualMachineTemplateVersion struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	VirtualMachineTemplateVersionSpec	`json:"spec"`
	Status	VirtualMachineTemplateVersionStatus	`json:"status,omitempty"`
}

type VirtualMachineTemplateVersionSpec struct {
	TemplateID	string	`json:"templateId"`

	Description	string	`json:"description,omitempty"`

	ImageID	string	`json:"imageId,omitempty"`

	KeyPairIDs	[]string	`json:"keyPairIds,omitempty"`

	VM	VirtualMachineSourceSpec	`json:"vm,omitempty"`
}

type VirtualMachineSourceSpec struct {
	ObjectMeta	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	kv1.VirtualMachineSpec	`json:"spec,omitempty"`
}

type VirtualMachineTemplateVersionStatus struct {
	Version	int	`json:"version,omitempty"`

	Conditions	[]Condition	`json:"conditions,omitempty"`
}
