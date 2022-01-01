package v1beta1

import (
	"github.com/rancher/wrangler/pkg/condition"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	BackupConditionReady	condition.Cond	= "Ready"

	BackupConditionProgressing	condition.Cond	= "InProgress"
)

type DeletionPolicy string

const (
	VirtualMachineRestoreDelete	DeletionPolicy	= "delete"

	VirtualMachineRestoreRetain	DeletionPolicy	= "retain"
)

type VirtualMachineBackup struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	VirtualMachineBackupSpec	`json:"spec"`

	Status	*VirtualMachineBackupStatus	`json:"status,omitempty"`
}

type VirtualMachineBackupSpec struct {
	Source corev1.TypedLocalObjectReference `json:"source"`
}

type VirtualMachineBackupStatus struct {
	SourceUID	*types.UID	`json:"sourceUID,omitempty"`

	CreationTime	*metav1.Time	`json:"creationTime,omitempty"`

	BackupTarget	*BackupTarget	`json:"backupTarget,omitempty"`

	SourceSpec	*VirtualMachineSourceSpec	`json:"source,omitempty"`

	VolumeBackups	[]VolumeBackup	`json:"volumeBackups,omitempty"`

	SecretBackups	[]SecretBackup	`json:"secretBackups,omitempty"`

	ReadyToUse	*bool	`json:"readyToUse,omitempty"`

	Error	*Error	`json:"error,omitempty"`

	Conditions	[]Condition	`json:"conditions,omitempty"`
}

type BackupTarget struct {
	Endpoint	string	`json:"endpoint,omitempty"`
	BucketName	string	`json:"bucketName,omitempty"`
	BucketRegion	string	`json:"bucketRegion,omitempty"`
}

type Error struct {
	Time	*metav1.Time	`json:"time,omitempty"`

	Message	*string	`json:"message,omitempty"`
}

type VolumeBackup struct {
	Name	*string	`json:"name,omitempty"`

	VolumeName	string	`json:"volumeName"`

	CreationTime	*metav1.Time	`json:"creationTime,omitempty"`

	PersistentVolumeClaim	PersistentVolumeClaimSourceSpec	`json:"persistentVolumeClaim"`

	LonghornBackupName	*string	`json:"longhornBackupName,omitempty"`

	ReadyToUse	*bool	`json:"readyToUse,omitempty"`

	Error	*Error	`json:"error,omitempty"`
}

type SecretBackup struct {
	Name	string	`json:"name,omitempty"`

	Data	map[string][]byte	`json:"data,omitempty"`
}

type PersistentVolumeClaimSourceSpec struct {
	ObjectMeta	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	corev1.PersistentVolumeClaimSpec	`json:"spec,omitempty"`
}

type VirtualMachineRestore struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	VirtualMachineRestoreSpec	`json:"spec"`

	Status	*VirtualMachineRestoreStatus	`json:"status,omitempty"`
}

type VirtualMachineRestoreSpec struct {
	Target	corev1.TypedLocalObjectReference	`json:"target"`

	VirtualMachineBackupName	string	`json:"virtualMachineBackupName"`

	VirtualMachineBackupNamespace	string	`json:"virtualMachineBackupNamespace"`

	NewVM	bool	`json:"newVM,omitempty"`

	DeletionPolicy	DeletionPolicy	`json:"deletionPolicy,omitempty"`
}

type VirtualMachineRestoreStatus struct {
	VolumeRestores	[]VolumeRestore	`json:"restores,omitempty"`

	RestoreTime	*metav1.Time	`json:"restoreTime,omitempty"`

	DeletedVolumes	[]string	`json:"deletedVolumes,omitempty"`

	Complete	*bool	`json:"complete,omitempty"`

	Conditions	[]Condition	`json:"conditions,omitempty"`

	TargetUID	*types.UID	`json:"targetUID,omitempty"`
}

type VolumeRestore struct {
	VolumeName	string	`json:"volumeName,omitempty"`

	PersistentVolumeClaim	PersistentVolumeClaimSourceSpec	`json:"persistentVolumeClaimSpec,omitempty"`

	VolumeBackupName	string	`json:"volumeBackupName,omitempty"`
}
