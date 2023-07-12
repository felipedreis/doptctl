package benchmark

import (
	"google.golang.org/grpc"
)

type Command struct {
	subCommand string
}

func NewBenchmarkCommand(subCommand string) *Command {
	return &Command{subCommand: subCommand}
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) {
	//client := api.NewBenchmarkClient(conn)
}
