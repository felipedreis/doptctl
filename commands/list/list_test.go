package list

import (
	"bytes"
	"context"
	"io"
	"net"
	"os"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	doptApi "github.com/felipedreis/doptimas-proto-go/api"
)

type mockAgentServer struct {
	doptApi.UnimplementedAgentServiceServer
}

func (s *mockAgentServer) ListAgents(ctx context.Context, req *doptApi.ListAgentRequest) (*doptApi.ListAgentResponse, error) {
	return &doptApi.ListAgentResponse{
		Agents: []*doptApi.Agent{
			{Name: "agent-1", Metaheuristic: "GA", Path: "/path/1"},
			{Name: "agent-2", Metaheuristic: "SA", Path: "/path/2"},
		},
	}, nil
}

type mockRegionServer struct {
	doptApi.UnimplementedRegionServiceServer
}

func (s *mockRegionServer) ListRegions(ctx context.Context, req *doptApi.ListRegionsRequest) (*doptApi.ListRegionsResponse, error) {
	return &doptApi.ListRegionsResponse{
		Regions: []*doptApi.Region{
			{Name: "region-1", Time: 100, Path: "/r1", NumberOfSolutions: 10},
		},
	}, nil
}

func TestListCommand_Execute(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	doptApi.RegisterAgentServiceServer(s, &mockAgentServer{})
	doptApi.RegisterRegionServiceServer(s, &mockRegionServer{})
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

	t.Run("list agents", func(t *testing.T) {
		output := captureOutput(func() {
			cmd := NewCommand([]string{"agents"})
			cmd.Execute(conn, nil)
		})
		if !strings.Contains(output, "agent-1") || !strings.Contains(output, "agent-2") {
			t.Errorf("Expected output to contain agent names, got: %s", output)
		}
	})

	t.Run("list regions", func(t *testing.T) {
		output := captureOutput(func() {
			cmd := NewCommand([]string{"regions"})
			cmd.Execute(conn, nil)
		})
		if !strings.Contains(output, "region-1") {
			t.Errorf("Expected output to contain region name, got: %s", output)
		}
	})
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
