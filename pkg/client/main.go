package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/rexposadas/teleport/api"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.TeleportServiceClient
}

// NewClient returns a client prepared to connect to the server.
func NewClient() *Client {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &Client{
		conn:   conn,
		client: pb.NewTeleportServiceClient(conn),
	}
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Start(ctx context.Context, cmd string, args []string) (*pb.StartResponse, error) {
	resp, err := c.client.Start(ctx, &pb.StartRequest{Command: cmd, Args: args})
	if err != nil {
		return nil, fmt.Errorf("call start job: %w", err)
	}

	return resp, nil
}

func (c *Client) GetStatus(ctx context.Context, jobID string) (*pb.GetStatusResponse, error) {
	resp, err := c.client.GetStatus(ctx, &pb.GetStatusRequest{JobId: jobID})
	if err != nil {
		return nil, fmt.Errorf("failed to get status of job: %w", err)
	}

	return resp, nil
}
