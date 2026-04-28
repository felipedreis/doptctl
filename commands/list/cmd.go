package list

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	doptApi "github.com/felipedreis/doptimas-proto-go/api"
)

type Command struct {
	entityType string
	showHelp   bool
}

func NewCommand(args []string) *Command {
	if len(args) == 0 {
		return &Command{showHelp: true}
	}
	return &Command{entityType: args[0]}
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) {
	if cmd.showHelp {
		cmd.Help()
		return
	}

	switch cmd.entityType {
	case "agents":
		agentClient := doptApi.NewAgentServiceClient(conn)
		listAgents(agentClient)
	case "regions":
		regionClient := doptApi.NewRegionServiceClient(conn)
		listRegions(regionClient)
	default:
		cmd.Help()
	}
}

func (cmd Command) Help() {
	Help()
}

func listAgents(client doptApi.AgentServiceClient) {
	var req = new(doptApi.ListAgentRequest)
	ans, error := client.ListAgents(context.Background(), req)
	if error == nil {
		fmt.Println(ans)
		for _, agent := range ans.Agents {
			fmt.Println(agent)
		}

	} else {
		log.Fatal(error)
	}
}

func listRegions(client doptApi.RegionServiceClient) {
	var req = new(doptApi.ListRegionsRequest)
	ans, error := client.ListRegions(context.Background(), req)

	if error == nil {
		fmt.Println(ans)
	}
}
