package simulation

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"

	doptApi "github.com/felipedreis/doptimas-proto-go/api"
)

type Command struct {
	subCommand string
	params     []string
	showHelp   bool
}

func NewCommand(args []string) *Command {
	if len(args) == 0 {
		return &Command{showHelp: true}
	}
	return &Command{subCommand: args[0], params: args[1:]}
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) error {
	if cmd.showHelp {
		cmd.Help()
		return nil
	}

	switch cmd.subCommand {
	case "status":
		return cmd.handleStatus(conn)
	case "start":
		return cmd.handleStart(conn)
	case "stop":
		return cmd.handleStop(conn)
	case "extractData":
		return cmd.handleExtractData(conn)
	default:
		cmd.Help()
		return nil
	}
}

func (cmd Command) Help() {
	Help()
}

func (cmd Command) handleStatus(conn *grpc.ClientConn) error {
	client := doptApi.NewSimulationServiceClient(conn)
	resp, err := client.StatSimulation(context.Background(), &doptApi.StatSimulationRequest{})
	if err != nil {
		return fmt.Errorf("error getting simulation status: %w", err)
	}
	fmt.Printf("Status: %s\n", resp.Status)
	if resp.Message != "" {
		fmt.Printf("Message: %s\n", resp.Message)
	}
	return nil
}

func (cmd Command) handleStart(conn *grpc.ClientConn) error {
	if len(cmd.params) < 1 {
		cmd.Help()
		return nil
	}
	configPath := cmd.params[0]
	content, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading configuration file: %w", err)
	}

	client := doptApi.NewSimulationServiceClient(conn)
	resp, err := client.StartSimulation(context.Background(), &doptApi.StartSimulationRequest{
		ConfigurationFileContent: string(content),
	})
	if err != nil {
		return fmt.Errorf("error starting simulation: %w", err)
	}
	fmt.Printf("Simulation started: %s\n", resp.Status)
	if resp.Message != "" {
		fmt.Printf("Message: %s\n", resp.Message)
	}
	return nil
}

func (cmd Command) handleStop(conn *grpc.ClientConn) error {
	client := doptApi.NewSimulationServiceClient(conn)
	resp, err := client.StopSimulation(context.Background(), &doptApi.StopSimulationRequest{})
	if err != nil {
		return fmt.Errorf("error stopping simulation: %w", err)
	}
	fmt.Printf("Simulation stopped: %s\n", resp.Status)
	if resp.Message != "" {
		fmt.Printf("Message: %s\n", resp.Message)
	}
	return nil
}

func (cmd Command) handleExtractData(conn *grpc.ClientConn) error {
	fmt.Println("ExtractData is not supported by the current API")
	return nil
}
