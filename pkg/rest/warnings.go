package rest

import (
	"fmt"
	"github.com/golang/glog"
	"io"
	"sync"
)

type WarningHandler interface {
	HandleWarningHeader(code int, agent string, text string)
}

var (
	defaultWarningHandler     WarningHandler = WarningLogger{}
	defaultWarningHandlerLock sync.RWMutex
)

func SetDefaultWarningHandler(l WarningHandler) {
	defaultWarningHandlerLock.Lock()
	defer defaultWarningHandlerLock.Unlock()
	defaultWarningHandler = l
}

func getDefaultWarningHandler() WarningHandler {
	defaultWarningHandlerLock.RLock()
	defer defaultWarningHandlerLock.RUnlock()
	l := defaultWarningHandler
	return l
}

type NoWarnings struct{}

func (NoWarnings) HandleWarningHeader(code int, agent string, message string) {}

type WarningLogger struct{}

func (WarningLogger) HandleWarningHeader(code int, agent string, message string) {
	if code != 299 || len(message) == 0 {
		return
	}
	glog.Warning(message)
}

type warningWriter struct {
	out          io.Writer
	opts         WarningWriterOptions
	writtenLock  sync.Mutex
	writtenCount int
	written      map[string]struct{}
}

type WarningWriterOptions struct {
	Deduplicate bool
	Color       bool
}

func NewWarningWriter(out io.Writer, opts WarningWriterOptions) *warningWriter {
	h := &warningWriter{out: out, opts: opts}
	if opts.Deduplicate {
		h.written = map[string]struct{}{}
	}
	return h
}

const (
	yellowColor = "\u001b[33;1m"
	resetColor  = "\u001b[0m"
)

func (w *warningWriter) HandleWarningHeader(code int, agent string, message string) {
	if code != 299 || len(message) == 0 {
		return
	}

	w.writtenLock.Lock()
	defer w.writtenLock.Unlock()

	if w.opts.Deduplicate {
		if _, alreadyWritten := w.written[message]; alreadyWritten {
			return
		}
		w.written[message] = struct{}{}
	}
	w.writtenCount++

	if w.opts.Color {
		_, _ = fmt.Fprintf(w.out, "%sWarning:%s %s\n", yellowColor, resetColor, message)
	} else {
		_, _ = fmt.Fprintf(w.out, "Warning: %s\n", message)
	}
}

func (w *warningWriter) WarningCount() int {
	w.writtenLock.Lock()
	defer w.writtenLock.Unlock()
	return w.writtenCount
}
