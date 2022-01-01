package fake

import (
	"context"

	k8scnicncfiov1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeNetworkAttachmentDefinitions struct {
	Fake	*FakeK8sCniCncfIoV1
	ns	string
}

var networkattachmentdefinitionsResource = schema.GroupVersionResource{Group: "k8s.cni.cncf.io", Version: "v1", Resource: "network-attachment-definitions"}

var networkattachmentdefinitionsKind = schema.GroupVersionKind{Group: "k8s.cni.cncf.io", Version: "v1", Kind: "NetworkAttachmentDefinition"}

func (c *FakeNetworkAttachmentDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *k8scnicncfiov1.NetworkAttachmentDefinition, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkattachmentdefinitionsResource, c.ns, name), &k8scnicncfiov1.NetworkAttachmentDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8scnicncfiov1.NetworkAttachmentDefinition), err
}

func (c *FakeNetworkAttachmentDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *k8scnicncfiov1.NetworkAttachmentDefinitionList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkattachmentdefinitionsResource, networkattachmentdefinitionsKind, c.ns, opts), &k8scnicncfiov1.NetworkAttachmentDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &k8scnicncfiov1.NetworkAttachmentDefinitionList{ListMeta: obj.(*k8scnicncfiov1.NetworkAttachmentDefinitionList).ListMeta}
	for _, item := range obj.(*k8scnicncfiov1.NetworkAttachmentDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeNetworkAttachmentDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkattachmentdefinitionsResource, c.ns, opts))

}

func (c *FakeNetworkAttachmentDefinitions) Create(ctx context.Context, networkAttachmentDefinition *k8scnicncfiov1.NetworkAttachmentDefinition, opts v1.CreateOptions) (result *k8scnicncfiov1.NetworkAttachmentDefinition, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkattachmentdefinitionsResource, c.ns, networkAttachmentDefinition), &k8scnicncfiov1.NetworkAttachmentDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8scnicncfiov1.NetworkAttachmentDefinition), err
}

func (c *FakeNetworkAttachmentDefinitions) Update(ctx context.Context, networkAttachmentDefinition *k8scnicncfiov1.NetworkAttachmentDefinition, opts v1.UpdateOptions) (result *k8scnicncfiov1.NetworkAttachmentDefinition, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkattachmentdefinitionsResource, c.ns, networkAttachmentDefinition), &k8scnicncfiov1.NetworkAttachmentDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8scnicncfiov1.NetworkAttachmentDefinition), err
}

func (c *FakeNetworkAttachmentDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkattachmentdefinitionsResource, c.ns, name), &k8scnicncfiov1.NetworkAttachmentDefinition{})

	return err
}

func (c *FakeNetworkAttachmentDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(networkattachmentdefinitionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &k8scnicncfiov1.NetworkAttachmentDefinitionList{})
	return err
}

func (c *FakeNetworkAttachmentDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *k8scnicncfiov1.NetworkAttachmentDefinition, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkattachmentdefinitionsResource, c.ns, name, pt, data, subresources...), &k8scnicncfiov1.NetworkAttachmentDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8scnicncfiov1.NetworkAttachmentDefinition), err
}
