package util

import (
	"encoding/json"
	"net/http"

	"github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
)

func ResponseBody(obj interface{}) []byte {
	__traceStack()

	respBody, err := json.Marshal(obj)
	if err != nil {
		return []byte(`{\"errors\":[\"Failed to parse response body\"]}`)
	}
	return respBody
}

func ResponseOKWithBody(rw http.ResponseWriter, obj interface{}) {
	__traceStack()

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(ResponseBody(obj))
}

func ResponseOK(rw http.ResponseWriter) {
	__traceStack()

	rw.WriteHeader(http.StatusOK)
}

func ResponseError(rw http.ResponseWriter, statusCode int, err error) {
	__traceStack()

	ResponseErrorMsg(rw, statusCode, err.Error())
}

func ResponseErrorMsg(rw http.ResponseWriter, statusCode int, errMsg string) {
	__traceStack()

	rw.WriteHeader(statusCode)
	_, _ = rw.Write(ResponseBody(v1beta1.ErrorResponse{Errors: []string{errMsg}}))
}
