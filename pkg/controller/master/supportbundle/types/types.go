package types

import "time"

const (
	StateNone	= ""
	StateGenerating	= "generating"
	StateError	= "error"
	StateReady	= "ready"

	SupportBundleLabelKey	= "rancher/supportbundle"
	DrainKey		= "kubevirt.io/drain"

	AppManager	= "support-bundle-manager"
	AppAgent	= "support-bundle-agent"

	BundleCreationTimeout	= 8 * time.Minute
	NodeBundleWaitTimeout	= "5m"
)

type ManagerStatus struct {
	Phase	string

	Error	bool

	ErrorMessage	string

	Progress	int

	Filename	string

	Filesize	int64
}
