package describe

import (
	"github.com/felipedreis/doptimas-proto-go/api"
	"google.golang.org/grpc"
)

type Command struct {
	entityType string
	entityId   string
	showHelp   bool
}

func NewCommand(args []string) *Command {
	if len(args) < 2 {
		return &Command{showHelp: true}
	}
	return &Command{entityType: args[0], entityId: args[1]}
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) {
	if cmd.showHelp {
		cmd.Help()
		return
	}

	switch cmd.entityType {
	case "agent":
		client := api.NewAgentServiceClient(conn)
		describeAgent(client)
	case "region":
		client := api.NewRegionServiceClient(conn)
		describeRegion(client)
	default:
		cmd.Help()
	}
}

func (cmd Command) Help() {
	Help()
}

func describeAgent(client api.AgentServiceClient) {

}

func describeRegion(client api.RegionServiceClient) {

}
