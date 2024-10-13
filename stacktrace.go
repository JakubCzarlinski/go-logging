package logging

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/JakubCzarlinski/go-pooling"
)

type stackTrace struct {
	callers []uintptr
}

func (s *stackTrace) Reset(struct{}) {
	s.callers = s.callers[:0]
}

var stackTracePool = pooling.NewPool(func() *stackTrace {
	return &stackTrace{callers: make([]uintptr, 50)}
})

func createStackMessage(msg string, colour ColourFunc) string {
	// Get the stack in the stack up.
	stack := stackTracePool.Get()
	defer stackTracePool.Reset(stack, struct{}{})
	numCallers := runtime.Callers(3, stack.callers)
	stack.callers = stack.callers[:numCallers]
	frames := runtime.CallersFrames(stack.callers)

	timestamp := getTimestamp()

	lines := make([]string, numCallers)
	for i := numCallers - 1; i >= 0; i-- {
		frame, _ := frames.Next()
		prefix := getLinePrefix(timestamp, i+1)
		lines[i] = fmt.Sprintf(
			"%s%s:%d %s", prefix, frame.File, frame.Line, frame.Function,
		)
	}
	output := strings.Join(lines, "\n")

	msgLines := strings.Split(msg, "\n")
	numLines := len(msgLines)

	msgLines[numLines-1] = colour(msgLines[numLines-1])
	for i := 0; i < numLines; i++ {
		msgLines[i] = getLinePrefix(timestamp, numCallers+numLines-i) + msgLines[i]
	}
	msg = strings.Join(msgLines, "\n") + "\n"

	if output != "" {
		output += "\n"
	}

	return output + msg
}
