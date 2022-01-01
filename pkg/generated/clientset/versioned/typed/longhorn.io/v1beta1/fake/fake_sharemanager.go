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

type FakeShareManagers struct {
	Fake	*FakeLonghornV1beta1
	ns	string
}

var sharemanagersResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "sharemanagers"}

var sharemanagersKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "ShareManager"}

func (c *FakeShareManagers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ShareManager, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sharemanagersResource, c.ns, name), &v1beta1.ShareManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ShareManager), err
}

func (c *FakeShareManagers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ShareManagerList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(sharemanagersResource, sharemanagersKind, c.ns, opts), &v1beta1.ShareManagerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ShareManagerList{ListMeta: obj.(*v1beta1.ShareManagerList).ListMeta}
	for _, item := range obj.(*v1beta1.ShareManagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeShareManagers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sharemanagersResource, c.ns, opts))

}

func (c *FakeShareManagers) Create(ctx context.Context, shareManager *v1beta1.ShareManager, opts v1.CreateOptions) (result *v1beta1.ShareManager, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sharemanagersResource, c.ns, shareManager), &v1beta1.ShareManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ShareManager), err
}

func (c *FakeShareManagers) Update(ctx context.Context, shareManager *v1beta1.ShareManager, opts v1.UpdateOptions) (result *v1beta1.ShareManager, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sharemanagersResource, c.ns, shareManager), &v1beta1.ShareManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ShareManager), err
}

func (c *FakeShareManagers) UpdateStatus(ctx context.Context, shareManager *v1beta1.ShareManager, opts v1.UpdateOptions) (*v1beta1.ShareManager, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sharemanagersResource, "status", c.ns, shareManager), &v1beta1.ShareManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ShareManager), err
}

func (c *FakeShareManagers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sharemanagersResource, c.ns, name), &v1beta1.ShareManager{})

	return err
}

func (c *FakeShareManagers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(sharemanagersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ShareManagerList{})
	return err
}

func (c *FakeShareManagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ShareManager, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sharemanagersResource, c.ns, name, pt, data, subresources...), &v1beta1.ShareManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ShareManager), err
}
