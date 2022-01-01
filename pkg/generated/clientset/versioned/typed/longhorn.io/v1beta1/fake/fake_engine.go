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

type FakeEngines struct {
	Fake	*FakeLonghornV1beta1
	ns	string
}

var enginesResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "engines"}

var enginesKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Engine"}

func (c *FakeEngines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Engine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(enginesResource, c.ns, name), &v1beta1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Engine), err
}

func (c *FakeEngines) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.EngineList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(enginesResource, enginesKind, c.ns, opts), &v1beta1.EngineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.EngineList{ListMeta: obj.(*v1beta1.EngineList).ListMeta}
	for _, item := range obj.(*v1beta1.EngineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeEngines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(enginesResource, c.ns, opts))

}

func (c *FakeEngines) Create(ctx context.Context, engine *v1beta1.Engine, opts v1.CreateOptions) (result *v1beta1.Engine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(enginesResource, c.ns, engine), &v1beta1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Engine), err
}

func (c *FakeEngines) Update(ctx context.Context, engine *v1beta1.Engine, opts v1.UpdateOptions) (result *v1beta1.Engine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(enginesResource, c.ns, engine), &v1beta1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Engine), err
}

func (c *FakeEngines) UpdateStatus(ctx context.Context, engine *v1beta1.Engine, opts v1.UpdateOptions) (*v1beta1.Engine, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(enginesResource, "status", c.ns, engine), &v1beta1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Engine), err
}

func (c *FakeEngines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(enginesResource, c.ns, name), &v1beta1.Engine{})

	return err
}

func (c *FakeEngines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(enginesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.EngineList{})
	return err
}

func (c *FakeEngines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Engine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(enginesResource, c.ns, name, pt, data, subresources...), &v1beta1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Engine), err
}
