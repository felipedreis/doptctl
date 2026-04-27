package describe

import (
	"github.com/felipedreis/doptimas-proto-go/api"
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
	switch cmd.entityType {
	case "agent":
		client := api.NewAgentServiceClient(conn)
		describeAgent(client)
	case "region":
		client := api.NewRegionServiceClient(conn)
		describeRegion(client)
	}
}

func describeAgent(client api.AgentServiceClient) {

}

func describeRegion(client api.RegionServiceClient) {

}
