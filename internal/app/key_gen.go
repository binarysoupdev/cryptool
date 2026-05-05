package app

import (
	"crypto/rand"
	"errors"
	"os"

	"github.com/binarysoupdev/cryptool/internal/util"
	"github.com/binarysoupdev/go-commando/command"
	"github.com/binarysoupdev/got-style/style"
)

// A command for generating a new keyfile.
// Supports 128, 192, and 256-bit keys.
type KeyGenCommand struct {
	command.FlagCommandBase
}

// Create a new KeyGen command.
func NewKeyGenCommand() *KeyGenCommand {
	return &KeyGenCommand{
		FlagCommandBase: command.NewFlagCommandBase("keygen", "generate a new encryption key"),
	}
}

// Run the command. See usage for details.
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
		return util.ChainError(err, "error writing keyfile")
	}

	style.Create.Printf("[+] %s\n", *out)
	return nil
}
