package describe

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

	"github.com/felipedreis/doptimas-proto-go/api"
)

type mockAgentServer struct {
	api.UnimplementedAgentServiceServer
}

func (s *mockAgentServer) DescribeAgent(ctx context.Context, req *api.DescribeAgentRequest) (*api.DescribeAgentResponse, error) {
	if req.AgentId == "agent-1" {
		return &api.DescribeAgentResponse{
			AgentId:   "agent-1",
			Heuristic: "GA",
			Lifetime:  1000,
		}, nil
	}
	return nil, nil // Simplified for test
}

type mockRegionServer struct {
	api.UnimplementedRegionServiceServer
}

func (s *mockRegionServer) DescribeRegion(ctx context.Context, req *api.DescribeRegionRequest) (*api.DescribeRegionResponse, error) {
	if req.RegionId == "region-1" {
		return &api.DescribeRegionResponse{
			RegionId: "region-1",
			Started:  true,
		}, nil
	}
	return nil, nil // Simplified for test
}

func TestDescribeCommand_Execute(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	api.RegisterAgentServiceServer(s, &mockAgentServer{})
	api.RegisterRegionServiceServer(s, &mockRegionServer{})
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

	t.Run("describe agent", func(t *testing.T) {
		output := captureOutput(func() {
			cmd := NewCommand([]string{"agent", "agent-1"})
			cmd.Execute(conn, nil)
		})
		if !strings.Contains(output, "agent-1") || !strings.Contains(output, "GA") {
			t.Errorf("Expected output to contain agent details, got: %s", output)
		}
	})

	t.Run("describe region", func(t *testing.T) {
		output := captureOutput(func() {
			cmd := NewCommand([]string{"region", "region-1"})
			cmd.Execute(conn, nil)
		})
		if !strings.Contains(output, "region-1") || !strings.Contains(output, "true") {
			t.Errorf("Expected output to contain region details, got: %s", output)
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
