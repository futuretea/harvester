package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	apiv1 "kubevirt.io/client-go/api/v1"
)

type FakeKubeVirts struct {
	Fake	*FakeKubevirtV1
	ns	string
}

var kubevirtsResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "kubevirts"}

var kubevirtsKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "KubeVirt"}

func (c *FakeKubeVirts) Get(ctx context.Context, name string, options v1.GetOptions) (result *apiv1.KubeVirt, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubevirtsResource, c.ns, name), &apiv1.KubeVirt{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.KubeVirt), err
}

func (c *FakeKubeVirts) List(ctx context.Context, opts v1.ListOptions) (result *apiv1.KubeVirtList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubevirtsResource, kubevirtsKind, c.ns, opts), &apiv1.KubeVirtList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apiv1.KubeVirtList{ListMeta: obj.(*apiv1.KubeVirtList).ListMeta}
	for _, item := range obj.(*apiv1.KubeVirtList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeKubeVirts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubevirtsResource, c.ns, opts))

}

func (c *FakeKubeVirts) Create(ctx context.Context, kubeVirt *apiv1.KubeVirt, opts v1.CreateOptions) (result *apiv1.KubeVirt, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubevirtsResource, c.ns, kubeVirt), &apiv1.KubeVirt{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.KubeVirt), err
}

func (c *FakeKubeVirts) Update(ctx context.Context, kubeVirt *apiv1.KubeVirt, opts v1.UpdateOptions) (result *apiv1.KubeVirt, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubevirtsResource, c.ns, kubeVirt), &apiv1.KubeVirt{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.KubeVirt), err
}

func (c *FakeKubeVirts) UpdateStatus(ctx context.Context, kubeVirt *apiv1.KubeVirt, opts v1.UpdateOptions) (*apiv1.KubeVirt, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubevirtsResource, "status", c.ns, kubeVirt), &apiv1.KubeVirt{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.KubeVirt), err
}

func (c *FakeKubeVirts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubevirtsResource, c.ns, name), &apiv1.KubeVirt{})

	return err
}

func (c *FakeKubeVirts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(kubevirtsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apiv1.KubeVirtList{})
	return err
}

func (c *FakeKubeVirts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiv1.KubeVirt, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubevirtsResource, c.ns, name, pt, data, subresources...), &apiv1.KubeVirt{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.KubeVirt), err
}
