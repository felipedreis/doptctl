package context

import (
	"fmt"
	"google.golang.org/grpc"
)

type Command struct {
	commandType string
	args        []string
	opts        map[string]string
	showHelp    bool
}

func NewCommand(args []string) *Command {
	if len(args) == 0 {
		return &Command{showHelp: true}
	}
	var command = &Command{commandType: args[0], args: args[1:], opts: make(map[string]string)}

	for _, arg := range command.args {
		if arg == "--set-current" || arg == "-s" {
			command.opts["--set-current"] = "true"
		}
	}

	return command
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) {
	if cmd.showHelp {
		cmd.Help()
		return
	}

	switch cmd.commandType {
	case "list":
		listContexts()
	case "configure":
		configureNewContext(cmd.opts)
	case "set":
		if len(cmd.args) == 0 {
			cmd.Help()
			return
		}
		setCurrentContext(cmd.args[0])
	default:
		cmd.Help()
	}
}

func (cmd Command) Help() {
	Help()
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
