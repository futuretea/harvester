package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
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

type BackingImageDataSourceHandler func(string, *v1beta1.BackingImageDataSource) (*v1beta1.BackingImageDataSource, error)

type BackingImageDataSourceController interface {
	generic.ControllerMeta
	BackingImageDataSourceClient

	OnChange(ctx context.Context, name string, sync BackingImageDataSourceHandler)
	OnRemove(ctx context.Context, name string, sync BackingImageDataSourceHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() BackingImageDataSourceCache
}

type BackingImageDataSourceClient interface {
	Create(*v1beta1.BackingImageDataSource) (*v1beta1.BackingImageDataSource, error)
	Update(*v1beta1.BackingImageDataSource) (*v1beta1.BackingImageDataSource, error)
	UpdateStatus(*v1beta1.BackingImageDataSource) (*v1beta1.BackingImageDataSource, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1beta1.BackingImageDataSource, error)
	List(namespace string, opts metav1.ListOptions) (*v1beta1.BackingImageDataSourceList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.BackingImageDataSource, err error)
}

type BackingImageDataSourceCache interface {
	Get(namespace, name string) (*v1beta1.BackingImageDataSource, error)
	List(namespace string, selector labels.Selector) ([]*v1beta1.BackingImageDataSource, error)

	AddIndexer(indexName string, indexer BackingImageDataSourceIndexer)
	GetByIndex(indexName, key string) ([]*v1beta1.BackingImageDataSource, error)
}

type BackingImageDataSourceIndexer func(obj *v1beta1.BackingImageDataSource) ([]string, error)

type backingImageDataSourceController struct {
	controller	controller.SharedController
	client		*client.Client
	gvk		schema.GroupVersionKind
	groupResource	schema.GroupResource
}

func NewBackingImageDataSourceController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) BackingImageDataSourceController {
	__traceStack()

	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &backingImageDataSourceController{
		controller:	c,
		client:		c.Client(),
		gvk:		gvk,
		groupResource: schema.GroupResource{
			Group:		gvk.Group,
			Resource:	resource,
		},
	}
}

func FromBackingImageDataSourceHandlerToHandler(sync BackingImageDataSourceHandler) generic.Handler {
	__traceStack()

	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1beta1.BackingImageDataSource
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1beta1.BackingImageDataSource))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *backingImageDataSourceController) Updater() generic.Updater {
	__traceStack()

	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1beta1.BackingImageDataSource))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateBackingImageDataSourceDeepCopyOnChange(client BackingImageDataSourceClient, obj *v1beta1.BackingImageDataSource, handler func(obj *v1beta1.BackingImageDataSource) (*v1beta1.BackingImageDataSource, error)) (*v1beta1.BackingImageDataSource, error) {
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

func (c *backingImageDataSourceController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *backingImageDataSourceController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *backingImageDataSourceController) OnChange(ctx context.Context, name string, sync BackingImageDataSourceHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, FromBackingImageDataSourceHandlerToHandler(sync))
}

func (c *backingImageDataSourceController) OnRemove(ctx context.Context, name string, sync BackingImageDataSourceHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromBackingImageDataSourceHandlerToHandler(sync)))
}

func (c *backingImageDataSourceController) Enqueue(namespace, name string) {
	__traceStack()

	c.controller.Enqueue(namespace, name)
}

func (c *backingImageDataSourceController) EnqueueAfter(namespace, name string, duration time.Duration) {
	__traceStack()

	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *backingImageDataSourceController) Informer() cache.SharedIndexInformer {
	__traceStack()

	return c.controller.Informer()
}

func (c *backingImageDataSourceController) GroupVersionKind() schema.GroupVersionKind {
	__traceStack()

	return c.gvk
}

func (c *backingImageDataSourceController) Cache() BackingImageDataSourceCache {
	__traceStack()

	return &backingImageDataSourceCache{
		indexer:	c.Informer().GetIndexer(),
		resource:	c.groupResource,
	}
}

