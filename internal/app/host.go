package app

import (
	"github.com/binarysoupdev/cryptool/net/host"
	"github.com/binarysoupdev/go-commando/command"
	"github.com/binarysoupdev/go-extensions/errors"
	"github.com/binarysoupdev/got-style/style"
)

type HostCommand struct {
	command.FlagCommandBase
}

func NewHostCommand() *HostCommand {
	return &HostCommand{
		FlagCommandBase: command.NewFlagCommandBase("host", "run the connect demo as host"),
	}
}

func (cmd HostCommand) Run(args []string) error {
	port := cmd.Flags.String("port", "8080", "the port to run on")
	cmd.Flags.Parse(args)

	host, err := host.Open(*port)
	if err != nil {
		return errors.Chain(err, "error opening host")
	}
	defer host.Close()

	style.BoldInfo.Printf("Listening at: %s\n", host.Addr())

	c, err := host.Accept()
	if err != nil {
		return errors.Chain(err, "error accepting connection")
	}

	style.Create.Printf("Accepted Client: %s\n", c.RemoteAddr())
	return nil
}
