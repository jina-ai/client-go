package main

import (
	"context"
	"fmt"

	"github.com/jina-ai/client-go/jina"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	host := "localhost:12345"
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Gateway", host, "is unhealthy")
		return
	}

	client := jina.NewJinaGatewayDryRunRPCClient(conn)
	status, err := client.DryRun(context.Background(), &emptypb.Empty{})
	if err != nil {
		fmt.Println("Gateway", host, "is unhealthy")
		return
	}
	if status.Code == jina.StatusProto_SUCCESS {
		fmt.Println("Gateway", host, "is unhealthy")
		return
	}
	fmt.Println("Gateway", host, "is unhealthy")
}
