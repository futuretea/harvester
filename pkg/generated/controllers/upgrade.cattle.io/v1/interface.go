package v1

import (
	"github.com/rancher/lasso/pkg/controller"
	v1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	__traceStack()

	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	Plan() PlanController
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

func (c *version) Plan() PlanController {
	__traceStack()

	return NewPlanController(schema.GroupVersionKind{Group: "upgrade.cattle.io", Version: "v1", Kind: "Plan"}, "plans", true, c.controllerFactory)
}
