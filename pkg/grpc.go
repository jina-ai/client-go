package client

import (
	"context"
	"fmt"

	"github.com/deepankarm/client-go/pkg/jina"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	Host      string
	conn      *grpc.ClientConn
	rpcClient jina.JinaRPCClient
	ctx       context.Context
}

func NewGRPCClient(host string) (*GRPCClient, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		Host:      host,
		conn:      conn,
		rpcClient: jina.NewJinaRPCClient(conn),
		ctx:       context.Background(),
	}, nil
}

func (c *GRPCClient) POST(requests <-chan jina.DataRequestProto) error {
	stream, err := c.rpcClient.Call(c.ctx)
	if err != nil {
		return err
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				break
			}
			_ = resp
			fmt.Println(resp)

			// requests <- *resp
		}
	}()

	func() {
		for {
			req, ok := <-requests
			if !ok {
				break
			}
			if err := stream.Send(&req); err != nil {
				panic(err)
			}
		}
		stream.CloseSend()
	}()

	return nil
}
