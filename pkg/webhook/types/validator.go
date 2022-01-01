package types

import (
	"k8s.io/apimachinery/pkg/runtime"
)

type Validator interface {
	Create(request *Request, newObj runtime.Object) error

	Update(request *Request, oldObj runtime.Object, newObj runtime.Object) error

	Delete(request *Request, oldObj runtime.Object) error

	Connect(request *Request, newObj runtime.Object) error

	Resource() Resource
}

type ValidatorAdapter struct {
	validator Validator
}

func NewValidatorAdapter(validator Validator) Mutator {
	__traceStack()

	return &ValidatorAdapter{validator: validator}
}

func (c *ValidatorAdapter) Create(request *Request, newObj runtime.Object) (PatchOps, error) {
	__traceStack()

	return nil, c.validator.Create(request, newObj)
}

func (c *ValidatorAdapter) Update(request *Request, oldObj runtime.Object, newObj runtime.Object) (PatchOps, error) {
	__traceStack()

	return nil, c.validator.Update(request, oldObj, newObj)
}

func (c *ValidatorAdapter) Delete(request *Request, oldObj runtime.Object) (PatchOps, error) {
	__traceStack()

	return nil, c.validator.Delete(request, oldObj)
}

func (c *ValidatorAdapter) Connect(request *Request, newObj runtime.Object) (PatchOps, error) {
	__traceStack()

	return nil, c.validator.Connect(request, newObj)
}

func (c *ValidatorAdapter) Resource() Resource {
	__traceStack()

	return c.validator.Resource()
}

type DefaultValidator struct {
}

func (v *DefaultValidator) Create(request *Request, newObj runtime.Object) error {
	__traceStack()

	return nil
}

func (v *DefaultValidator) Update(request *Request, oldObj runtime.Object, newObj runtime.Object) error {
	__traceStack()

	return nil
}

func (v *DefaultValidator) Delete(request *Request, oldObj runtime.Object) error {
	__traceStack()

	return nil
}

func (v *DefaultValidator) Connect(request *Request, newObj runtime.Object) error {
	__traceStack()

	return nil
}
