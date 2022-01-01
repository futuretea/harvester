package fake

import (
	"context"

	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeBackupVolumes struct {
	Fake	*FakeLonghornV1beta1
	ns	string
}

var backupvolumesResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "backupvolumes"}

var backupvolumesKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "BackupVolume"}

func (c *FakeBackupVolumes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackupVolume, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupvolumesResource, c.ns, name), &v1beta1.BackupVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupVolume), err
}

func (c *FakeBackupVolumes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackupVolumeList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupvolumesResource, backupvolumesKind, c.ns, opts), &v1beta1.BackupVolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BackupVolumeList{ListMeta: obj.(*v1beta1.BackupVolumeList).ListMeta}
	for _, item := range obj.(*v1beta1.BackupVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeBackupVolumes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupvolumesResource, c.ns, opts))

}

func (c *FakeBackupVolumes) Create(ctx context.Context, backupVolume *v1beta1.BackupVolume, opts v1.CreateOptions) (result *v1beta1.BackupVolume, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupvolumesResource, c.ns, backupVolume), &v1beta1.BackupVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupVolume), err
}

func (c *FakeBackupVolumes) Update(ctx context.Context, backupVolume *v1beta1.BackupVolume, opts v1.UpdateOptions) (result *v1beta1.BackupVolume, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupvolumesResource, c.ns, backupVolume), &v1beta1.BackupVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupVolume), err
}

func (c *FakeBackupVolumes) UpdateStatus(ctx context.Context, backupVolume *v1beta1.BackupVolume, opts v1.UpdateOptions) (*v1beta1.BackupVolume, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupvolumesResource, "status", c.ns, backupVolume), &v1beta1.BackupVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupVolume), err
}

func (c *FakeBackupVolumes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupvolumesResource, c.ns, name), &v1beta1.BackupVolume{})

	return err
}

func (c *FakeBackupVolumes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(backupvolumesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.BackupVolumeList{})
	return err
}

func (c *FakeBackupVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackupVolume, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupvolumesResource, c.ns, name, pt, data, subresources...), &v1beta1.BackupVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupVolume), err
}
