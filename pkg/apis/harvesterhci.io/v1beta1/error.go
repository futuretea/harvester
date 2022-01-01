package v1beta1

type ErrorResponse struct {
	Errors []string `json:"errors,omitempty"`
}
