package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/binarysoupdev/cryptool/internal/app"
	"github.com/binarysoupdev/go-commando/command"
	"github.com/binarysoupdev/got-style/style"
)

func main() {
	ls := flag.Bool("ls", false, "list all commands")
	flag.Parse()

	runner := command.NewRunner(
		app.NewKeyGenCommand(),
		app.NewEncryptCommand(),
		app.NewDecryptCommand(),
		app.NewHostCommand(),
		app.NewClientCommand(),
	)

	if *ls || len(os.Args) < 2 {
		runner.ListCommands()
		return
	}

	if err := runner.RunCommand(os.Args[1], os.Args[2:]); err != nil {
		fmt.Printf("%s %s\n", style.BoldError.Sprint("ERROR:"), err.Error())
	}
}
