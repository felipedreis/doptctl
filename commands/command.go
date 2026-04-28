package commands

import (
	"doptctl/commands/benchmark"
	"doptctl/commands/context"
	"doptctl/commands/describe"
	"doptctl/commands/list"
	"doptctl/commands/simulation"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strings"
)

type Command interface {
	Execute(conn *grpc.ClientConn, opts map[string]string)
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
		command.Execute(nil, nil)
	} else {
		ctx, error := context.LoadContext()

		if error != nil {
			log.Fatalf("Couldn't load current context")
		}

		conn := getConnection(ctx.URL())
		defer conn.Close()
		command.Execute(conn, globalOpts)
	}
}

func getConnection(serverAddr string) *grpc.ClientConn {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return conn
}

func getCommand(input []string) Command {
	if len(input) == 0 {
		Help()
		return nil
	}

	var commandName = input[0]

	switch commandName {
	case "benchmark":
		if len(input) < 2 {
			help([]string{"benchmark"})
			return nil
		}
		return benchmark.NewBenchmarkCommand(input[1])
	case "context":
		if len(input) < 2 {
			help([]string{"context"})
			return nil
		}
		return context.NewContextCommand(input[1], input[2:])
	case "describe":
		if len(input) < 3 {
			help([]string{"describe"})
			return nil
		}
		return describe.NewDescribeCommand(input[1], input[2])
	case "list":
		if len(input) < 2 {
			help([]string{"list"})
			return nil
		}
		return list.NewListCommand(input[1])
	case "simulation":
		if len(input) < 2 {
			help([]string{"simulation"})
			return nil
		}
		return simulation.NewSimulationCommand(input[1], input[2:])
	case "help":
		help(input[1:])
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
