package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type PreferencesGetter interface {
	Preferences(namespace string) PreferenceInterface
}

type PreferenceInterface interface {
	Create(ctx context.Context, preference *v1beta1.Preference, opts v1.CreateOptions) (*v1beta1.Preference, error)
	Update(ctx context.Context, preference *v1beta1.Preference, opts v1.UpdateOptions) (*v1beta1.Preference, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Preference, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.PreferenceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Preference, err error)
	PreferenceExpansion
}

type preferences struct {
	client	rest.Interface
	ns	string
}

func newPreferences(c *HarvesterhciV1beta1Client, namespace string) *preferences {
	__traceStack()

	return &preferences{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *preferences) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Preference, err error) {
	__traceStack()

	result = &v1beta1.Preference{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("preferences").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *preferences) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PreferenceList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.PreferenceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("preferences").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *preferences) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("preferences").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *preferences) Create(ctx context.Context, preference *v1beta1.Preference, opts v1.CreateOptions) (result *v1beta1.Preference, err error) {
	__traceStack()

	result = &v1beta1.Preference{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("preferences").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(preference).
		Do(ctx).
		Into(result)
	return
}

func (c *preferences) Update(ctx context.Context, preference *v1beta1.Preference, opts v1.UpdateOptions) (result *v1beta1.Preference, err error) {
	__traceStack()

	result = &v1beta1.Preference{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("preferences").
		Name(preference.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(preference).
		Do(ctx).
		Into(result)
	return
}

func (c *preferences) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("preferences").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *preferences) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("preferences").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *preferences) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Preference, err error) {
	__traceStack()

	result = &v1beta1.Preference{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("preferences").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
