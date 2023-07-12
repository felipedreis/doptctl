package simulation

import (
	"google.golang.org/grpc"
)

type Command struct {
	subCommand string
	params     []string
}

func NewSimulationCommand(subCommand string, params []string) *Command {
	return &Command{subCommand: subCommand, params: params}
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) {

}
