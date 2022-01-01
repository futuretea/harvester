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

type FakeReplicas struct {
	Fake	*FakeLonghornV1beta1
	ns	string
}

var replicasResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "replicas"}

var replicasKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Replica"}

func (c *FakeReplicas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Replica, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(replicasResource, c.ns, name), &v1beta1.Replica{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Replica), err
}

func (c *FakeReplicas) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ReplicaList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(replicasResource, replicasKind, c.ns, opts), &v1beta1.ReplicaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ReplicaList{ListMeta: obj.(*v1beta1.ReplicaList).ListMeta}
	for _, item := range obj.(*v1beta1.ReplicaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeReplicas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(replicasResource, c.ns, opts))

}

func (c *FakeReplicas) Create(ctx context.Context, replica *v1beta1.Replica, opts v1.CreateOptions) (result *v1beta1.Replica, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(replicasResource, c.ns, replica), &v1beta1.Replica{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Replica), err
}

func (c *FakeReplicas) Update(ctx context.Context, replica *v1beta1.Replica, opts v1.UpdateOptions) (result *v1beta1.Replica, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(replicasResource, c.ns, replica), &v1beta1.Replica{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Replica), err
}

func (c *FakeReplicas) UpdateStatus(ctx context.Context, replica *v1beta1.Replica, opts v1.UpdateOptions) (*v1beta1.Replica, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(replicasResource, "status", c.ns, replica), &v1beta1.Replica{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Replica), err
}

func (c *FakeReplicas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(replicasResource, c.ns, name), &v1beta1.Replica{})

	return err
}

func (c *FakeReplicas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(replicasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ReplicaList{})
	return err
}

func (c *FakeReplicas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Replica, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(replicasResource, c.ns, name, pt, data, subresources...), &v1beta1.Replica{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Replica), err
}
