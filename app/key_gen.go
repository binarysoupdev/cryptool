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
		FlagCommandBase: command.NewFlagCommandBase("keygen", "generate a new encryption key"),
	}
}

func (cmd KeyGenCommand) Run(args []string) error {
	out := cmd.Flags.String("o", "", "the keyfile output")
	length := cmd.Flags.Uint("l", 32, "the key length (16, 24, or 32)")
	cmd.Flags.Parse(args)

	if *out == "" {
		return errors.New("output filepath cannot be empty")
	}
	if *length != 16 && *length != 24 && *length != 32 {
		return errors.New("invalid key length")
	}

	key := make([]byte, *length)
	rand.Read(key)

	err := os.WriteFile(*out, key, 0666)
	if err != nil {
		return chainError(err, "error writing keyfile")
	}

	style.Create.Printf("[+] %s\n", *out)
	return nil
}
