package v1beta1

import (
	"github.com/rancher/wrangler/pkg/condition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	UpgradeCompleted	condition.Cond	= "Completed"

	ImageReady	condition.Cond	= "ImageReady"

	RepoProvisioned	condition.Cond	= "RepoReady"

	NodesPrepared	condition.Cond	= "NodesPrepared"

	NodesUpgraded	condition.Cond	= "NodesUpgraded"

	SystemServicesUpgraded	condition.Cond	= "SystemServicesUpgraded"
)

type Upgrade struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	UpgradeSpec	`json:"spec"`
	Status	UpgradeStatus	`json:"status,omitempty"`
}

type UpgradeSpec struct {
	Version	string	`json:"version"`

	Image	string	`json:"image"`
}

type UpgradeStatus struct {
	PreviousVersion	string	`json:"previousVersion,omitempty"`

	ImageID	string	`json:"imageID,omitempty"`

	RepoInfo	string	`json:"repoInfo,omitempty"`

	SingleNode	string	`json:"singleNode,omitempty"`

	NodeStatuses	map[string]NodeUpgradeStatus	`json:"nodeStatuses,omitempty"`

	Conditions	[]Condition	`json:"conditions,omitempty"`
}

type NodeUpgradeStatus struct {
	State	string	`json:"state,omitempty"`
	Reason	string	`json:"reason,omitempty"`
	Message	string	`json:"message,omitempty"`
}
