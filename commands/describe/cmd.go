package describe

import (
	"doptctl/doptimas/api"
	"google.golang.org/grpc"
)

type Command struct {
	entityType string
	entityId   string
}

func NewDescribeCommand(entityType string, entityId string) (cmd *Command) {
	return &Command{entityType: entityType, entityId: entityId}
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) {
	client := api.NewSimulationClient(conn)

	switch cmd.entityType {
	case "agent":
		describeAgent(client)
	case "region":
		describeRegion(client)
	}
}

func describeAgent(client api.SimulationClient) {

}

func describeRegion(client api.SimulationClient) {

}
