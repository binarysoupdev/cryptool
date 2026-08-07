package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/binarysoupdev/cryptool/net/client"
	"github.com/binarysoupdev/go-commando/command"
	"github.com/binarysoupdev/go-extensions/errors"
	"github.com/binarysoupdev/got-style/style"
)

type ClientCommand struct {
	command.FlagCommandBase
}

func NewClientCommand() *ClientCommand {
	return &ClientCommand{
		FlagCommandBase: command.NewFlagCommandBase("client", "run the connect demo as client"),
	}
}

func (cmd ClientCommand) Run(args []string) error {
	addr := cmd.Flags.String("addr", "", "the address to connect to")
	cmd.Flags.Parse(args)

	if *addr == "" {
		return errors.New("\"addr\" cannot be empty")
	}

	client, err := client.Connect(*addr)
	if err != nil {
		return errors.Chain(err, "error connecting client")
	}
	defer client.Close()
	style.Create.Printf("Connected to: %s\n", client.RemoteAddr())

	// wait for input
	bufio.NewReader(os.Stdin).ReadString('\n')

	err = client.SendMessage([]byte("Hello Host!"))
	if err != nil {
		return errors.Chain(err, "error sending message")
	}
	fmt.Println("Sent message.")

	_, err = client.ReadMessage()
	if err != nil {
		return errors.Chain(err, "error receiving response")
	}

	return nil
}
