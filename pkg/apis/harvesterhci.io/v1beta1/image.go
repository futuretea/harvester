package v1beta1

import (
	"github.com/rancher/wrangler/pkg/condition"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ImageInitialized	condition.Cond	= "Initialized"
	ImageImported		condition.Cond	= "Imported"
)

const (
	VirtualMachineImageSourceTypeDownload		= "download"
	VirtualMachineImageSourceTypeUpload		= "upload"
	VirtualMachineImageSourceTypeExportVolume	= "export-from-volume"
)

type VirtualMachineImage struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	VirtualMachineImageSpec		`json:"spec"`
	Status	VirtualMachineImageStatus	`json:"status,omitempty"`
}

type VirtualMachineImageSpec struct {
	Description	string	`json:"description,omitempty"`

	DisplayName	string	`json:"displayName"`

	SourceType	string	`json:"sourceType"`

	PVCName	string	`json:"pvcName"`

	PVCNamespace	string	`json:"pvcNamespace"`

	URL	string	`json:"url"`

	Checksum	string	`json:"checksum"`
}

type VirtualMachineImageStatus struct {
	AppliedURL	string	`json:"appliedUrl,omitempty"`

	Progress	int	`json:"progress,omitempty"`

	Size	int64	`json:"size,omitempty"`

	StorageClassName	string	`json:"storageClassName,omitempty"`

	Conditions	[]Condition	`json:"conditions,omitempty"`
}

type Condition struct {
	Type	condition.Cond	`json:"type"`

	Status	v1.ConditionStatus	`json:"status"`

	LastUpdateTime	string	`json:"lastUpdateTime,omitempty"`

	LastTransitionTime	string	`json:"lastTransitionTime,omitempty"`

	Reason	string	`json:"reason,omitempty"`

	Message	string	`json:"message,omitempty"`
}
