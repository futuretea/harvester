package v1beta1

import (
	harvesterhci "github.com/harvester/harvester/pkg/apis/harvesterhci.io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	KeyPairResourceName				= "keypairs"
	PreferenceResourceName				= "preferences"
	SettingResourceName				= "settings"
	SupportBundleResourceName			= "supportbundles"
	UpgradeResourceName				= "upgrades"
	VersionResourceName				= "versions"
	VirtualMachineBackupResourceName		= "virtualmachinebackups"
	VirtualMachineImageResourceName			= "virtualmachineimages"
	VirtualMachineRestoreResourceName		= "virtualmachinerestores"
	VirtualMachineTemplateResourceName		= "virtualmachinetemplates"
	VirtualMachineTemplateVersionResourceName	= "virtualmachinetemplateversions"
)

var SchemeGroupVersion = schema.GroupVersion{Group: harvesterhci.GroupName, Version: "v1beta1"}

func Kind(kind string) schema.GroupKind {
	__traceStack()

	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func Resource(resource string) schema.GroupResource {
	__traceStack()

	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder	= runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme	= SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	__traceStack()

	scheme.AddKnownTypes(SchemeGroupVersion,
		&KeyPair{},
		&KeyPairList{},
		&Preference{},
		&PreferenceList{},
		&Setting{},
		&SettingList{},
		&SupportBundle{},
		&SupportBundleList{},
		&Upgrade{},
		&UpgradeList{},
		&Version{},
		&VersionList{},
		&VirtualMachineBackup{},
		&VirtualMachineBackupList{},
		&VirtualMachineImage{},
		&VirtualMachineImageList{},
		&VirtualMachineRestore{},
		&VirtualMachineRestoreList{},
		&VirtualMachineTemplate{},
		&VirtualMachineTemplateList{},
		&VirtualMachineTemplateVersion{},
		&VirtualMachineTemplateVersionList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
