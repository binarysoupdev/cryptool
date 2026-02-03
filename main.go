package main

import (
	"flag"
	"fmt"

	"github.com/binarysoupdev/cryptool/app"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Simple cryptography tool to encrypt/decrypt a file with a password.")
		flag.PrintDefaults()
	}

	file := flag.String("i", "", "the file to encrypt/decrypt")
	rm := flag.Bool("rm", false, "remove the old file")
	flag.Parse()

	err := app.Run(*file, *rm)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}
