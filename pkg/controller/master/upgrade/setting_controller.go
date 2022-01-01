package upgrade

import (
	"github.com/sirupsen/logrus"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
)

type settingHandler struct {
	versionSyncer *versionSyncer
}

func (h *settingHandler) OnChanged(key string, setting *harvesterv1.Setting) (*harvesterv1.Setting, error) {
	__traceStack()

	if setting == nil || setting.DeletionTimestamp != nil || setting.Name != "server-version" {
		return setting, nil
	}
	if err := h.versionSyncer.sync(); err != nil {
		logrus.Errorf("failed syncing version metadata: %v", err)
	}
	return setting, nil
}
