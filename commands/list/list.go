package list

import (
	"fmt"

	"google.golang.org/grpc"

	doptApi "doptctl/doptimas/api"
)

type Command struct {
	entityType string
	entityId   string
}

func NewListCommand(entityType string, entityId string) *Command {
	return &Command{entityType: entityType, entityId: entityId}
}

func (*Command) Execute(conn *grpc.ClientConn, subcommand []string, opts map[string]string) {
	simulationClient := doptApi.NewSimulationClient(conn)
	fmt.Println(simulationClient)
	fmt.Println("calling list execute")
}
