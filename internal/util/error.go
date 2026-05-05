package util

import "fmt"

// Chain an error with a new message string.
func ChainError(err error, msg string) error {
	return fmt.Errorf("%s\n  %s", msg, err.Error())
}
