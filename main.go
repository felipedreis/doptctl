package main

import (
	//"context"
	"flag"
	"fmt"

	//"time"
	"doptctl/commands"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	commands.Run(flag.Args())
}
