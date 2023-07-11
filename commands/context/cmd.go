package context

import (
	"fmt"
	"google.golang.org/grpc"
)

type Command struct {
	commandType string
	args        []string
	opts        map[string]string
}

func NewContextCommand(commandType string, args []string) *Command {
	var command = &Command{commandType: commandType, args: args, opts: make(map[string]string)}

	for _, arg := range command.args {
		if arg == "--set-current" || arg == "-s" {
			command.opts["--set-current"] = "true"
		}
	}

	return command
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) {
	switch cmd.commandType {
	case "list":
		listContexts()
	case "configure":
		configureNewContext(cmd.opts)
	case "set":
		setCurrentContext(cmd.args[0])
	}
}

func listContexts() {
	ctxs := getAllContexts()

	for _, ctx := range ctxs {
		fmt.Println(ctx.Name + ", " + ctx.Host + ", " + ctx.Port)
	}
}

func configureNewContext(opts map[string]string) {
	ctx := createNewContext()
	if _, ok := opts["--set-current"]; ok {
		setCurrentContext(ctx.Name)
	}
}

func setCurrentContext(contextName string) {
	overrideCurrentContext(contextName)
}
