package longhorn

import (
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/client-go/rest"
)

type Factory struct {
	*generic.Factory
}

func NewFactoryFromConfigOrDie(config *rest.Config) *Factory {
	__traceStack()

	f, err := NewFactoryFromConfig(config)
	if err != nil {
		panic(err)
	}
	return f
}

func NewFactoryFromConfig(config *rest.Config) (*Factory, error) {
	__traceStack()

	return NewFactoryFromConfigWithOptions(config, nil)
}

func NewFactoryFromConfigWithNamespace(config *rest.Config, namespace string) (*Factory, error) {
	__traceStack()

	return NewFactoryFromConfigWithOptions(config, &FactoryOptions{
		Namespace: namespace,
	})
}

type FactoryOptions = generic.FactoryOptions

func NewFactoryFromConfigWithOptions(config *rest.Config, opts *FactoryOptions) (*Factory, error) {
	__traceStack()

	f, err := generic.NewFactoryFromConfigWithOptions(config, opts)
	return &Factory{
		Factory: f,
	}, err
}

func NewFactoryFromConfigWithOptionsOrDie(config *rest.Config, opts *FactoryOptions) *Factory {
	__traceStack()

	f, err := NewFactoryFromConfigWithOptions(config, opts)
	if err != nil {
		panic(err)
	}
	return f
}

func (c *Factory) Longhorn() Interface {
	__traceStack()

	return New(c.ControllerFactory())
}
