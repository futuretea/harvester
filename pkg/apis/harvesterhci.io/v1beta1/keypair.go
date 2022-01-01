package v1beta1

import (
	"github.com/rancher/wrangler/pkg/condition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	KeyPairValidated condition.Cond = "validated"
)

type KeyPair struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	KeyPairSpec	`json:"spec"`
	Status	KeyPairStatus	`json:"status,omitempty"`
}

type KeyPairSpec struct {
	PublicKey string `json:"publicKey"`
}

type KeyPairStatus struct {
	FingerPrint	string	`json:"fingerPrint,omitempty"`

	Conditions	[]Condition	`json:"conditions,omitempty"`
}

type KeyGenInput struct {
	Name		string	`json:"name"`
	Namespace	string	`json:"namespace"`
}
