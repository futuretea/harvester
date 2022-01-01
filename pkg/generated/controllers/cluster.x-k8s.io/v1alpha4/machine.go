package v1alpha4

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
	v1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

type MachineHandler func(string, *v1alpha4.Machine) (*v1alpha4.Machine, error)

type MachineController interface {
	generic.ControllerMeta
	MachineClient

	OnChange(ctx context.Context, name string, sync MachineHandler)
	OnRemove(ctx context.Context, name string, sync MachineHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() MachineCache
}

type MachineClient interface {
	Create(*v1alpha4.Machine) (*v1alpha4.Machine, error)
	Update(*v1alpha4.Machine) (*v1alpha4.Machine, error)
	UpdateStatus(*v1alpha4.Machine) (*v1alpha4.Machine, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1alpha4.Machine, error)
	List(namespace string, opts metav1.ListOptions) (*v1alpha4.MachineList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha4.Machine, err error)
}

type MachineCache interface {
	Get(namespace, name string) (*v1alpha4.Machine, error)
	List(namespace string, selector labels.Selector) ([]*v1alpha4.Machine, error)

	AddIndexer(indexName string, indexer MachineIndexer)
	GetByIndex(indexName, key string) ([]*v1alpha4.Machine, error)
}

type MachineIndexer func(obj *v1alpha4.Machine) ([]string, error)

type machineController struct {
	controller	controller.SharedController
	client		*client.Client
	gvk		schema.GroupVersionKind
	groupResource	schema.GroupResource
}

func NewMachineController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) MachineController {
	__traceStack()

	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &machineController{
		controller:	c,
		client:		c.Client(),
		gvk:		gvk,
		groupResource: schema.GroupResource{
			Group:		gvk.Group,
			Resource:	resource,
		},
	}
}

func FromMachineHandlerToHandler(sync MachineHandler) generic.Handler {
	__traceStack()

	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1alpha4.Machine
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1alpha4.Machine))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *machineController) Updater() generic.Updater {
	__traceStack()

	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1alpha4.Machine))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateMachineDeepCopyOnChange(client MachineClient, obj *v1alpha4.Machine, handler func(obj *v1alpha4.Machine) (*v1alpha4.Machine, error)) (*v1alpha4.Machine, error) {
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

func (c *machineController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *machineController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *machineController) OnChange(ctx context.Context, name string, sync MachineHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, FromMachineHandlerToHandler(sync))
}

func (c *machineController) OnRemove(ctx context.Context, name string, sync MachineHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromMachineHandlerToHandler(sync)))
}

func (c *machineController) Enqueue(namespace, name string) {
	__traceStack()

	c.controller.Enqueue(namespace, name)
}

func (c *machineController) EnqueueAfter(namespace, name string, duration time.Duration) {
	__traceStack()

	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *machineController) Informer() cache.SharedIndexInformer {
	__traceStack()

	return c.controller.Informer()
}

func (c *machineController) GroupVersionKind() schema.GroupVersionKind {
	__traceStack()

	return c.gvk
}

func (c *machineController) Cache() MachineCache {
	__traceStack()

	return &machineCache{
		indexer:	c.Informer().GetIndexer(),
		resource:	c.groupResource,
	}
}

func (c *machineController) Create(obj *v1alpha4.Machine) (*v1alpha4.Machine, error) {
	__traceStack()

	result := &v1alpha4.Machine{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *machineController) Update(obj *v1alpha4.Machine) (*v1alpha4.Machine, error) {
	__traceStack()

	result := &v1alpha4.Machine{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *machineController) UpdateStatus(obj *v1alpha4.Machine) (*v1alpha4.Machine, error) {
	__traceStack()

	result := &v1alpha4.Machine{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *machineController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	__traceStack()

	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *machineController) Get(namespace, name string, options metav1.GetOptions) (*v1alpha4.Machine, error) {
	__traceStack()

	result := &v1alpha4.Machine{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *machineController) List(namespace string, opts metav1.ListOptions) (*v1alpha4.MachineList, error) {
	__traceStack()

	result := &v1alpha4.MachineList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *machineController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *machineController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1alpha4.Machine, error) {
	__traceStack()

	result := &v1alpha4.Machine{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type machineCache struct {
	indexer		cache.Indexer
	resource	schema.GroupResource
}

func (c *machineCache) Get(namespace, name string) (*v1alpha4.Machine, error) {
	__traceStack()

	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1alpha4.Machine), nil
}

func (c *machineCache) List(namespace string, selector labels.Selector) (ret []*v1alpha4.Machine, err error) {
	__traceStack()

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha4.Machine))
	})

	return ret, err
}

func (c *machineCache) AddIndexer(indexName string, indexer MachineIndexer) {
	__traceStack()

	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1alpha4.Machine))
		},
	}))
}

func (c *machineCache) GetByIndex(indexName, key string) (result []*v1alpha4.Machine, err error) {
	__traceStack()

	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1alpha4.Machine, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1alpha4.Machine))
	}
	return result, nil
}

type MachineStatusHandler func(obj *v1alpha4.Machine, status v1alpha4.MachineStatus) (v1alpha4.MachineStatus, error)

type MachineGeneratingHandler func(obj *v1alpha4.Machine, status v1alpha4.MachineStatus) ([]runtime.Object, v1alpha4.MachineStatus, error)

func RegisterMachineStatusHandler(ctx context.Context, controller MachineController, condition condition.Cond, name string, handler MachineStatusHandler) {
	__traceStack()

	statusHandler := &machineStatusHandler{
		client:		controller,
		condition:	condition,
		handler:	handler,
	}
	controller.AddGenericHandler(ctx, name, FromMachineHandlerToHandler(statusHandler.sync))
}

func RegisterMachineGeneratingHandler(ctx context.Context, controller MachineController, apply apply.Apply,
	condition condition.Cond, name string, handler MachineGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	__traceStack()

	statusHandler := &machineGeneratingHandler{
		MachineGeneratingHandler:	handler,
		apply:				apply,
		name:				name,
		gvk:				controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterMachineStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type machineStatusHandler struct {
	client		MachineClient
	condition	condition.Cond
	handler		MachineStatusHandler
}

func (a *machineStatusHandler) sync(key string, obj *v1alpha4.Machine) (*v1alpha4.Machine, error) {
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

type machineGeneratingHandler struct {
	MachineGeneratingHandler
	apply	apply.Apply
	opts	generic.GeneratingHandlerOptions
	gvk	schema.GroupVersionKind
	name	string
}

func (a *machineGeneratingHandler) Remove(key string, obj *v1alpha4.Machine) (*v1alpha4.Machine, error) {
	__traceStack()

	if obj != nil {
		return obj, nil
	}

	obj = &v1alpha4.Machine{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *machineGeneratingHandler) Handle(obj *v1alpha4.Machine, status v1alpha4.MachineStatus) (v1alpha4.MachineStatus, error) {
	__traceStack()

	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.MachineGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
