package list

import (
	"context"
	"doptctl/commands/output"
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
	ans, err := client.ListAgents(context.Background(), req)
	if err == nil {
		header := []string{"Name", "Metaheuristic", "Path"}
		data := [][]string{}
		for _, agent := range ans.Agents {
			data = append(data, []string{agent.Name, agent.Metaheuristic, agent.Path})
		}
		output.PrintTable(header, data)
	} else {
		log.Fatal(err)
	}
}

func listRegions(client doptApi.RegionServiceClient) {
	var req = new(doptApi.ListRegionsRequest)
	ans, err := client.ListRegions(context.Background(), req)

	if err == nil {
		header := []string{"Name", "Time", "Path", "Solutions"}
		data := [][]string{}
		for _, region := range ans.Regions {
			data = append(data, []string{
				region.Name,
				fmt.Sprintf("%d", region.Time),
				region.Path,
				fmt.Sprintf("%d", region.NumberOfSolutions),
			})
		}
		output.PrintTable(header, data)
	} else {
		log.Fatal(err)
	}
}
