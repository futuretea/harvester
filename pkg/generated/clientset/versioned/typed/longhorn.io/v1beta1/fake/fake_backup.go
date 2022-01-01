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

type FakeBackups struct {
	Fake	*FakeLonghornV1beta1
	ns	string
}

var backupsResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "backups"}

var backupsKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Backup"}

func (c *FakeBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Backup, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupsResource, c.ns, name), &v1beta1.Backup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Backup), err
}

func (c *FakeBackups) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackupList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupsResource, backupsKind, c.ns, opts), &v1beta1.BackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BackupList{ListMeta: obj.(*v1beta1.BackupList).ListMeta}
	for _, item := range obj.(*v1beta1.BackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupsResource, c.ns, opts))

}

func (c *FakeBackups) Create(ctx context.Context, backup *v1beta1.Backup, opts v1.CreateOptions) (result *v1beta1.Backup, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupsResource, c.ns, backup), &v1beta1.Backup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Backup), err
}

func (c *FakeBackups) Update(ctx context.Context, backup *v1beta1.Backup, opts v1.UpdateOptions) (result *v1beta1.Backup, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupsResource, c.ns, backup), &v1beta1.Backup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Backup), err
}

func (c *FakeBackups) UpdateStatus(ctx context.Context, backup *v1beta1.Backup, opts v1.UpdateOptions) (*v1beta1.Backup, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupsResource, "status", c.ns, backup), &v1beta1.Backup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Backup), err
}

func (c *FakeBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupsResource, c.ns, name), &v1beta1.Backup{})

	return err
}

func (c *FakeBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(backupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.BackupList{})
	return err
}

func (c *FakeBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Backup, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupsResource, c.ns, name, pt, data, subresources...), &v1beta1.Backup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Backup), err
}
