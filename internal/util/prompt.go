package util

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

// Prompt a password from the stdin. If stdin is a terminal, echo is temporarily disabled.
func PromptPassword(prompt string) string {
	fmt.Printf("%s PASSWORD: ", prompt)

	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return readStdin()
	}

	fmt.Println()
	return string(password)
}

func readStdin() string {
	password, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return password[:len(password)-1]
}
