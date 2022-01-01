package server

import (
	"fmt"

	"github.com/rancher/wrangler/pkg/kubeconfig"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetConfig(kubeConfig string) (clientcmd.ClientConfig, error) {
	__traceStack()

	if isManual(kubeConfig) {
		return kubeconfig.GetNonInteractiveClientConfig(kubeConfig), nil
	}

	return getEmbedded()
}

func isManual(kubeConfig string) bool {
	__traceStack()

	if kubeConfig != "" {
		return true
	}
	_, inClusterErr := rest.InClusterConfig()
	return inClusterErr == nil
}

func getEmbedded() (clientcmd.ClientConfig, error) {
	__traceStack()

	return nil, fmt.Errorf("embedded only supported on linux")
}
