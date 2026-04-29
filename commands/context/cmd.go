package context

import (
	"doptctl/commands/output"
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

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) error {
	if cmd.showHelp {
		cmd.Help()
		return nil
	}

	switch cmd.commandType {
	case "list":
		return listContexts()
	case "configure":
		return configureNewContext(cmd.opts)
	case "set":
		if len(cmd.args) == 0 {
			cmd.Help()
			return nil
		}
		return setCurrentContext(cmd.args[0])
	default:
		cmd.Help()
		return nil
	}
}

func (cmd Command) Help() {
	Help()
}

func listContexts() error {
	ctxs, err := getAllContexts()
	if err != nil {
		return err
	}

	header := []string{"", "Name", "Host", "Port"}
	data := [][]string{}
	for _, ctx := range ctxs {
		currentIndicator := ""
		if ctx.current {
			currentIndicator = "*"
		}
		data = append(data, []string{currentIndicator, ctx.Name, ctx.Host, ctx.Port})
	}

	output.PrintTable(header, data)
	return nil
}

func configureNewContext(opts map[string]string) error {
	ctx, err := createNewContext()
	if err != nil {
		return err
	}
	if ctx == nil {
		return nil
	}
	if _, ok := opts["--set-current"]; ok {
		return setCurrentContext(ctx.Name)
	}
	return nil
}

func setCurrentContext(contextName string) error {
	return overrideCurrentContext(contextName)
}
