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

type FakeBackupTargets struct {
	Fake	*FakeLonghornV1beta1
	ns	string
}

var backuptargetsResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "backuptargets"}

var backuptargetsKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "BackupTarget"}

func (c *FakeBackupTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackupTarget, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backuptargetsResource, c.ns, name), &v1beta1.BackupTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupTarget), err
}

func (c *FakeBackupTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackupTargetList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(backuptargetsResource, backuptargetsKind, c.ns, opts), &v1beta1.BackupTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BackupTargetList{ListMeta: obj.(*v1beta1.BackupTargetList).ListMeta}
	for _, item := range obj.(*v1beta1.BackupTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeBackupTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backuptargetsResource, c.ns, opts))

}

func (c *FakeBackupTargets) Create(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.CreateOptions) (result *v1beta1.BackupTarget, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backuptargetsResource, c.ns, backupTarget), &v1beta1.BackupTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupTarget), err
}

func (c *FakeBackupTargets) Update(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.UpdateOptions) (result *v1beta1.BackupTarget, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backuptargetsResource, c.ns, backupTarget), &v1beta1.BackupTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupTarget), err
}

func (c *FakeBackupTargets) UpdateStatus(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.UpdateOptions) (*v1beta1.BackupTarget, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backuptargetsResource, "status", c.ns, backupTarget), &v1beta1.BackupTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupTarget), err
}

func (c *FakeBackupTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backuptargetsResource, c.ns, name), &v1beta1.BackupTarget{})

	return err
}

func (c *FakeBackupTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(backuptargetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.BackupTargetList{})
	return err
}

func (c *FakeBackupTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackupTarget, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backuptargetsResource, c.ns, name, pt, data, subresources...), &v1beta1.BackupTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupTarget), err
}
