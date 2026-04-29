package describe

import (
	"context"
	"doptctl/commands/output"
	"fmt"

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

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) error {
	if cmd.showHelp {
		cmd.Help()
		return nil
	}

	switch cmd.entityType {
	case "agent":
		client := api.NewAgentServiceClient(conn)
		return cmd.describeAgent(client)
	case "region":
		client := api.NewRegionServiceClient(conn)
		return cmd.describeRegion(client)
	default:
		cmd.Help()
		return nil
	}
}

func (cmd Command) Help() {
	Help()
}

func (cmd Command) describeAgent(client api.AgentServiceClient) error {
	req := &api.DescribeAgentRequest{AgentId: cmd.entityId}
	resp, err := client.DescribeAgent(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error describing agent: %w", err)
	}

	data := [][]string{
		{"Agent ID", resp.AgentId},
		{"Heuristic", resp.Heuristic},
		{"Lifetime", fmt.Sprintf("%d", resp.Lifetime)},
		{"Start Time", fmt.Sprintf("%d", resp.StartTime)},
		{"Current Time", fmt.Sprintf("%d", resp.CurrentTime)},
		{"Executions", fmt.Sprintf("%d", resp.CompleteExecutions)},
		{"Req. Solutions", fmt.Sprintf("%d", resp.RequiredSolutions)},
		{"Memory Tax", fmt.Sprintf("%.2f", resp.MemoryTax)},
	}

	if resp.BestSolution != nil && len(resp.BestSolution.Y) > 0 {
		data = append(data, []string{"Best Solution", fmt.Sprintf("%.4f", resp.BestSolution.Y[0])})
	}

	output.PrintVertical(data)
	return nil
}

func (cmd Command) describeRegion(client api.RegionServiceClient) error {
	req := &api.DescribeRegionRequest{RegionId: cmd.entityId}
	resp, err := client.DescribeRegion(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error describing region: %w", err)
	}

	data := [][]string{
		{"Region ID", resp.RegionId},
		{"Started", fmt.Sprintf("%t", resp.Started)},
		{"Started Time", fmt.Sprintf("%d", resp.StartedTime)},
		{"Current Time", fmt.Sprintf("%d", resp.CurrentTime)},
		{"Solutions", fmt.Sprintf("%d", resp.NumberOfSolutions)},
	}

	if resp.BestSolution != nil && len(resp.BestSolution.Y) > 0 {
		data = append(data, []string{"Best Solution", fmt.Sprintf("%.4f", resp.BestSolution.Y[0])})
	}

	output.PrintVertical(data)
	return nil
}
