package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"sync"

	"github.com/jina-ai/client-go/jina"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCClient struct {
	Host      string
	conn      *grpc.ClientConn
	rpcClient jina.JinaRPCClient
	ctx       context.Context
}

func getSecureDialOptions() grpc.DialOption {
	return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: false}))
}

func getInsecureDialOptions() grpc.DialOption {
	return grpc.WithTransportCredentials(insecure.NewCredentials())
}

func getHostAndDialOptions(host string) (string, grpc.DialOption) {
	if strings.HasPrefix(host, "grpcs://") {
		host = strings.TrimPrefix(host, "grpcs://")
		return fmt.Sprint(host + ":443"), getSecureDialOptions()
	}
	if strings.HasPrefix(host, "grpc://") {
		host = strings.TrimPrefix(host, "grpc://")
		return host, getInsecureDialOptions()
	}
	return host, getInsecureDialOptions()
}

func NewGRPCClient(host string) (*GRPCClient, error) {
	host, dialOptions := getHostAndDialOptions(host)
	conn, err := grpc.Dial(host, dialOptions)
	if err != nil {
		fmt.Println("Error in grpc dial", err)
		return nil, err
	}
	return &GRPCClient{
		Host:      host,
		conn:      conn,
		rpcClient: jina.NewJinaRPCClient(conn),
		ctx:       context.Background(),
	}, nil
}

func (c GRPCClient) POST(requests <-chan *jina.DataRequestProto, onDone, onError, onAlways CallbackType) error {
	var wg sync.WaitGroup

	stream, err := c.rpcClient.Call(c.ctx)
	if err != nil {
		return err
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if resp == nil {
				break
			}
			if err != nil {
				fmt.Println("Error in receiving response", err)
				if onError != nil {
					onError(resp)
				}
			} else if onDone != nil {
				onDone(resp)
			}
			if onAlways != nil {
				onAlways(resp)
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

type GRPCHealthCheckClient struct {
	Host      string
	conn      *grpc.ClientConn
	rpcClient jina.JinaGatewayDryRunRPCClient
	ctx       context.Context
}

func NewGRPCHealthCheckClient(host string) (*GRPCHealthCheckClient, error) {
	host, dialOptions := getHostAndDialOptions(host)
	conn, err := grpc.Dial(host, dialOptions)
	if err != nil {
		return nil, err
	}
	return &GRPCHealthCheckClient{
		Host:      host,
		conn:      conn,
		rpcClient: jina.NewJinaGatewayDryRunRPCClient(conn),
		ctx:       context.Background(),
	}, nil
}

func (c *GRPCHealthCheckClient) HealthCheck() (bool, error) {
	status, err := c.rpcClient.DryRun(c.ctx, &emptypb.Empty{})
	if err != nil {
		return false, err
	}
	if status.Code == jina.StatusProto_SUCCESS {
		return true, nil
	}
	return false, nil
}
