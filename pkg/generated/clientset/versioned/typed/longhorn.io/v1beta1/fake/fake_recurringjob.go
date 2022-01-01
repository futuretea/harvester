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

type FakeRecurringJobs struct {
	Fake	*FakeLonghornV1beta1
	ns	string
}

var recurringjobsResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "recurringjobs"}

var recurringjobsKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "RecurringJob"}

func (c *FakeRecurringJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.RecurringJob, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(recurringjobsResource, c.ns, name), &v1beta1.RecurringJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecurringJob), err
}

func (c *FakeRecurringJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.RecurringJobList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(recurringjobsResource, recurringjobsKind, c.ns, opts), &v1beta1.RecurringJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.RecurringJobList{ListMeta: obj.(*v1beta1.RecurringJobList).ListMeta}
	for _, item := range obj.(*v1beta1.RecurringJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeRecurringJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(recurringjobsResource, c.ns, opts))

}

func (c *FakeRecurringJobs) Create(ctx context.Context, recurringJob *v1beta1.RecurringJob, opts v1.CreateOptions) (result *v1beta1.RecurringJob, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(recurringjobsResource, c.ns, recurringJob), &v1beta1.RecurringJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecurringJob), err
}

func (c *FakeRecurringJobs) Update(ctx context.Context, recurringJob *v1beta1.RecurringJob, opts v1.UpdateOptions) (result *v1beta1.RecurringJob, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(recurringjobsResource, c.ns, recurringJob), &v1beta1.RecurringJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecurringJob), err
}

func (c *FakeRecurringJobs) UpdateStatus(ctx context.Context, recurringJob *v1beta1.RecurringJob, opts v1.UpdateOptions) (*v1beta1.RecurringJob, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(recurringjobsResource, "status", c.ns, recurringJob), &v1beta1.RecurringJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecurringJob), err
}

func (c *FakeRecurringJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(recurringjobsResource, c.ns, name), &v1beta1.RecurringJob{})

	return err
}

func (c *FakeRecurringJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(recurringjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.RecurringJobList{})
	return err
}

func (c *FakeRecurringJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.RecurringJob, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(recurringjobsResource, c.ns, name, pt, data, subresources...), &v1beta1.RecurringJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecurringJob), err
}
