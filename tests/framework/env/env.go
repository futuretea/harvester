package env

import (
	"os"
	"strings"
)

const (
	envUseExistingCluster	= "USE_EXISTING_CLUSTER"

	envKeepTestingCluster	= "KEEP_TESTING_CLUSTER"

	envSkipHarvesterInstallation	= "SKIP_HARVESTER_INSTALLATION"

	envKeepHarvesterInstallation	= "KEEP_HARVESTER_INSTALLATION"

	envKeepTestingResource	= "KEEP_TESTING_RESOURCE"

	envDontUseEmulation	= "DONT_USE_EMULATION"

	envEnableE2ETests	= "ENABLE_E2E_TESTS"

	envPreloadingImages	= "PRELOADING_IMAGES"

	envWebhookImage	= "WEBHOOK_IMAGE_NAME"
)

func IsTrue(key string) bool {
	__traceStack()

	return strings.EqualFold(os.Getenv(key), "true")
}

func IsUsingExistingCluster() bool {
	__traceStack()

	return IsTrue(envUseExistingCluster)
}

func IsKeepingTestingCluster() bool {
	__traceStack()

	return IsTrue(envKeepTestingCluster)
}

func IsSkipHarvesterInstallation() bool {
	__traceStack()

	return IsTrue(envSkipHarvesterInstallation)
}

func IsKeepingHarvesterInstallation() bool {
	__traceStack()

	return IsTrue(envKeepHarvesterInstallation)
}

func IsKeepingTestingResource() bool {
	__traceStack()

	return IsTrue(envKeepTestingResource)
}

func IsUsingEmulation() bool {
	__traceStack()

	return !IsTrue(envDontUseEmulation)
}

func IsE2ETestsEnabled() bool {
	__traceStack()

	return IsTrue(envEnableE2ETests)
}

func GetPreloadingImages() []string {
	__traceStack()

	images := []string{}
	for _, image := range strings.Split(os.Getenv(envPreloadingImages), ",") {
		images = append(images, strings.TrimSpace(image))
	}
	return images
}

func GetWebhookImage() (string, string) {
	__traceStack()

	image := os.Getenv(envWebhookImage)
	if image == "" {
		return "", ""
	}

	tokens := strings.Split(image, ":")
	if len(tokens) > 1 {
		return tokens[0], tokens[1]
	}
	return tokens[0], ""
}
