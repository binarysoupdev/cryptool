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
	key := cmd.Flags.String("key", "", "encrypt using a keyfile (16, 24, or 32 bytes)")
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

	if *key == "" {
		err = cmd.encryptFromPassword(bytes, *out)
	} else {
		err = cmd.encryptFromKeyfile(*key, bytes, *out)
	}
	if err != nil {
		return err
	}

	if *rm {
		os.Remove(*in)
		style.Delete.Printf("[-] %s\n", *in)
	}
	return nil
}

func (cmd EncryptCommand) encryptFromPassword(in []byte, out string) error {
	password := promptPassword("New")
	verify := promptPassword("Verify")

	if verify != password {
		return errors.New("passwords do not match")
	}

	c, salt := crypt.NewFromPassword(password)
	ciphertext := c.Encrypt(in)

	return cmd.writeCiphertextFile(append(salt, ciphertext...), out)
}

func (cmd EncryptCommand) encryptFromKeyfile(key string, in []byte, out string) error {
	bytes, err := os.ReadFile(key)
	if err != nil {
		return chainError(err, "error reading keyfile")
	}

	c, err := crypt.New(bytes)
	if err != nil {
		return err
	}

	return cmd.writeCiphertextFile(c.Encrypt(in), out)
}

func (EncryptCommand) writeCiphertextFile(bytes []byte, out string) error {
	err := os.WriteFile(out, bytes, 0666)
	if err != nil {
		return chainError(err, "error writing encrypted file")
	}

	style.Create.Printf("[+] %s\n", out)
	return nil
}
