package locallogging

import (
	"fmt"
	"runtime"
	"strings"
)

const maxStackLength = 50

// Error is the type that implements the error interface.
// It contains the underlying err and its stacktrace.
type ErrorStackTrace struct {
	Err        error
	StackTrace string
}

func (m ErrorStackTrace) Error() string {
	return m.Err.Error() + m.StackTrace
}

// Wrap annotates the given error with a stack trace
func Wrap(err error) ErrorStackTrace {
	return ErrorStackTrace{Err: err, StackTrace: getStackTrace()}
}

func getStackTrace() string {
	stackBuf := make([]uintptr, maxStackLength)
	length := runtime.Callers(3, stackBuf[:])
	stack := stackBuf[:length]

	trace := ""
	frames := runtime.CallersFrames(stack)
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			trace = trace + fmt.Sprintf("\n\tFile: %s, Line: %d. Function: %s", frame.File, frame.Line, frame.Function)
		}
		if !more {
			break
		}
	}
	return trace
}
