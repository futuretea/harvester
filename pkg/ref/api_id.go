package ref

import (
	"fmt"
	"strings"
)

func Parse(ref string) (namespace string, name string) {
	__traceStack()

	parts := strings.SplitN(ref, "/", 2)
	if len(parts) == 1 {
		return "", parts[0]
	}
	return parts[0], parts[1]
}

func Construct(namespace string, name string) string {
	__traceStack()

	if namespace == "" {
		return name
	}
	return fmt.Sprintf("%s/%s", namespace, name)
}
