package util

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func PromptPassword(prompt string) string {
	fmt.Printf("%s PASSWORD: ", prompt)

	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	fmt.Println()
	return string(password)
}
