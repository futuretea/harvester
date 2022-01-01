package setting

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/util"
)

func (h *Handler) syncAdditionalTrustedCAs(setting *harvesterv1.Setting) error {
	__traceStack()

	backupConfig := map[string]string{
		"AWS_CERT": setting.Value,
	}
	if err := h.updateBackupSecret(backupConfig); err != nil {
		return err
	}

	if err := h.syncAdditionalCASecrets(setting); err != nil {
		return err
	}

	return h.redeployDeployment(h.namespace, "harvester")
}

func (h *Handler) syncAdditionalCASecrets(setting *harvesterv1.Setting) error {
	__traceStack()

	namespaces := []string{h.namespace, util.LonghornSystemNamespaceName, util.CattleSystemNamespaceName}

	for _, namespace := range namespaces {
		secret, err := h.secretCache.Get(namespace, util.AdditionalCASecretName)
		if errors.IsNotFound(err) {
			newSecret := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:		util.AdditionalCASecretName,
					Namespace:	namespace,
				},
				Data: map[string][]byte{
					util.AdditionalCAFileName: []byte(setting.Value),
				},
			}

			if _, err := h.secrets.Create(newSecret); err != nil {
				return err
			}
			continue
		} else if err != nil {
			return err
		}

		toUpdate := secret.DeepCopy()
		toUpdate.Data = map[string][]byte{
			util.AdditionalCAFileName: []byte(setting.Value),
		}
		if _, err := h.secrets.Update(toUpdate); err != nil {
			return err
		}
	}
	return nil
}
