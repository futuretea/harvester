package runtime

import (
	"context"
	"fmt"

	restclient "k8s.io/client-go/rest"

	"github.com/harvester/harvester/tests/framework/env"
	"github.com/harvester/harvester/tests/framework/helm"
	"github.com/harvester/harvester/tests/framework/ready"
)

func Destruct(ctx context.Context, kubeConfig *restclient.Config) error {
	__traceStack()

	if env.IsKeepingHarvesterInstallation() || env.IsSkipHarvesterInstallation() {
		return nil
	}

	err := uninstallHarvesterCharts(ctx, kubeConfig)
	if err != nil {
		return err
	}

	return nil
}

func uninstallHarvesterCharts(ctx context.Context, kubeConfig *restclient.Config) error {
	__traceStack()

	_, err := helm.UninstallChart(testChartReleaseName, testHarvesterNamespace)
	if err != nil {
		return fmt.Errorf("failed to uninstall harvester chart: %v", err)
	}

	namespaceReadyCondition, err := ready.NewNamespaceCondition(kubeConfig, testHarvesterNamespace)
	if err != nil {
		return fmt.Errorf("faield to create namespace ready condition from kubernetes config: %w", err)
	}
	namespaceReadyCondition.AddDeploymentsClean(testDeploymentManifest...)
	namespaceReadyCondition.AddDaemonSetsClean(testDaemonSetManifest...)

	return namespaceReadyCondition.Wait(ctx)
}
