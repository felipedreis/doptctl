package benchmark

import (
	"google.golang.org/grpc"

	// doptApi "github.com/felipedreis/doptimas-proto-go/api"
)

type Command struct {
	subCommand string
}

func NewBenchmarkCommand(subCommand string) *Command {
	return &Command{subCommand: subCommand}
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) {
	//client := doptApi.NewBenchmarkServiceClient(conn)
}
