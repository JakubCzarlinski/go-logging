package logging

import (
	"fmt"
	"strings"
	"time"

	"github.com/JakubCzarlinski/go-pooling"
)

type stringBuilder struct {
	*strings.Builder
}

func (sb *stringBuilder) Reset(struct{}) {
	sb.Builder.Reset()
}

var stringBuilderPool = pooling.NewPool(func() *stringBuilder {
	builder := &strings.Builder{}
	builder.Grow(64)
	return &stringBuilder{builder}
})

func getTimestamp() string {
	if !UseTimestamp {
		return ""
	}
	return time.Now().Format("2006-01-02 15:04:05")
}

func getLinePrefix(timestamp string, i int) string {
	if !UsePrefix {
		return ""
	}
	if !UseLineLabels {
		return Magenta(fmt.Sprintf("%s ", timestamp))
	}
	return Magenta(fmt.Sprintf("%s %d. ", timestamp, i))
}

func format(msg string) string {
	if !UseLogger {
		return ""
	}
	timestamp := getTimestamp()

	sb := stringBuilderPool.Get()
	defer stringBuilderPool.Reset(sb, struct{}{})

	// Go through each character in the message.
	// If we find a newline character, we want to insert the prefix.
	// We also want to insert the prefix at the beginning of the message.
	line := 0
	sb.WriteString(getLinePrefix(timestamp, line))
	min := 0
	for i, c := range msg {
		if c == '\n' {
			sb.WriteString(msg[min:i])
			sb.WriteRune('\n')
			min = i + 1
			line++
			sb.WriteString(getLinePrefix(timestamp, line))
		}
	}
	if min < len(msg) {
		sb.WriteString(msg[min:])
	}
	sb.WriteRune('\n')

	return sb.String()
}
