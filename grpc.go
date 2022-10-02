package client

import (
	"context"
	"sync"

	"github.com/jina-ai/client-go/jina"
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

func (c *GRPCClient) POST(requests <-chan *jina.DataRequestProto, onDone, onError, onAlways CallbackType) error {
	var wg sync.WaitGroup

	stream, err := c.rpcClient.Call(c.ctx)
	if err != nil {
		return err
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				if onAlways != nil {
					onAlways(resp)
				}
				if onError != nil {
					onError(resp)
				}
			}
			if onAlways != nil {
				onAlways(resp)
			}
			if onDone != nil {
				onDone(resp)
			}
			wg.Done()
		}
	}()

	for {
		req, ok := <-requests
		if !ok {
			break
		}
		if err := stream.Send(req); err != nil {
			panic(err)
		}
		wg.Add(1)
	}
	wg.Wait()
	stream.CloseSend()
	return nil
}
