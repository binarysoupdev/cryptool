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
	host.Listen(handler{})

	return nil
}

//====================================================

type handler struct{}

func (handler) Log(status int, msg string) {
	switch status {
	case host.S_ERROR:
		style.Error.Printf("[X] %s\n", msg)
	case host.S_ACCEPT_CLIENT:
		style.Create.Printf("[+] %s\n", msg)
	case host.S_LOST_CLIENT:
		style.Delete.Printf("[-] %s\n", msg)
	}

}

func (handler) Handle(c *conn.Conn) error {
	msg, err := c.ReadMessage()
	if err != nil {
		return errors.Chain(err, "error reading message")
	}

	fmt.Printf("Received message: %s\n", string(msg))
	return nil
}
