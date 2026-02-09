package util

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

func PromptPassword(prompt string) string {
	fmt.Printf("%s PASSWORD: ", prompt)
	fd := int(os.Stdin.Fd())

	if term.IsTerminal(fd) {
		return readTerminal(fd)
	} else {
		return readStdin()
	}
}

func readTerminal(fd int) string {
	password, err := term.ReadPassword(fd)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	return string(password)
}

func readStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}
