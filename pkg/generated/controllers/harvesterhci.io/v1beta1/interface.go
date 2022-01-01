package v1beta1

import (
	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	__traceStack()

	schemes.Register(v1beta1.AddToScheme)
}

type Interface interface {
	KeyPair() KeyPairController
	Preference() PreferenceController
	Setting() SettingController
	SupportBundle() SupportBundleController
	Upgrade() UpgradeController
	Version() VersionController
	VirtualMachineBackup() VirtualMachineBackupController
	VirtualMachineImage() VirtualMachineImageController
	VirtualMachineRestore() VirtualMachineRestoreController
	VirtualMachineTemplate() VirtualMachineTemplateController
	VirtualMachineTemplateVersion() VirtualMachineTemplateVersionController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	__traceStack()

	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) KeyPair() KeyPairController {
	__traceStack()

	return NewKeyPairController(schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "KeyPair"}, "keypairs", true, c.controllerFactory)
}
func (c *version) Preference() PreferenceController {
	__traceStack()

	return NewPreferenceController(schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "Preference"}, "preferences", true, c.controllerFactory)
}
func (c *version) Setting() SettingController {
	__traceStack()

	return NewSettingController(schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "Setting"}, "settings", false, c.controllerFactory)
}
func (c *version) SupportBundle() SupportBundleController {
	__traceStack()

	return NewSupportBundleController(schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "SupportBundle"}, "supportbundles", true, c.controllerFactory)
}
func (c *version) Upgrade() UpgradeController {
	__traceStack()

	return NewUpgradeController(schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "Upgrade"}, "upgrades", true, c.controllerFactory)
}
func (c *version) Version() VersionController {
	__traceStack()

	return NewVersionController(schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "Version"}, "versions", true, c.controllerFactory)
}
func (c *version) VirtualMachineBackup() VirtualMachineBackupController {
	__traceStack()

	return NewVirtualMachineBackupController(schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "VirtualMachineBackup"}, "virtualmachinebackups", true, c.controllerFactory)
}
func (c *version) VirtualMachineImage() VirtualMachineImageController {
	__traceStack()

	return NewVirtualMachineImageController(schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "VirtualMachineImage"}, "virtualmachineimages", true, c.controllerFactory)
}
func (c *version) VirtualMachineRestore() VirtualMachineRestoreController {
	__traceStack()

	return NewVirtualMachineRestoreController(schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "VirtualMachineRestore"}, "virtualmachinerestores", true, c.controllerFactory)
}
func (c *version) VirtualMachineTemplate() VirtualMachineTemplateController {
	__traceStack()

	return NewVirtualMachineTemplateController(schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "VirtualMachineTemplate"}, "virtualmachinetemplates", true, c.controllerFactory)
}
func (c *version) VirtualMachineTemplateVersion() VirtualMachineTemplateVersionController {
	__traceStack()

	return NewVirtualMachineTemplateVersionController(schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "VirtualMachineTemplateVersion"}, "virtualmachinetemplateversions", true, c.controllerFactory)
}
