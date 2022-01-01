package version

import "fmt"

var (
	Version		= "v0.0.0-dev"
	GitCommit	= "HEAD"
)

func FriendlyVersion() string {
	__traceStack()

	return fmt.Sprintf("%s (%s)", Version, GitCommit)
}
