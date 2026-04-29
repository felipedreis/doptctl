package benchmark

import (
	"google.golang.org/grpc"

	// doptApi "github.com/felipedreis/doptimas-proto-go/api"
)

type Command struct {
	subCommand string
	showHelp   bool
}

func NewCommand(args []string) *Command {
	if len(args) == 0 {
		return &Command{showHelp: true}
	}
	return &Command{subCommand: args[0]}
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) error {
	if cmd.showHelp {
		cmd.Help()
		return nil
	}
	//client := doptApi.NewBenchmarkServiceClient(conn)
	return nil
}

func (cmd Command) Help() {
	Help()
}
