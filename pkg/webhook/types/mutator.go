package types

import (
	"k8s.io/apimachinery/pkg/runtime"
)

type Mutator Admitter

type DefaultMutator struct {
}

func (v *DefaultMutator) Create(request *Request, newObj runtime.Object) (PatchOps, error) {
	__traceStack()

	return nil, nil
}

func (v *DefaultMutator) Update(request *Request, oldObj runtime.Object, newObj runtime.Object) (PatchOps, error) {
	__traceStack()

	return nil, nil
}

func (v *DefaultMutator) Delete(request *Request, oldObj runtime.Object) (PatchOps, error) {
	__traceStack()

	return nil, nil
}

func (v *DefaultMutator) Connect(request *Request, newObj runtime.Object) (PatchOps, error) {
	__traceStack()

	return nil, nil
}
