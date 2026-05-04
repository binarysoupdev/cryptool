package app

import (
	"errors"
	"os"

	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/binarysoupdev/go-commando/command"
	"github.com/binarysoupdev/got-style/style"
)

type DecryptCommand struct {
	command.FlagCommandBase
}

func NewDecryptCommand() *DecryptCommand {
	return &DecryptCommand{
		FlagCommandBase: command.NewFlagCommandBase("decrypt", "decrypt the given the ciphertext file"),
	}
}

func (cmd DecryptCommand) Run(args []string) error {
	in := cmd.Flags.String("i", "", "the ciphertext input file")
	out := cmd.Flags.String("o", "", "the plaintext output file")
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

	err = decrypt(bytes, *out)
	if err != nil {
		return err
	}

	if *rm {
		os.Remove(*in)
		style.Delete.Printf("[-] %s\n", *in)
	}
	return nil
}

func decrypt(in []byte, out string) error {
	password := promptPassword("Enter")

	plaintext, err := crypt.Load(password, in[:crypt.SALT_SIZE]).Decrypt(in[crypt.SALT_SIZE:])
	if err != nil {
		return chainError(err, "error decrypting ciphertext")
	}

	err = os.WriteFile(out, plaintext, 0666)
	if err != nil {
		return chainError(err, "error writing decrypted file")
	}

	style.Create.PrintF("[+] %s\n", out)
	return nil
}
