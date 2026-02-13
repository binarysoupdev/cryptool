package util

import "fmt"

// Append a new message string to an existing error and return the resulting error object.
func ChainError(err error, msg string) error {
	return fmt.Errorf("%s\n  %s", msg, err.Error())
}
