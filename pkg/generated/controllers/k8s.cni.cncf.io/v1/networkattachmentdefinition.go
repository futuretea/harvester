package v1

import (
	"context"
	"time"

	v1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
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

type NetworkAttachmentDefinitionHandler func(string, *v1.NetworkAttachmentDefinition) (*v1.NetworkAttachmentDefinition, error)

type NetworkAttachmentDefinitionController interface {
	generic.ControllerMeta
	NetworkAttachmentDefinitionClient

	OnChange(ctx context.Context, name string, sync NetworkAttachmentDefinitionHandler)
	OnRemove(ctx context.Context, name string, sync NetworkAttachmentDefinitionHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() NetworkAttachmentDefinitionCache
}

type NetworkAttachmentDefinitionClient interface {
	Create(*v1.NetworkAttachmentDefinition) (*v1.NetworkAttachmentDefinition, error)
	Update(*v1.NetworkAttachmentDefinition) (*v1.NetworkAttachmentDefinition, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.NetworkAttachmentDefinition, error)
	List(namespace string, opts metav1.ListOptions) (*v1.NetworkAttachmentDefinitionList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NetworkAttachmentDefinition, err error)
}

type NetworkAttachmentDefinitionCache interface {
	Get(namespace, name string) (*v1.NetworkAttachmentDefinition, error)
	List(namespace string, selector labels.Selector) ([]*v1.NetworkAttachmentDefinition, error)

	AddIndexer(indexName string, indexer NetworkAttachmentDefinitionIndexer)
	GetByIndex(indexName, key string) ([]*v1.NetworkAttachmentDefinition, error)
}

type NetworkAttachmentDefinitionIndexer func(obj *v1.NetworkAttachmentDefinition) ([]string, error)

type networkAttachmentDefinitionController struct {
	controller	controller.SharedController
	client		*client.Client
	gvk		schema.GroupVersionKind
	groupResource	schema.GroupResource
}

func NewNetworkAttachmentDefinitionController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) NetworkAttachmentDefinitionController {
	__traceStack()

	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &networkAttachmentDefinitionController{
		controller:	c,
		client:		c.Client(),
		gvk:		gvk,
		groupResource: schema.GroupResource{
			Group:		gvk.Group,
			Resource:	resource,
		},
	}
}

func FromNetworkAttachmentDefinitionHandlerToHandler(sync NetworkAttachmentDefinitionHandler) generic.Handler {
	__traceStack()

	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.NetworkAttachmentDefinition
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.NetworkAttachmentDefinition))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *networkAttachmentDefinitionController) Updater() generic.Updater {
	__traceStack()

	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.NetworkAttachmentDefinition))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateNetworkAttachmentDefinitionDeepCopyOnChange(client NetworkAttachmentDefinitionClient, obj *v1.NetworkAttachmentDefinition, handler func(obj *v1.NetworkAttachmentDefinition) (*v1.NetworkAttachmentDefinition, error)) (*v1.NetworkAttachmentDefinition, error) {
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

func (c *networkAttachmentDefinitionController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *networkAttachmentDefinitionController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *networkAttachmentDefinitionController) OnChange(ctx context.Context, name string, sync NetworkAttachmentDefinitionHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, FromNetworkAttachmentDefinitionHandlerToHandler(sync))
}

func (c *networkAttachmentDefinitionController) OnRemove(ctx context.Context, name string, sync NetworkAttachmentDefinitionHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromNetworkAttachmentDefinitionHandlerToHandler(sync)))
}

func (c *networkAttachmentDefinitionController) Enqueue(namespace, name string) {
	__traceStack()

	c.controller.Enqueue(namespace, name)
}

func (c *networkAttachmentDefinitionController) EnqueueAfter(namespace, name string, duration time.Duration) {
	__traceStack()

	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *networkAttachmentDefinitionController) Informer() cache.SharedIndexInformer {
	__traceStack()

	return c.controller.Informer()
}

func (c *networkAttachmentDefinitionController) GroupVersionKind() schema.GroupVersionKind {
	__traceStack()

	return c.gvk
}

func (c *networkAttachmentDefinitionController) Cache() NetworkAttachmentDefinitionCache {
	__traceStack()

	return &networkAttachmentDefinitionCache{
		indexer:	c.Informer().GetIndexer(),
		resource:	c.groupResource,
	}
}

func (c *networkAttachmentDefinitionController) Create(obj *v1.NetworkAttachmentDefinition) (*v1.NetworkAttachmentDefinition, error) {
	__traceStack()

	result := &v1.NetworkAttachmentDefinition{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *networkAttachmentDefinitionController) Update(obj *v1.NetworkAttachmentDefinition) (*v1.NetworkAttachmentDefinition, error) {
	__traceStack()

	result := &v1.NetworkAttachmentDefinition{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *networkAttachmentDefinitionController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	__traceStack()

	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *networkAttachmentDefinitionController) Get(namespace, name string, options metav1.GetOptions) (*v1.NetworkAttachmentDefinition, error) {
	__traceStack()

	result := &v1.NetworkAttachmentDefinition{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *networkAttachmentDefinitionController) List(namespace string, opts metav1.ListOptions) (*v1.NetworkAttachmentDefinitionList, error) {
	__traceStack()

	result := &v1.NetworkAttachmentDefinitionList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *networkAttachmentDefinitionController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *networkAttachmentDefinitionController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.NetworkAttachmentDefinition, error) {
	__traceStack()

	result := &v1.NetworkAttachmentDefinition{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type networkAttachmentDefinitionCache struct {
	indexer		cache.Indexer
	resource	schema.GroupResource
}

func (c *networkAttachmentDefinitionCache) Get(namespace, name string) (*v1.NetworkAttachmentDefinition, error) {
	__traceStack()

	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.NetworkAttachmentDefinition), nil
}

func (c *networkAttachmentDefinitionCache) List(namespace string, selector labels.Selector) (ret []*v1.NetworkAttachmentDefinition, err error) {
	__traceStack()

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NetworkAttachmentDefinition))
	})

	return ret, err
}

func (c *networkAttachmentDefinitionCache) AddIndexer(indexName string, indexer NetworkAttachmentDefinitionIndexer) {
	__traceStack()

	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.NetworkAttachmentDefinition))
		},
	}))
}

func (c *networkAttachmentDefinitionCache) GetByIndex(indexName, key string) (result []*v1.NetworkAttachmentDefinition, err error) {
	__traceStack()

	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.NetworkAttachmentDefinition, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.NetworkAttachmentDefinition))
	}
	return result, nil
}
