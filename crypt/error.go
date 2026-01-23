package crypt

import "fmt"

func chainError(err error, msg string) error {
	return fmt.Errorf("%s\n  %s", msg, err.Error())
}
