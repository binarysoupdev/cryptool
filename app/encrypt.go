package app

import (
	"errors"
	"os"

	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/binarysoupdev/go-commando/command"
	"github.com/binarysoupdev/got-style/style"
)

type EncryptCommand struct {
	command.FlagCommandBase
}

func NewEncryptCommand() *EncryptCommand {
	return &EncryptCommand{
		FlagCommandBase: command.NewFlagCommandBase("encrypt", "encrypt the given the plaintext file"),
	}
}

func (cmd EncryptCommand) Run(args []string) error {
	in := cmd.Flags.String("i", "", "the plaintext input file")
	out := cmd.Flags.String("o", "", "the ciphertext output")
	rm := cmd.Flags.Bool("rm", false, "remove the plaintext file")
	cmd.Flags.Parse(args)

	if *in == "" {
		return errors.New("input filepath cannot be empty")
	}
	if *out == "" {
		return errors.New("output filepath cannot be empty")
	}

	bytes, err := os.ReadFile(*in)
	if err != nil {
		return chainError(err, "error reading plaintext file")
	}

	err = encrypt(bytes, *out)
	if err != nil {
		return err
	}

	if *rm {
		os.Remove(*in)
		style.Delete.Printf("[-] %s\n", *in)
	}
	return nil
}

func encrypt(in []byte, out string) error {
	password := promptPassword("New")
	verify := promptPassword("Verify")

	if verify != password {
		return errors.New("passwords do not match")
	}

	c, salt := crypt.New(password)
	ciphertext := c.Encrypt(in)

	err := os.WriteFile(out, append(salt, ciphertext...), 0666)
	if err != nil {
		return chainError(err, "error writing encrypted file")
	}

	style.Create.Printf("[+] %s\n", out)
	return nil
}
