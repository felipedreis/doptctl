package list

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	doptApi "doptctl/doptimas/api"
)

type Command struct {
	entityType string
}

func NewListCommand(entityType string) Command {
	return Command{entityType: entityType}
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) {
	simulationClient := doptApi.NewSimulationClient(conn)

	switch cmd.entityType {
	case "agents":
		listAgents(simulationClient)
	case "regions":
		listRegions(simulationClient)
	}
}

func listAgents(client doptApi.SimulationClient) {
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

func listRegions(client doptApi.SimulationClient) {
	var req = new(doptApi.ListRegionsRequest)
	ans, error := client.ListRegions(context.Background(), req)

	if error == nil {
		fmt.Println(ans)
	}
}
