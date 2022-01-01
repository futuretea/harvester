package v1

import (
	v1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	__traceStack()

	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	NetworkAttachmentDefinition() NetworkAttachmentDefinitionController
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

func (c *version) NetworkAttachmentDefinition() NetworkAttachmentDefinitionController {
	__traceStack()

	return NewNetworkAttachmentDefinitionController(schema.GroupVersionKind{Group: "k8s.cni.cncf.io", Version: "v1", Kind: "NetworkAttachmentDefinition"}, "network-attachment-definitions", true, c.controllerFactory)
}
