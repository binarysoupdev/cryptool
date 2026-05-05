package app

import (
	"errors"
	"os"

	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/binarysoupdev/go-commando/command"
	"github.com/binarysoupdev/got-style/style"
)

// A command for decrypting ciphertext files.
// Supports both password and file-based keys.
type DecryptCommand struct {
	command.FlagCommandBase
}

// Create a new Decrypt command.
func NewDecryptCommand() *DecryptCommand {
	return &DecryptCommand{
		FlagCommandBase: command.NewFlagCommandBase("decrypt", "decrypt the given the ciphertext file"),
	}
}

// Run the command. See usage for details.
func (cmd DecryptCommand) Run(args []string) error {
	in := cmd.Flags.String("i", "", "the ciphertext input file")
	out := cmd.Flags.String("o", "", "the plaintext output file")
	key := cmd.Flags.String("key", "", "decrypt using a keyfile (16, 24, or 32 bytes)")
	rm := cmd.Flags.Bool("rm", false, "remove the ciphertext file")
	cmd.Flags.Parse(args)

	if *in == "" {
		return errors.New("input filepath cannot be empty")
	}
	if *out == "" {
		return errors.New("output filepath cannot be empty")
	}

	bytes, err := os.ReadFile(*in)
	if err != nil {
		return chainError(err, "error reading ciphertext file")
	}

	if *key == "" {
		err = cmd.decryptFromPassword(bytes, *out)
	} else {
		err = cmd.decryptFromKeyfile(*key, bytes, *out)
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

func (cmd DecryptCommand) decryptFromPassword(in []byte, out string) error {
	password := promptPassword("Enter")

	plaintext, err := crypt.LoadFromPassword(password, in[:crypt.SALT_SIZE]).Decrypt(in[crypt.SALT_SIZE:])
	if err != nil {
		return chainError(err, "error decrypting ciphertext")
	}

	return cmd.writePlaintextFile(plaintext, out)
}

func (cmd DecryptCommand) decryptFromKeyfile(key string, in []byte, out string) error {
	bytes, err := os.ReadFile(key)
	if err != nil {
		return chainError(err, "error reading keyfile")
	}

	c, err := crypt.New(bytes)
	if err != nil {
		return err
	}

	plaintext, err := c.Decrypt(in)
	if err != nil {
		return chainError(err, "error decrypting ciphertext")
	}

	return cmd.writePlaintextFile(plaintext, out)
}

func (DecryptCommand) writePlaintextFile(bytes []byte, out string) error {
	err := os.WriteFile(out, bytes, 0666)
	if err != nil {
		return chainError(err, "error writing decrypted file")
	}

	style.Create.Printf("[+] %s\n", out)
	return nil
}
