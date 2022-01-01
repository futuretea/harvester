package fake

import (
	v1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/harvesterhci.io/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHarvesterhciV1beta1 struct {
	*testing.Fake
}

func (c *FakeHarvesterhciV1beta1) KeyPairs(namespace string) v1beta1.KeyPairInterface {
	__traceStack()

	return &FakeKeyPairs{c, namespace}
}

func (c *FakeHarvesterhciV1beta1) Preferences(namespace string) v1beta1.PreferenceInterface {
	__traceStack()

	return &FakePreferences{c, namespace}
}

func (c *FakeHarvesterhciV1beta1) Settings() v1beta1.SettingInterface {
	__traceStack()

	return &FakeSettings{c}
}

func (c *FakeHarvesterhciV1beta1) SupportBundles(namespace string) v1beta1.SupportBundleInterface {
	__traceStack()

	return &FakeSupportBundles{c, namespace}
}

func (c *FakeHarvesterhciV1beta1) Upgrades(namespace string) v1beta1.UpgradeInterface {
	__traceStack()

	return &FakeUpgrades{c, namespace}
}

func (c *FakeHarvesterhciV1beta1) Versions(namespace string) v1beta1.VersionInterface {
	__traceStack()

	return &FakeVersions{c, namespace}
}

func (c *FakeHarvesterhciV1beta1) VirtualMachineBackups(namespace string) v1beta1.VirtualMachineBackupInterface {
	__traceStack()

	return &FakeVirtualMachineBackups{c, namespace}
}

func (c *FakeHarvesterhciV1beta1) VirtualMachineImages(namespace string) v1beta1.VirtualMachineImageInterface {
	__traceStack()

	return &FakeVirtualMachineImages{c, namespace}
}

func (c *FakeHarvesterhciV1beta1) VirtualMachineRestores(namespace string) v1beta1.VirtualMachineRestoreInterface {
	__traceStack()

	return &FakeVirtualMachineRestores{c, namespace}
}

func (c *FakeHarvesterhciV1beta1) VirtualMachineTemplates(namespace string) v1beta1.VirtualMachineTemplateInterface {
	__traceStack()

	return &FakeVirtualMachineTemplates{c, namespace}
}

func (c *FakeHarvesterhciV1beta1) VirtualMachineTemplateVersions(namespace string) v1beta1.VirtualMachineTemplateVersionInterface {
	__traceStack()

	return &FakeVirtualMachineTemplateVersions{c, namespace}
}

func (c *FakeHarvesterhciV1beta1) RESTClient() rest.Interface {
	__traceStack()

	var ret *rest.RESTClient
	return ret
}
