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
	showHelp   bool
}

func NewCommand(args []string) *Command {
	if len(args) == 0 {
		return &Command{showHelp: true}
	}
	return &Command{subCommand: args[0], params: args[1:]}
}

func (cmd Command) Execute(conn *grpc.ClientConn, opts map[string]string) {
	if cmd.showHelp {
		cmd.Help()
		return
	}

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
		cmd.Help()
	}
}

func (cmd Command) Help() {
	Help()
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
		cmd.Help()
		return
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
