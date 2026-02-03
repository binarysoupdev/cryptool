package app

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func promptPassword(prompt string) string {
	fmt.Printf("%s PASSWORD:\n", prompt)

	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	return string(password)
}
