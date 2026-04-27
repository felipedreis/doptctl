package simulation

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	doptApi "github.com/felipedreis/doptimas-proto-go/api"
)

type Command struct {
	subCommand string
	params     []string
}

func NewSimulationCommand(subCommand string, params []string) *Command {
	return &Command{subCommand: subCommand, params: params}
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) {
	switch cmd.subCommand {
	case "status":
		cmd.handleStatus(conn)
	case "start":
		cmd.handleStart(conn)
	case "stop":
		cmd.handleStop(conn)
	case "extractData":
		cmd.handleExtractData(conn)
	default:
		log.Fatalf("Unknown simulation subcommand: %s", cmd.subCommand)
	}
}

func (cmd Command) handleStatus(conn *grpc.ClientConn) {
	client := doptApi.NewSimulationServiceClient(conn)
	resp, err := client.StatSimulation(context.Background(), &doptApi.StatSimulationRequest{})
	if err != nil {
		log.Fatalf("Error getting simulation status: %v", err)
	}
	fmt.Printf("Status: %s\n", resp.Status)
	if resp.Message != "" {
		fmt.Printf("Message: %s\n", resp.Message)
	}
}

func (cmd Command) handleStart(conn *grpc.ClientConn) {
	if len(cmd.params) < 1 {
		log.Fatal("Missing configuration file path")
	}
	configPath := cmd.params[0]
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	client := doptApi.NewSimulationServiceClient(conn)
	resp, err := client.StartSimulation(context.Background(), &doptApi.StartSimulationRequest{
		ConfigurationFileContent: string(content),
	})
	if err != nil {
		log.Fatalf("Error starting simulation: %v", err)
	}
	fmt.Printf("Simulation started: %s\n", resp.Status)
	if resp.Message != "" {
		fmt.Printf("Message: %s\n", resp.Message)
	}
}

func (cmd Command) handleStop(conn *grpc.ClientConn) {
	client := doptApi.NewSimulationServiceClient(conn)
	resp, err := client.StopSimulation(context.Background(), &doptApi.StopSimulationRequest{})
	if err != nil {
		log.Fatalf("Error stopping simulation: %v", err)
	}
	fmt.Printf("Simulation stopped: %s\n", resp.Status)
	if resp.Message != "" {
		fmt.Printf("Message: %s\n", resp.Message)
	}
}

func (cmd Command) handleExtractData(conn *grpc.ClientConn) {
	fmt.Println("ExtractData is not supported by the current API")
}
