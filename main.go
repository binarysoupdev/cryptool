package main

import (
	"flag"
	"fmt"

	"github.com/binarysoupdev/cryptool/app"
	"github.com/binarysoupdev/got-style/style"
)

func main() {
	flag.Usage = func() {
		style.Info.Println("Simple cryptography tool to encrypt/decrypt a file with a password:")
		flag.PrintDefaults()
	}

	file := flag.String("i", "", "the file to encrypt/decrypt")
	rm := flag.Bool("rm", false, "remove the old file")
	flag.Parse()

	err := app.Run(*file, *rm)
	if err != nil {
		style.BoldError.Print("ERROR: ")
		fmt.Println(err)
	}
}
