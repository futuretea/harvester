package printer

import (
	"fmt"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

var _ ginkgo.Reporter = NewlineReporter{}

type NewlineReporter struct{}

func (NewlineReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
	__traceStack()

}

func (NewlineReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary)	{ __traceStack() }

func (NewlineReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary)	{ __traceStack() }

func (NewlineReporter) SpecWillRun(specSummary *types.SpecSummary)	{ __traceStack() }

func (NewlineReporter) SpecDidComplete(specSummary *types.SpecSummary)	{ __traceStack() }

func (NewlineReporter) SpecSuiteDidEnd(summary *types.SuiteSummary)	{ __traceStack(); fmt.Printf("\n") }
