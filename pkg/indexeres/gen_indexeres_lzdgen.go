package indexeres

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var gopaths []string

func init() {
	gopath := os.Getenv("GOPATH")
	gopaths = strings.Split(gopath, ";")
	for i := range gopaths {
		gopaths[i] += "/src/"
	}
}

func __traceStack() {
	caller, file, line := __caller()
	fmt.Printf("[lazydog][%s] %s:%d caller= %s  \n", __curGid(), __prettyFile(file), line, __prettyCaller(caller))
}

func __caller() (string, string, int) {

	fpcs := make([]uintptr, 1)

	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return "", "n/a", 0
	}

	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return "", "n/a", 0
	}

	file, line := fun.FileLine(0)

	return fun.Name(), file, line
}

func __prettyCaller(caller string) string {
	return string(string(caller[strings.LastIndex(caller, ".")+1:]))
}

func __prettyFile(file string) string {
	for _, gopath := range gopaths {
		if strings.Contains(file, gopath) {
			return strings.Replace(file, gopath, "", -1)
		}
	}
	return file
}

func __curGid() string {
	var buf [64]byte
	runtime.Stack(buf[:], false)
	return strings.Fields(strings.TrimPrefix(string(buf[:]), "goroutine "))[0]
}
