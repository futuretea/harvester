package logs

import (
	"io"
	"os"
	"runtime"

	"github.com/mattn/go-isatty"
)

func IsTerminal(w io.Writer) bool {
	__traceStack()

	if v, ok := (w).(*os.File); ok {
		return isatty.IsTerminal(v.Fd())
	}
	return false
}

func IsSmartTerminal(w io.Writer) bool {
	__traceStack()

	return isSmartTerminal(w, runtime.GOOS, os.LookupEnv)
}

func isSmartTerminal(w io.Writer, GOOS string, lookupEnv func(string) (string, bool)) bool {
	__traceStack()

	if !IsTerminal(w) {
		return false
	}

	getenv := func(e string) string {
		v, _ := lookupEnv(e)
		return v
	}

	if _, set := lookupEnv("NO_COLOR"); set {
		return false
	}

	if getenv("TERM") == "dumb" {
		return false
	}

	if GOOS == "windows" && getenv("WT_SESSION") == "" {
		return false
	}

	if getenv("HAS_JOSH_K_SEAL_OF_APPROVAL") == "true" && getenv("TRAVIS") == "true" {
		return false
	}

	return true
}
