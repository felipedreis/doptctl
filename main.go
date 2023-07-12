package main

import (
	//"context"
	"flag"
	//"time"
	"doptctl/commands"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		commands.Help()
	} else {
		commands.Run(args)
	}
}
