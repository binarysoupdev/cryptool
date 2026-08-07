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

	style.BoldInfo.Printf("Listening at: %s\n", host.Addr())
	host.Listen(handler{}, 1)

	return nil
}

//====================================================

type handler struct{}

func (handler) Log(status int, clientID int, v any) {
	switch status {
	case host.S_ERROR:
		style.Error.Printf("[ID:%d] [X] %v\n", clientID, v)
	case host.S_ACCEPT_CLIENT:
		style.Create.Printf("[ID:%d] [+] %v\n", clientID, v)
	case host.S_LOST_CLIENT:
		style.Delete.Printf("[ID:%d] [-] %v\n", clientID, v)
	}
}

func (handler) Handle(_ int, c *conn.Conn) error {
	msg, err := c.ReadMessage()
	if err != nil {
		return errors.Chain(err, "error reading message")
	}

	fmt.Printf("Received message: %s\n", string(msg))
	return nil
}
