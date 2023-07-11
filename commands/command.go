package commands

import (
	"doptctl/commands/context"
	"doptctl/commands/list"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strings"
)

type Command interface {
	Execute(conn *grpc.ClientConn, opts map[string]string)
}

func getCommand(input []string) Command {
	var commandName = input[0]

	switch commandName {
	case "list":
		return list.NewListCommand(input[1])
	case "context":
		return context.NewContextCommand(input[1], input[2:])
	}
	return nil
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

func Run(input []string) {

	var lastIndex int
	var globalOpts map[string]string

	for i, value := range input {
		if strings.Contains(value, "--") {
			opts := strings.Split(value, "=")
			globalOpts[opts[0]] = opts[1]
		} else {
			lastIndex = i
			break
		}
	}

	var command = getCommand(input[lastIndex:])

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
