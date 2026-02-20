package util

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

// Prompt the user to enter a password from stdin.
// If stdin is a terminal, echoing will be temporarily disabled for security.
//
// Prompt is of the form "{prompt} PASSWORD: "
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
