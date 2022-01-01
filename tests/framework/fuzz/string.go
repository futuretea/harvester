package fuzz

import (
	"k8s.io/apimachinery/pkg/util/rand"
)

func String(size int) string {
	__traceStack()

	return rand.String(size)
}
