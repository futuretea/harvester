package snapshot

import (
	v1beta1 "github.com/harvester/harvester/pkg/generated/controllers/snapshot.storage.k8s.io/v1beta1"
	"github.com/rancher/lasso/pkg/controller"
)

type Interface interface {
	V1beta1() v1beta1.Interface
}

type group struct {
	controllerFactory controller.SharedControllerFactory
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	__traceStack()

	return &group{
		controllerFactory: controllerFactory,
	}
}

func (g *group) V1beta1() v1beta1.Interface {
	__traceStack()

	return v1beta1.New(g.controllerFactory)
}
