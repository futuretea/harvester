package v1beta1

import (
	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type HarvesterhciV1beta1Interface interface {
	RESTClient() rest.Interface
	KeyPairsGetter
	PreferencesGetter
	SettingsGetter
	SupportBundlesGetter
	UpgradesGetter
	VersionsGetter
	VirtualMachineBackupsGetter
	VirtualMachineImagesGetter
	VirtualMachineRestoresGetter
	VirtualMachineTemplatesGetter
	VirtualMachineTemplateVersionsGetter
}

type HarvesterhciV1beta1Client struct {
	restClient rest.Interface
}

func (c *HarvesterhciV1beta1Client) KeyPairs(namespace string) KeyPairInterface {
	__traceStack()

	return newKeyPairs(c, namespace)
}

func (c *HarvesterhciV1beta1Client) Preferences(namespace string) PreferenceInterface {
	__traceStack()

	return newPreferences(c, namespace)
}

func (c *HarvesterhciV1beta1Client) Settings() SettingInterface {
	__traceStack()

	return newSettings(c)
}

func (c *HarvesterhciV1beta1Client) SupportBundles(namespace string) SupportBundleInterface {
	__traceStack()

	return newSupportBundles(c, namespace)
}

func (c *HarvesterhciV1beta1Client) Upgrades(namespace string) UpgradeInterface {
	__traceStack()

	return newUpgrades(c, namespace)
}

func (c *HarvesterhciV1beta1Client) Versions(namespace string) VersionInterface {
	__traceStack()

	return newVersions(c, namespace)
}

func (c *HarvesterhciV1beta1Client) VirtualMachineBackups(namespace string) VirtualMachineBackupInterface {
	__traceStack()

	return newVirtualMachineBackups(c, namespace)
}

func (c *HarvesterhciV1beta1Client) VirtualMachineImages(namespace string) VirtualMachineImageInterface {
	__traceStack()

	return newVirtualMachineImages(c, namespace)
}

func (c *HarvesterhciV1beta1Client) VirtualMachineRestores(namespace string) VirtualMachineRestoreInterface {
	__traceStack()

	return newVirtualMachineRestores(c, namespace)
}

func (c *HarvesterhciV1beta1Client) VirtualMachineTemplates(namespace string) VirtualMachineTemplateInterface {
	__traceStack()

	return newVirtualMachineTemplates(c, namespace)
}

func (c *HarvesterhciV1beta1Client) VirtualMachineTemplateVersions(namespace string) VirtualMachineTemplateVersionInterface {
	__traceStack()

	return newVirtualMachineTemplateVersions(c, namespace)
}

func NewForConfig(c *rest.Config) (*HarvesterhciV1beta1Client, error) {
	__traceStack()

	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &HarvesterhciV1beta1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *HarvesterhciV1beta1Client {
	__traceStack()

	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *HarvesterhciV1beta1Client {
	__traceStack()

	return &HarvesterhciV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	__traceStack()

	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

func (c *HarvesterhciV1beta1Client) RESTClient() rest.Interface {
	__traceStack()

	if c == nil {
		return nil
	}
	return c.restClient
}
