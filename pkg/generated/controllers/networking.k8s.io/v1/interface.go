package v1

import (
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	v1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	__traceStack()

	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	Ingress() IngressController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	__traceStack()

	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) Ingress() IngressController {
	__traceStack()

	return NewIngressController(schema.GroupVersionKind{Group: "networking.k8s.io", Version: "v1", Kind: "Ingress"}, "ingresses", true, c.controllerFactory)
}