func (c *backingImageDataSourceController) Create(obj *v1beta1.BackingImageDataSource) (*v1beta1.BackingImageDataSource, error) {
	__traceStack()

	result := &v1beta1.BackingImageDataSource{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *backingImageDataSourceController) Update(obj *v1beta1.BackingImageDataSource) (*v1beta1.BackingImageDataSource, error) {
	__traceStack()

	result := &v1beta1.BackingImageDataSource{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *backingImageDataSourceController) UpdateStatus(obj *v1beta1.BackingImageDataSource) (*v1beta1.BackingImageDataSource, error) {
	__traceStack()

	result := &v1beta1.BackingImageDataSource{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *backingImageDataSourceController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	__traceStack()

	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *backingImageDataSourceController) Get(namespace, name string, options metav1.GetOptions) (*v1beta1.BackingImageDataSource, error) {
	__traceStack()

	result := &v1beta1.BackingImageDataSource{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *backingImageDataSourceController) List(namespace string, opts metav1.ListOptions) (*v1beta1.BackingImageDataSourceList, error) {
	__traceStack()

	result := &v1beta1.BackingImageDataSourceList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *backingImageDataSourceController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *backingImageDataSourceController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1beta1.BackingImageDataSource, error) {
	__traceStack()

	result := &v1beta1.BackingImageDataSource{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type backingImageDataSourceCache struct {
	indexer		cache.Indexer
	resource	schema.GroupResource
}

func (c *backingImageDataSourceCache) Get(namespace, name string) (*v1beta1.BackingImageDataSource, error) {
	__traceStack()

	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1beta1.BackingImageDataSource), nil
}

func (c *backingImageDataSourceCache) List(namespace string, selector labels.Selector) (ret []*v1beta1.BackingImageDataSource, err error) {
	__traceStack()

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.BackingImageDataSource))
	})

	return ret, err
}

func (c *backingImageDataSourceCache) AddIndexer(indexName string, indexer BackingImageDataSourceIndexer) {
	__traceStack()

	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1beta1.BackingImageDataSource))
		},
	}))
}

func (c *backingImageDataSourceCache) GetByIndex(indexName, key string) (result []*v1beta1.BackingImageDataSource, err error) {
	__traceStack()

	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1beta1.BackingImageDataSource, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1beta1.BackingImageDataSource))
	}
	return result, nil
}

type BackingImageDataSourceStatusHandler func(obj *v1beta1.BackingImageDataSource, status v1beta1.BackingImageDataSourceStatus) (v1beta1.BackingImageDataSourceStatus, error)

type BackingImageDataSourceGeneratingHandler func(obj *v1beta1.BackingImageDataSource, status v1beta1.BackingImageDataSourceStatus) ([]runtime.Object, v1beta1.BackingImageDataSourceStatus, error)

func RegisterBackingImageDataSourceStatusHandler(ctx context.Context, controller BackingImageDataSourceController, condition condition.Cond, name string, handler BackingImageDataSourceStatusHandler) {
	__traceStack()

	statusHandler := &backingImageDataSourceStatusHandler{
		client:		controller,
		condition:	condition,
		handler:	handler,
	}
	controller.AddGenericHandler(ctx, name, FromBackingImageDataSourceHandlerToHandler(statusHandler.sync))
}

func RegisterBackingImageDataSourceGeneratingHandler(ctx context.Context, controller BackingImageDataSourceController, apply apply.Apply,
	condition condition.Cond, name string, handler BackingImageDataSourceGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	__traceStack()

	statusHandler := &backingImageDataSourceGeneratingHandler{
		BackingImageDataSourceGeneratingHandler:	handler,
		apply:						apply,
		name:						name,
		gvk:						controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterBackingImageDataSourceStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type backingImageDataSourceStatusHandler struct {
	client		BackingImageDataSourceClient
	condition	condition.Cond
	handler		BackingImageDataSourceStatusHandler
}

func (a *backingImageDataSourceStatusHandler) sync(key string, obj *v1beta1.BackingImageDataSource) (*v1beta1.BackingImageDataSource, error) {
	__traceStack()

	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {

		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {

			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type backingImageDataSourceGeneratingHandler struct {
	BackingImageDataSourceGeneratingHandler
	apply	apply.Apply
	opts	generic.GeneratingHandlerOptions
	gvk	schema.GroupVersionKind
	name	string
}

func (a *backingImageDataSourceGeneratingHandler) Remove(key string, obj *v1beta1.BackingImageDataSource) (*v1beta1.BackingImageDataSource, error) {
	__traceStack()

	if obj != nil {
		return obj, nil
	}

	obj = &v1beta1.BackingImageDataSource{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *backingImageDataSourceGeneratingHandler) Handle(obj *v1beta1.BackingImageDataSource, status v1beta1.BackingImageDataSourceStatus) (v1beta1.BackingImageDataSourceStatus, error) {
	__traceStack()

	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.BackingImageDataSourceGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
