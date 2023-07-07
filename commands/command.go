package commands

import (
	"doptctl/commands/list"
)

type Command interface {
	Execute(grpc, subcommand []string, opts map[string]string)
}

func GetCommand(input []string) *Command {
	var commandName = input[0]

	switch commandName {
	case "list":
		return list.NewListCommand(input[1], input[2])

	}
}
