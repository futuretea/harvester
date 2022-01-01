package v1

import (
	"context"
	"time"

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
	v1 "kubevirt.io/client-go/api/v1"
)

type VirtualMachineHandler func(string, *v1.VirtualMachine) (*v1.VirtualMachine, error)

type VirtualMachineController interface {
	generic.ControllerMeta
	VirtualMachineClient

	OnChange(ctx context.Context, name string, sync VirtualMachineHandler)
	OnRemove(ctx context.Context, name string, sync VirtualMachineHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() VirtualMachineCache
}

type VirtualMachineClient interface {
	Create(*v1.VirtualMachine) (*v1.VirtualMachine, error)
	Update(*v1.VirtualMachine) (*v1.VirtualMachine, error)
	UpdateStatus(*v1.VirtualMachine) (*v1.VirtualMachine, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.VirtualMachine, error)
	List(namespace string, opts metav1.ListOptions) (*v1.VirtualMachineList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.VirtualMachine, err error)
}

type VirtualMachineCache interface {
	Get(namespace, name string) (*v1.VirtualMachine, error)
	List(namespace string, selector labels.Selector) ([]*v1.VirtualMachine, error)

	AddIndexer(indexName string, indexer VirtualMachineIndexer)
	GetByIndex(indexName, key string) ([]*v1.VirtualMachine, error)
}

type VirtualMachineIndexer func(obj *v1.VirtualMachine) ([]string, error)

type virtualMachineController struct {
	controller	controller.SharedController
	client		*client.Client
	gvk		schema.GroupVersionKind
	groupResource	schema.GroupResource
}

func NewVirtualMachineController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) VirtualMachineController {
	__traceStack()

	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &virtualMachineController{
		controller:	c,
		client:		c.Client(),
		gvk:		gvk,
		groupResource: schema.GroupResource{
			Group:		gvk.Group,
			Resource:	resource,
		},
	}
}

func FromVirtualMachineHandlerToHandler(sync VirtualMachineHandler) generic.Handler {
	__traceStack()

	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.VirtualMachine
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.VirtualMachine))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *virtualMachineController) Updater() generic.Updater {
	__traceStack()

	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.VirtualMachine))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateVirtualMachineDeepCopyOnChange(client VirtualMachineClient, obj *v1.VirtualMachine, handler func(obj *v1.VirtualMachine) (*v1.VirtualMachine, error)) (*v1.VirtualMachine, error) {
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

func (c *virtualMachineController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *virtualMachineController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *virtualMachineController) OnChange(ctx context.Context, name string, sync VirtualMachineHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, FromVirtualMachineHandlerToHandler(sync))
}

func (c *virtualMachineController) OnRemove(ctx context.Context, name string, sync VirtualMachineHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromVirtualMachineHandlerToHandler(sync)))
}

func (c *virtualMachineController) Enqueue(namespace, name string) {
	__traceStack()

	c.controller.Enqueue(namespace, name)
}

func (c *virtualMachineController) EnqueueAfter(namespace, name string, duration time.Duration) {
	__traceStack()

	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *virtualMachineController) Informer() cache.SharedIndexInformer {
	__traceStack()

	return c.controller.Informer()
}

func (c *virtualMachineController) GroupVersionKind() schema.GroupVersionKind {
	__traceStack()

	return c.gvk
}

func (c *virtualMachineController) Cache() VirtualMachineCache {
	__traceStack()

	return &virtualMachineCache{
		indexer:	c.Informer().GetIndexer(),
		resource:	c.groupResource,
	}
}

func (c *virtualMachineController) Create(obj *v1.VirtualMachine) (*v1.VirtualMachine, error) {
	__traceStack()

	result := &v1.VirtualMachine{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *virtualMachineController) Update(obj *v1.VirtualMachine) (*v1.VirtualMachine, error) {
	__traceStack()

	result := &v1.VirtualMachine{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *virtualMachineController) UpdateStatus(obj *v1.VirtualMachine) (*v1.VirtualMachine, error) {
	__traceStack()

	result := &v1.VirtualMachine{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *virtualMachineController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	__traceStack()

	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *virtualMachineController) Get(namespace, name string, options metav1.GetOptions) (*v1.VirtualMachine, error) {
	__traceStack()

	result := &v1.VirtualMachine{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *virtualMachineController) List(namespace string, opts metav1.ListOptions) (*v1.VirtualMachineList, error) {
	__traceStack()

	result := &v1.VirtualMachineList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *virtualMachineController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *virtualMachineController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.VirtualMachine, error) {
	__traceStack()

	result := &v1.VirtualMachine{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type virtualMachineCache struct {
	indexer		cache.Indexer
	resource	schema.GroupResource
}

func (c *virtualMachineCache) Get(namespace, name string) (*v1.VirtualMachine, error) {
	__traceStack()

	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.VirtualMachine), nil
}

func (c *virtualMachineCache) List(namespace string, selector labels.Selector) (ret []*v1.VirtualMachine, err error) {
	__traceStack()

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualMachine))
	})

	return ret, err
}

func (c *virtualMachineCache) AddIndexer(indexName string, indexer VirtualMachineIndexer) {
	__traceStack()

	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.VirtualMachine))
		},
	}))
}

func (c *virtualMachineCache) GetByIndex(indexName, key string) (result []*v1.VirtualMachine, err error) {
	__traceStack()

	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.VirtualMachine, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.VirtualMachine))
	}
	return result, nil
}

type VirtualMachineStatusHandler func(obj *v1.VirtualMachine, status v1.VirtualMachineStatus) (v1.VirtualMachineStatus, error)

type VirtualMachineGeneratingHandler func(obj *v1.VirtualMachine, status v1.VirtualMachineStatus) ([]runtime.Object, v1.VirtualMachineStatus, error)

func RegisterVirtualMachineStatusHandler(ctx context.Context, controller VirtualMachineController, condition condition.Cond, name string, handler VirtualMachineStatusHandler) {
	__traceStack()

	statusHandler := &virtualMachineStatusHandler{
		client:		controller,
		condition:	condition,
		handler:	handler,
	}
	controller.AddGenericHandler(ctx, name, FromVirtualMachineHandlerToHandler(statusHandler.sync))
}

func RegisterVirtualMachineGeneratingHandler(ctx context.Context, controller VirtualMachineController, apply apply.Apply,
	condition condition.Cond, name string, handler VirtualMachineGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	__traceStack()

	statusHandler := &virtualMachineGeneratingHandler{
		VirtualMachineGeneratingHandler:	handler,
		apply:					apply,
		name:					name,
		gvk:					controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterVirtualMachineStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type virtualMachineStatusHandler struct {
	client		VirtualMachineClient
	condition	condition.Cond
	handler		VirtualMachineStatusHandler
}

func (a *virtualMachineStatusHandler) sync(key string, obj *v1.VirtualMachine) (*v1.VirtualMachine, error) {
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

type virtualMachineGeneratingHandler struct {
	VirtualMachineGeneratingHandler
	apply	apply.Apply
	opts	generic.GeneratingHandlerOptions
	gvk	schema.GroupVersionKind
	name	string
}

func (a *virtualMachineGeneratingHandler) Remove(key string, obj *v1.VirtualMachine) (*v1.VirtualMachine, error) {
	__traceStack()

	if obj != nil {
		return obj, nil
	}

	obj = &v1.VirtualMachine{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *virtualMachineGeneratingHandler) Handle(obj *v1.VirtualMachine, status v1.VirtualMachineStatus) (v1.VirtualMachineStatus, error) {
	__traceStack()

	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.VirtualMachineGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
