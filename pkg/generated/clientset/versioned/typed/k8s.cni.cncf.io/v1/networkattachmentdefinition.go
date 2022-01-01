package v1

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type NetworkAttachmentDefinitionsGetter interface {
	NetworkAttachmentDefinitions(namespace string) NetworkAttachmentDefinitionInterface
}

type NetworkAttachmentDefinitionInterface interface {
	Create(ctx context.Context, networkAttachmentDefinition *v1.NetworkAttachmentDefinition, opts metav1.CreateOptions) (*v1.NetworkAttachmentDefinition, error)
	Update(ctx context.Context, networkAttachmentDefinition *v1.NetworkAttachmentDefinition, opts metav1.UpdateOptions) (*v1.NetworkAttachmentDefinition, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.NetworkAttachmentDefinition, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NetworkAttachmentDefinitionList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkAttachmentDefinition, err error)
	NetworkAttachmentDefinitionExpansion
}

type networkAttachmentDefinitions struct {
	client	rest.Interface
	ns	string
}

func newNetworkAttachmentDefinitions(c *K8sCniCncfIoV1Client, namespace string) *networkAttachmentDefinitions {
	__traceStack()

	return &networkAttachmentDefinitions{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *networkAttachmentDefinitions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NetworkAttachmentDefinition, err error) {
	__traceStack()

	result = &v1.NetworkAttachmentDefinition{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *networkAttachmentDefinitions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NetworkAttachmentDefinitionList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NetworkAttachmentDefinitionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *networkAttachmentDefinitions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *networkAttachmentDefinitions) Create(ctx context.Context, networkAttachmentDefinition *v1.NetworkAttachmentDefinition, opts metav1.CreateOptions) (result *v1.NetworkAttachmentDefinition, err error) {
	__traceStack()

	result = &v1.NetworkAttachmentDefinition{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkAttachmentDefinition).
		Do(ctx).
		Into(result)
	return
}

func (c *networkAttachmentDefinitions) Update(ctx context.Context, networkAttachmentDefinition *v1.NetworkAttachmentDefinition, opts metav1.UpdateOptions) (result *v1.NetworkAttachmentDefinition, err error) {
	__traceStack()

	result = &v1.NetworkAttachmentDefinition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		Name(networkAttachmentDefinition.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkAttachmentDefinition).
		Do(ctx).
		Into(result)
	return
}

func (c *networkAttachmentDefinitions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *networkAttachmentDefinitions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *networkAttachmentDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkAttachmentDefinition, err error) {
	__traceStack()

	result = &v1.NetworkAttachmentDefinition{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
