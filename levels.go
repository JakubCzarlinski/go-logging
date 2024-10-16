package logging

import (
	"fmt"
)

const (
	DEBUG int = iota
	INFO  int = iota
	WARN  int = iota
	ERROR int = iota
	FATAL int = iota
)

type ColourFunc func(string) string

var colours [](ColourFunc) = [](ColourFunc){
	White,                         // DEBUG
	Cyan,                          // INFO
	Yellow,                        // WARN
	Red,                           // ERROR
	Combine(Red, Bold, Underline), // FATAL
}

func shouldLog(level int) bool {
	return UseLogger && level >= MinLogLevel
}

func Debug(msg string) {
	withMessage(msg, DEBUG)
}

func DebugF(format string, args ...interface{}) {
	withMessageF(DEBUG, format, args...)
}

func Info(msg string) {
	withMessage(msg, INFO)
}

func InfoF(format string, args ...interface{}) {
	withMessageF(INFO, format, args...)
}

func Warn(msg string) {
	withMessage(msg, WARN)
}

func WarnF(format string, args ...interface{}) {
	withMessageF(WARN, format, args...)
}

func Error(msg string) error {
	withMessage(msg, ERROR)
	return Bubble(nil, msg)
}

func ErrorF(format string, args ...interface{}) error {
	return Bubble(nil, withMessageF(ERROR, format, args...))
}

func ErrorBubble(err error, msg string) error {
	return withBubble(ERROR, err, msg)
}

func ErrorBubbleF(err error, format string, args ...interface{}) error {
	return withBubbleF(ERROR, err, format, args...)
}

func ErrorStack(err error) error {
	fmt.Print(createStackMessage(err.Error(), colours[ERROR]))
	return err
}

func Fatal(msg string) {
	if !shouldLog(FATAL) {
		panic(msg)
	}
	panic(format(colours[FATAL](msg)))
}

func FatalF(f string, args ...interface{}) {
	if !shouldLog(FATAL) {
		panic(fmt.Sprintf(f, args...))
	}
	panic(format(colours[FATAL](fmt.Sprintf(f, args...))))
}

func FatalBubble(err error, msg string) {
	if !shouldLog(FATAL) {
		panic(err)
	}
	panic(withBubble(FATAL, err, msg))
}

func FatalBubbleF(err error, f string, args ...interface{}) {
	if !shouldLog(FATAL) {
		panic(err)
	}
	panic(withBubbleF(FATAL, err, f, args...))
}

func withMessage(msg string, level int) {
	if !shouldLog(level) {
		return
	}
	fmt.Print(format(colours[level](msg)))
}

func withMessageF(level int, f string, args ...interface{}) string {
	if !shouldLog(level) {
		return ""
	}
	str := format(colours[level](fmt.Sprintf(f, args...)))
	fmt.Print(str)
	return str
}

func withBubble(level int, err error, msg string) error {
	if !shouldLog(level) {
		return err
	}
	bubble := Bubble(err, colours[level](msg))
	fmt.Print(format(bubble.Error()))
	return bubble
}

func withBubbleF(level int, err error, f string, args ...interface{}) error {
	if !shouldLog(level) {
		return err
	}
	bubble := Bubble(err, colours[level](fmt.Sprintf(f, args...)))
	fmt.Print(format(bubble.Error()))
	return bubble
}
