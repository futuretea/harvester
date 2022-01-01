package logs

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"sigs.k8s.io/kind/pkg/log"
)

type Logger struct {
	writer		io.Writer
	bufferPool	*bufferPool
	writerMu	sync.Mutex
	verbosity	log.Level

	isSmartWriter	bool
}

var _ log.Logger = &Logger{}

func NewLogger(writer io.Writer, verbosity log.Level) *Logger {
	__traceStack()

	l := &Logger{
		verbosity:	verbosity,
		bufferPool:	newBufferPool(),
	}
	l.SetWriter(writer)
	return l
}

func (l *Logger) SetWriter(w io.Writer) {
	__traceStack()

	l.writerMu.Lock()
	defer l.writerMu.Unlock()
	l.writer = w
	l.isSmartWriter = IsSmartTerminal(w)
}

func (l *Logger) ColorEnabled() bool {
	__traceStack()

	l.writerMu.Lock()
	defer l.writerMu.Unlock()
	return l.isSmartWriter
}

func (l *Logger) getVerbosity() log.Level {
	__traceStack()

	return log.Level(atomic.LoadInt32((*int32)(&l.verbosity)))
}

func (l *Logger) SetVerbosity(verbosity log.Level) {
	__traceStack()

	atomic.StoreInt32((*int32)(&l.verbosity), int32(verbosity))
}

func (l *Logger) write(p []byte) (n int, err error) {
	__traceStack()

	l.writerMu.Lock()
	defer l.writerMu.Unlock()
	return l.writer.Write(p)
}

func (l *Logger) writeBuffer(buf *bytes.Buffer) {
	__traceStack()

	if buf.Len() == 0 || buf.Bytes()[buf.Len()-1] != '\n' {
		buf.WriteByte('\n')
	}

	_, _ = l.write(buf.Bytes())
}

func (l *Logger) print(message string) {
	__traceStack()

	buf := bytes.NewBufferString(message)
	l.writeBuffer(buf)
}

func (l *Logger) printf(format string, args ...interface{}) {
	__traceStack()

	buf := l.bufferPool.Get()
	fmt.Fprintf(buf, format, args...)
	l.writeBuffer(buf)
	l.bufferPool.Put(buf)
}

func addDebugHeader(buf *bytes.Buffer) {
	__traceStack()

	_, file, line, ok := runtime.Caller(3)

	if !ok {
		file = "???"
		line = 1
	} else {
		if slash := strings.LastIndex(file, "/"); slash >= 0 {
			path := file
			file = path[slash+1:]
			if dirsep := strings.LastIndex(path[:slash], "/"); dirsep >= 0 {
				file = path[dirsep+1:]
			}
		}
	}
	buf.Grow(len(file) + 11)
	buf.WriteString("DEBUG: ")
	buf.WriteString(file)
	buf.WriteByte(':')
	_, _ = fmt.Fprintf(buf, "%d", line)
	buf.WriteByte(']')
	buf.WriteByte(' ')
}

func (l *Logger) debug(message string) {
	__traceStack()

	buf := l.bufferPool.Get()
	addDebugHeader(buf)
	buf.WriteString(message)
	l.writeBuffer(buf)
	l.bufferPool.Put(buf)
}

func (l *Logger) debugf(format string, args ...interface{}) {
	__traceStack()

	buf := l.bufferPool.Get()
	addDebugHeader(buf)
	_, _ = fmt.Fprintf(buf, format, args...)
	l.writeBuffer(buf)
	l.bufferPool.Put(buf)
}

func (l *Logger) Warn(message string) {
	__traceStack()

	l.print(message)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	__traceStack()

	l.printf(format, args...)
}

func (l *Logger) Error(message string) {
	__traceStack()

	l.print(message)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	__traceStack()

	l.printf(format, args...)
}

func (l *Logger) V(level log.Level) log.InfoLogger {
	__traceStack()

	return infoLogger{
		logger:		l,
		level:		level,
		enabled:	level <= l.getVerbosity(),
	}
}

type infoLogger struct {
	logger	*Logger
	level	log.Level
	enabled	bool
}

func (i infoLogger) Enabled() bool {
	__traceStack()

	return i.enabled
}

func (i infoLogger) Info(message string) {
	__traceStack()

	if !i.enabled {
		return
	}

	if i.level > 0 {
		i.logger.debug(message)
	} else {
		i.logger.print(message)
	}
}

func (i infoLogger) Infof(format string, args ...interface{}) {
	__traceStack()

	if !i.enabled {
		return
	}

	if i.level > 0 {
		i.logger.debugf(format, args...)
	} else {
		i.logger.printf(format, args...)
	}
}

type bufferPool struct {
	sync.Pool
}

func newBufferPool() *bufferPool {
	__traceStack()

	return &bufferPool{
		sync.Pool{
			New: func() interface{} {

				return new(bytes.Buffer)
			},
		},
	}
}

func (b *bufferPool) Get() *bytes.Buffer {
	__traceStack()

	return b.Pool.Get().(*bytes.Buffer)
}

func (b *bufferPool) Put(x *bytes.Buffer) {
	__traceStack()

	if x.Len() > 256 {
		return
	}
	x.Reset()
	b.Pool.Put(x)
}
