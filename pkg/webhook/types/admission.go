package types

import (
	"fmt"
	"strings"

	"github.com/rancher/wrangler/pkg/webhook"
	"github.com/sirupsen/logrus"
	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/harvester/harvester/pkg/webhook/config"
	werror "github.com/harvester/harvester/pkg/webhook/error"
)

const (
	AdmissionTypeMutation	= "mutation"
	AdmissionTypeValidation	= "validation"
)

type PatchOps []string

type Admitter interface {
	Create(request *Request, newObj runtime.Object) (PatchOps, error)

	Update(request *Request, oldObj runtime.Object, newObj runtime.Object) (PatchOps, error)

	Delete(request *Request, oldObj runtime.Object) (PatchOps, error)

	Connect(request *Request, newObj runtime.Object) (PatchOps, error)

	Resource() Resource
}

type AdmissionHandler struct {
	admitter	Admitter
	admissionType	string
	options		*config.Options
}

func NewAdmissionHandler(admitter Admitter, admissionType string, options *config.Options) *AdmissionHandler {
	__traceStack()

	if err := admitter.Resource().Validate(); err != nil {
		panic(err.Error())
	}
	return &AdmissionHandler{
		admitter:	admitter,
		admissionType:	admissionType,
		options:	options,
	}
}

func (v *AdmissionHandler) Admit(response *webhook.Response, request *webhook.Request) error {
	__traceStack()

	v.admit(response, NewRequest(request, v.options))
	return nil
}

func (v *AdmissionHandler) admit(response *webhook.Response, req *Request) {
	__traceStack()

	logrus.Debugf("%s admitting %s", req, v.admissionType)

	oldObj, newObj, err := req.DecodeObjects()
	if err != nil {
		logrus.Errorf("%s fail to decode objects: %s", req, err)
		response.Result = werror.NewInternalError(err.Error()).AsResult()
		response.Allowed = false
		return
	}

	var patchOps PatchOps

	switch req.Operation {
	case admissionv1.Create:
		patchOps, err = v.admitter.Create(req, newObj)
	case admissionv1.Delete:
		patchOps, err = v.admitter.Delete(req, oldObj)
	case admissionv1.Update:
		patchOps, err = v.admitter.Update(req, oldObj, newObj)
	case admissionv1.Connect:
		patchOps, err = v.admitter.Connect(req, newObj)
	default:
		err = fmt.Errorf("unsupported operation %s", req.Operation)
	}

	if err != nil {
		var admitErr werror.AdmitError
		if e, ok := err.(werror.AdmitError); ok {
			admitErr = e
		} else {
			admitErr = werror.NewInternalError(err.Error())
		}
		response.Allowed = false
		response.Result = admitErr.AsResult()
		logrus.Debugf("%s operation is rejected: %s", req, admitErr)
		return
	}

	if len(patchOps) > 0 {
		patchType := admissionv1.PatchTypeJSONPatch
		patchData := fmt.Sprintf("[%s]", strings.Join(patchOps, ","))
		response.PatchType = &patchType
		response.Patch = []byte(patchData)
		logrus.Debugf("%s patchOps: %s", req, patchData)
	}

	logrus.Debugf("%s operation is allowed", req)
	response.Allowed = true
}

func (v *AdmissionHandler) decodeObjects(request *Request) (oldObj runtime.Object, newObj runtime.Object, err error) {
	__traceStack()

	operation := request.Operation
	if operation == admissionv1.Delete || operation == admissionv1.Update {
		oldObj, err = request.DecodeOldObject()
		if err != nil {
			return
		}
		if operation == admissionv1.Delete {

			return
		}
	}
	newObj, err = request.DecodeObject()
	return
}
