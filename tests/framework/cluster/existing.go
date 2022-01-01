package cluster

import (
	"io"

	"github.com/harvester/harvester/tests/framework/logs"
)

var _ Cluster = &ExistingCluster{}

type ExistingCluster struct {
}

func (e *ExistingCluster) String() string {
	__traceStack()

	return "this is an existing cluster"
}

func (e *ExistingCluster) GetKind() string {
	__traceStack()

	return ExistingClusterKind
}

func (e *ExistingCluster) Startup(output io.Writer) error {
	__traceStack()

	var logger = logs.NewLogger(output, 0)
	logger.V(0).Info("skip exist cluster")
	return nil
}

func (e *ExistingCluster) LoadImages(output io.Writer) error {
	__traceStack()

	var logger = logs.NewLogger(output, 0)
	logger.V(0).Info("skip loading images")
	return nil
}

func (e *ExistingCluster) Cleanup(output io.Writer) error {
	__traceStack()

	var logger = logs.NewLogger(output, 0)
	logger.V(0).Info("skip exist cluster")
	return nil
}
