package logging

import (
	"errors"
	"fmt"
)

func Bubble(err error, msg string) error {
	if err == nil || !UseLogger {
		return errors.New(msg)
	}
	return fmt.Errorf("%w\n%s", err, msg)
}

func BubbleF(err error, format string, args ...interface{}) error {
	if err == nil || !UseLogger {
		return fmt.Errorf(format, args...)
	}
	return fmt.Errorf("%s\n%w", fmt.Sprintf(format, args...), err)
}
