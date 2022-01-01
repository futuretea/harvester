package v1alpha4

import (
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
	v1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

func init() {
	__traceStack()

	schemes.Register(v1alpha4.AddToScheme)
}

type Interface interface {
	Machine() MachineController
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

func (c *version) Machine() MachineController {
	__traceStack()

	return NewMachineController(schema.GroupVersionKind{Group: "cluster.x-k8s.io", Version: "v1alpha4", Kind: "Machine"}, "machines", true, c.controllerFactory)
}
