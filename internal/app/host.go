package app

import (
	"fmt"

	"github.com/binarysoupdev/cryptool/net/conn"
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

	cmd.acceptLoop(host)
	return nil
}

func (cmd HostCommand) acceptLoop(h host.Host) {
	style.BoldInfo.Printf("Listening at: %s\n", h.Addr())
	for {
		err := cmd.acceptClient(h)
		if err != nil {
			cmd.logError(err)
		}
	}
}

func (cmd HostCommand) acceptClient(h host.Host) error {
	c, err := h.Accept()
	if err != nil {
		return errors.Chain(err, "error accepting connection")
	}

	style.Create.Printf("Accepted Client: %s\n", c.RemoteAddr())
	err = cmd.accept(c)
	style.Delete.Printf("Ended Client: %s\n", c.RemoteAddr())

	return err
}

func (cmd HostCommand) accept(c *conn.Conn) error {
	msg, err := c.ReadMessage()
	if err != nil {
		return errors.Chain(err, "error reading message")
	}

	fmt.Printf("Received message: %s\n", string(msg))
	return nil
}

func (cmd HostCommand) logError(err error) {
	fmt.Printf("%s %s\n", style.BoldError.Sprint("[X] "), err)
}
