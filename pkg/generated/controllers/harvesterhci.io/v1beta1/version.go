package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type VersionHandler func(string, *v1beta1.Version) (*v1beta1.Version, error)

type VersionController interface {
	generic.ControllerMeta
	VersionClient

	OnChange(ctx context.Context, name string, sync VersionHandler)
	OnRemove(ctx context.Context, name string, sync VersionHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() VersionCache
}

type VersionClient interface {
	Create(*v1beta1.Version) (*v1beta1.Version, error)
	Update(*v1beta1.Version) (*v1beta1.Version, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1beta1.Version, error)
	List(namespace string, opts metav1.ListOptions) (*v1beta1.VersionList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Version, err error)
}

type VersionCache interface {
	Get(namespace, name string) (*v1beta1.Version, error)
	List(namespace string, selector labels.Selector) ([]*v1beta1.Version, error)

	AddIndexer(indexName string, indexer VersionIndexer)
	GetByIndex(indexName, key string) ([]*v1beta1.Version, error)
}

type VersionIndexer func(obj *v1beta1.Version) ([]string, error)

type versionController struct {
	controller	controller.SharedController
	client		*client.Client
	gvk		schema.GroupVersionKind
	groupResource	schema.GroupResource
}

func NewVersionController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) VersionController {
	__traceStack()

	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &versionController{
		controller:	c,
		client:		c.Client(),
		gvk:		gvk,
		groupResource: schema.GroupResource{
			Group:		gvk.Group,
			Resource:	resource,
		},
	}
}

func FromVersionHandlerToHandler(sync VersionHandler) generic.Handler {
	__traceStack()

	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1beta1.Version
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1beta1.Version))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *versionController) Updater() generic.Updater {
	__traceStack()

	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1beta1.Version))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateVersionDeepCopyOnChange(client VersionClient, obj *v1beta1.Version, handler func(obj *v1beta1.Version) (*v1beta1.Version, error)) (*v1beta1.Version, error) {
	__traceStack()

	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *versionController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *versionController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *versionController) OnChange(ctx context.Context, name string, sync VersionHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, FromVersionHandlerToHandler(sync))
}

func (c *versionController) OnRemove(ctx context.Context, name string, sync VersionHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromVersionHandlerToHandler(sync)))
}

func (c *versionController) Enqueue(namespace, name string) {
	__traceStack()

	c.controller.Enqueue(namespace, name)
}

func (c *versionController) EnqueueAfter(namespace, name string, duration time.Duration) {
	__traceStack()

	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *versionController) Informer() cache.SharedIndexInformer {
	__traceStack()

	return c.controller.Informer()
}

func (c *versionController) GroupVersionKind() schema.GroupVersionKind {
	__traceStack()

	return c.gvk
}

func (c *versionController) Cache() VersionCache {
	__traceStack()

	return &versionCache{
		indexer:	c.Informer().GetIndexer(),
		resource:	c.groupResource,
	}
}

func (c *versionController) Create(obj *v1beta1.Version) (*v1beta1.Version, error) {
	__traceStack()

	result := &v1beta1.Version{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *versionController) Update(obj *v1beta1.Version) (*v1beta1.Version, error) {
	__traceStack()

	result := &v1beta1.Version{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *versionController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	__traceStack()

	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *versionController) Get(namespace, name string, options metav1.GetOptions) (*v1beta1.Version, error) {
	__traceStack()

	result := &v1beta1.Version{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *versionController) List(namespace string, opts metav1.ListOptions) (*v1beta1.VersionList, error) {
	__traceStack()

	result := &v1beta1.VersionList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *versionController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *versionController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1beta1.Version, error) {
	__traceStack()

	result := &v1beta1.Version{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type versionCache struct {
	indexer		cache.Indexer
	resource	schema.GroupResource
}

func (c *versionCache) Get(namespace, name string) (*v1beta1.Version, error) {
	__traceStack()

	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1beta1.Version), nil
}

func (c *versionCache) List(namespace string, selector labels.Selector) (ret []*v1beta1.Version, err error) {
	__traceStack()

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Version))
	})

	return ret, err
}

func (c *versionCache) AddIndexer(indexName string, indexer VersionIndexer) {
	__traceStack()

	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1beta1.Version))
		},
	}))
}

func (c *versionCache) GetByIndex(indexName, key string) (result []*v1beta1.Version, err error) {
	__traceStack()

	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1beta1.Version, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1beta1.Version))
	}
	return result, nil
}
