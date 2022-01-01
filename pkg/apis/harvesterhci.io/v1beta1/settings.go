package v1beta1

import (
	"github.com/rancher/wrangler/pkg/condition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	SettingConfigured condition.Cond = "configured"
)

type Setting struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Value	string	`json:"value,omitempty"`

	Default	string	`json:"default,omitempty"`

	Customized	bool	`json:"customized,omitempty"`

	Source	string	`json:"source,omitempty"`

	Status	SettingStatus	`json:"status,omitempty"`
}

type SettingStatus struct {
	Conditions []Condition `json:"conditions,omitempty"`
}
