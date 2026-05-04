package app

import (
	"crypto/rand"
	"errors"
	"os"

	"github.com/binarysoupdev/go-commando/command"
	"github.com/binarysoupdev/got-style/style"
)

type KeyGenCommand struct {
	command.FlagCommandBase
}

func NewKeyGenCommand() *KeyGenCommand {
	return &KeyGenCommand{
		FlagCommandBase: command.NewFlagCommandBase("key", "generate a new encryption key"),
	}
}

func (cmd KeyGenCommand) Run(args []string) error {
	out := cmd.Flags.String("o", "", "the keyfile output")
	cmd.Flags.Parse(args)

	if *out == "" {
		return errors.New("output filepath cannot be empty")
	}

	//TODO: add length parameter (16, 24, or 32)

	key := make([]byte, 32) //256-bit
	rand.Read(key)

	err := os.WriteFile(*out, key, 0666)
	if err != nil {
		return chainError(err, "error writing keyfile")
	}

	style.Create.Printf("[+] %s\n", *out)
	return nil
}
