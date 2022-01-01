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

type SettingsGetter interface {
	Settings() SettingInterface
}

type SettingInterface interface {
	Create(ctx context.Context, setting *v1beta1.Setting, opts v1.CreateOptions) (*v1beta1.Setting, error)
	Update(ctx context.Context, setting *v1beta1.Setting, opts v1.UpdateOptions) (*v1beta1.Setting, error)
	UpdateStatus(ctx context.Context, setting *v1beta1.Setting, opts v1.UpdateOptions) (*v1beta1.Setting, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Setting, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.SettingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Setting, err error)
	SettingExpansion
}

type settings struct {
	client rest.Interface
}

func newSettings(c *HarvesterhciV1beta1Client) *settings {
	__traceStack()

	return &settings{
		client: c.RESTClient(),
	}
}

func (c *settings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Setting, err error) {
	__traceStack()

	result = &v1beta1.Setting{}
	err = c.client.Get().
		Resource("settings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *settings) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SettingList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.SettingList{}
	err = c.client.Get().
		Resource("settings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *settings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("settings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *settings) Create(ctx context.Context, setting *v1beta1.Setting, opts v1.CreateOptions) (result *v1beta1.Setting, err error) {
	__traceStack()

	result = &v1beta1.Setting{}
	err = c.client.Post().
		Resource("settings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(setting).
		Do(ctx).
		Into(result)
	return
}

func (c *settings) Update(ctx context.Context, setting *v1beta1.Setting, opts v1.UpdateOptions) (result *v1beta1.Setting, err error) {
	__traceStack()

	result = &v1beta1.Setting{}
	err = c.client.Put().
		Resource("settings").
		Name(setting.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(setting).
		Do(ctx).
		Into(result)
	return
}

func (c *settings) UpdateStatus(ctx context.Context, setting *v1beta1.Setting, opts v1.UpdateOptions) (result *v1beta1.Setting, err error) {
	__traceStack()

	result = &v1beta1.Setting{}
	err = c.client.Put().
		Resource("settings").
		Name(setting.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(setting).
		Do(ctx).
		Into(result)
	return
}

func (c *settings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Resource("settings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *settings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("settings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *settings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Setting, err error) {
	__traceStack()

	result = &v1beta1.Setting{}
	err = c.client.Patch(pt).
		Resource("settings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
