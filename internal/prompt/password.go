package prompt

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

func Password(prompt string) string {
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
