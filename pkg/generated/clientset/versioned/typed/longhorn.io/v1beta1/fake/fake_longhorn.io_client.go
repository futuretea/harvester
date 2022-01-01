package fake

import (
	v1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/longhorn.io/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLonghornV1beta1 struct {
	*testing.Fake
}

func (c *FakeLonghornV1beta1) BackingImages(namespace string) v1beta1.BackingImageInterface {
	__traceStack()

	return &FakeBackingImages{c, namespace}
}

func (c *FakeLonghornV1beta1) BackingImageDataSources(namespace string) v1beta1.BackingImageDataSourceInterface {
	__traceStack()

	return &FakeBackingImageDataSources{c, namespace}
}

func (c *FakeLonghornV1beta1) BackingImageManagers(namespace string) v1beta1.BackingImageManagerInterface {
	__traceStack()

	return &FakeBackingImageManagers{c, namespace}
}

func (c *FakeLonghornV1beta1) Backups(namespace string) v1beta1.BackupInterface {
	__traceStack()

	return &FakeBackups{c, namespace}
}

func (c *FakeLonghornV1beta1) BackupTargets(namespace string) v1beta1.BackupTargetInterface {
	__traceStack()

	return &FakeBackupTargets{c, namespace}
}

func (c *FakeLonghornV1beta1) BackupVolumes(namespace string) v1beta1.BackupVolumeInterface {
	__traceStack()

	return &FakeBackupVolumes{c, namespace}
}

func (c *FakeLonghornV1beta1) Engines(namespace string) v1beta1.EngineInterface {
	__traceStack()

	return &FakeEngines{c, namespace}
}

func (c *FakeLonghornV1beta1) EngineImages(namespace string) v1beta1.EngineImageInterface {
	__traceStack()

	return &FakeEngineImages{c, namespace}
}

func (c *FakeLonghornV1beta1) InstanceManagers(namespace string) v1beta1.InstanceManagerInterface {
	__traceStack()

	return &FakeInstanceManagers{c, namespace}
}

func (c *FakeLonghornV1beta1) Nodes(namespace string) v1beta1.NodeInterface {
	__traceStack()

	return &FakeNodes{c, namespace}
}

func (c *FakeLonghornV1beta1) RecurringJobs(namespace string) v1beta1.RecurringJobInterface {
	__traceStack()

	return &FakeRecurringJobs{c, namespace}
}

func (c *FakeLonghornV1beta1) Replicas(namespace string) v1beta1.ReplicaInterface {
	__traceStack()

	return &FakeReplicas{c, namespace}
}

func (c *FakeLonghornV1beta1) Settings(namespace string) v1beta1.SettingInterface {
	__traceStack()

	return &FakeSettings{c, namespace}
}

func (c *FakeLonghornV1beta1) ShareManagers(namespace string) v1beta1.ShareManagerInterface {
	__traceStack()

	return &FakeShareManagers{c, namespace}
}

func (c *FakeLonghornV1beta1) Volumes(namespace string) v1beta1.VolumeInterface {
	__traceStack()

	return &FakeVolumes{c, namespace}
}

func (c *FakeLonghornV1beta1) RESTClient() rest.Interface {
	__traceStack()

	var ret *rest.RESTClient
	return ret
}
