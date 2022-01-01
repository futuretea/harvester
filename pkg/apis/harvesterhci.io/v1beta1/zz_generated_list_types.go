package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KeyPairList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`

	Items	[]KeyPair	`json:"items"`
}

func NewKeyPair(namespace, name string, obj KeyPair) *KeyPair {
	__traceStack()

	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("KeyPair").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type PreferenceList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`

	Items	[]Preference	`json:"items"`
}

func NewPreference(namespace, name string, obj Preference) *Preference {
	__traceStack()

	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Preference").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type SettingList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`

	Items	[]Setting	`json:"items"`
}

func NewSetting(namespace, name string, obj Setting) *Setting {
	__traceStack()

	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Setting").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type UpgradeList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`

	Items	[]Upgrade	`json:"items"`
}

func NewUpgrade(namespace, name string, obj Upgrade) *Upgrade {
	__traceStack()

	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Upgrade").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type VersionList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`

	Items	[]Version	`json:"items"`
}

func NewVersion(namespace, name string, obj Version) *Version {
	__traceStack()

	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Version").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type VirtualMachineBackupList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`

	Items	[]VirtualMachineBackup	`json:"items"`
}

func NewVirtualMachineBackup(namespace, name string, obj VirtualMachineBackup) *VirtualMachineBackup {
	__traceStack()

	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VirtualMachineBackup").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type VirtualMachineRestoreList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`

	Items	[]VirtualMachineRestore	`json:"items"`
}

func NewVirtualMachineRestore(namespace, name string, obj VirtualMachineRestore) *VirtualMachineRestore {
	__traceStack()

	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VirtualMachineRestore").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type VirtualMachineImageList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`

	Items	[]VirtualMachineImage	`json:"items"`
}

func NewVirtualMachineImage(namespace, name string, obj VirtualMachineImage) *VirtualMachineImage {
	__traceStack()

	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VirtualMachineImage").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type VirtualMachineTemplateList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`

	Items	[]VirtualMachineTemplate	`json:"items"`
}

func NewVirtualMachineTemplate(namespace, name string, obj VirtualMachineTemplate) *VirtualMachineTemplate {
	__traceStack()

	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VirtualMachineTemplate").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type VirtualMachineTemplateVersionList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`

	Items	[]VirtualMachineTemplateVersion	`json:"items"`
}

func NewVirtualMachineTemplateVersion(namespace, name string, obj VirtualMachineTemplateVersion) *VirtualMachineTemplateVersion {
	__traceStack()

	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VirtualMachineTemplateVersion").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type SupportBundleList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`

	Items	[]SupportBundle	`json:"items"`
}

func NewSupportBundle(namespace, name string, obj SupportBundle) *SupportBundle {
	__traceStack()

	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("SupportBundle").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}
