package commands

import (
	"doptctl/commands/benchmark"
	"doptctl/commands/context"
	"doptctl/commands/describe"
	"doptctl/commands/list"
	"doptctl/commands/simulation"
	"doptctl/doptimas/client"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"strings"
)

type Command interface {
	Execute(conn *grpc.ClientConn, opts map[string]string) error
	Help()
}

func Run(input []string) {
	var lastIndex int
	globalOpts := make(map[string]string)

	for i, value := range input {
		if strings.HasPrefix(value, "--") {
			opts := strings.Split(value, "=")
			if len(opts) == 2 {
				globalOpts[opts[0]] = opts[1]
			}
		} else {
			lastIndex = i
			break
		}
	}

	var command = getCommand(input[lastIndex:])
	if command == nil {
		return
	}

	// context commands operate over local files, and change current context
	// this if prevents loading the current context
	if input[lastIndex] == "context" {
		err := command.Execute(nil, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		ctx, err := context.LoadContext()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		conn, err := client.NewConnection(ctx.URL())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error connecting to %s: %v\n", ctx.URL(), err)
			os.Exit(1)
		}
		defer conn.Close()
		err = command.Execute(conn, globalOpts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}

func getCommand(input []string) Command {
	if len(input) == 0 {
		Help()
		return nil
	}

	var commandName = input[0]
	var args = input[1:]

	switch commandName {
	case "benchmark":
		return benchmark.NewCommand(args)
	case "context":
		return context.NewCommand(args)
	case "describe":
		return describe.NewCommand(args)
	case "list":
		return list.NewCommand(args)
	case "simulation":
		return simulation.NewCommand(args)
	case "help":
		help(args)
	default:
		Help()
	}
	return nil
}

func help(input []string) {
	if len(input) != 0 {
		var commandName = input[0]

		switch commandName {
		case "benchmark":
			benchmark.Help()
		case "context":
			context.Help()
		case "describe":
			describe.Help()
		case "list":
			list.Help()
		case "simulation":
			simulation.Help()
		default:
			fmt.Printf("error: Unknown command %s for doptctl\n", commandName)
			Help()
		}
	} else {
		Help()
	}
}
