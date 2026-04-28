package simulation

import (
	"context"
	"net"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	doptApi "github.com/felipedreis/doptimas-proto-go/api"
)

type mockSimulationServer struct {
	doptApi.UnimplementedSimulationServiceServer
	lastMethod string
}

func (s *mockSimulationServer) StatSimulation(ctx context.Context, req *doptApi.StatSimulationRequest) (*doptApi.StatSimulationResponse, error) {
	s.lastMethod = "StatSimulation"
	return &doptApi.StatSimulationResponse{Status: doptApi.SimulationStatus_READY}, nil
}

func (s *mockSimulationServer) StartSimulation(ctx context.Context, req *doptApi.StartSimulationRequest) (*doptApi.StatSimulationResponse, error) {
	s.lastMethod = "StartSimulation"
	return &doptApi.StatSimulationResponse{Status: doptApi.SimulationStatus_STARTED}, nil
}

func (s *mockSimulationServer) StopSimulation(ctx context.Context, req *doptApi.StopSimulationRequest) (*doptApi.StatSimulationResponse, error) {
	s.lastMethod = "StopSimulation"
	return &doptApi.StatSimulationResponse{Status: doptApi.SimulationStatus_STOPPED}, nil
}

func TestSimulationCommand_Execute(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	mockServer := &mockSimulationServer{}
	doptApi.RegisterSimulationServiceServer(s, mockServer)
	go func() {
		if err := s.Serve(lis); err != nil {
			return
		}
	}()
	defer s.Stop()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	// Test status
	t.Run("status", func(t *testing.T) {
		cmd := NewCommand([]string{"status"})
		cmd.Execute(conn, nil)
		if mockServer.lastMethod != "StatSimulation" {
			t.Errorf("Expected method StatSimulation, got %s", mockServer.lastMethod)
		}
	})

	// Test stop
	t.Run("stop", func(t *testing.T) {
		cmd := NewCommand([]string{"stop"})
		cmd.Execute(conn, nil)
		if mockServer.lastMethod != "StopSimulation" {
			t.Errorf("Expected method StopSimulation, got %s", mockServer.lastMethod)
		}
	})

	// Test start
	t.Run("start", func(t *testing.T) {
		tmpfile, err := os.CreateTemp("", "config*.json")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())

		if _, err := tmpfile.Write([]byte("{}")); err != nil {
			t.Fatal(err)
		}
		if err := tmpfile.Close(); err != nil {
			t.Fatal(err)
		}

		cmd := NewCommand([]string{"start", tmpfile.Name()})
		cmd.Execute(conn, nil)
		if mockServer.lastMethod != "StartSimulation" {
			t.Errorf("Expected method StartSimulation, got %s", mockServer.lastMethod)
		}
	})
}
