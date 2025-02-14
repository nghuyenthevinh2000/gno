package main

import (
	"os"

	"github.com/gnolang/gno/pkgs/command"
	"github.com/gnolang/gno/pkgs/crypto/keys/client"
)

func main() {

	cmd := command.NewStdCommand()

	// set default options.

	// customize call to command.
	// insert args and options here.
	// TODO: use flags or */pflags.

	exec := os.Args[0]
	args := os.Args[1:]

	err := client.RunMain(cmd, exec, args)
	if err != nil {
		cmd.ErrPrintln(err.Error())
		return // exit
	}
}
