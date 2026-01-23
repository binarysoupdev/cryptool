package util

import "fmt"

func ChainError(err error, msg string) error {
	return fmt.Errorf("%s\n  %s", msg, err.Error())
}
