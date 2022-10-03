package main

import (
	"flag"
	"fmt"

	"github.com/jina-ai/client-go"
)

func main() {
	host := flag.String("host", "", "host of the gateway")
	flag.Parse()

	if *host == "" {
		panic("Please pass a host to check the health of")
	}

	hcClient, err := client.NewGRPCHealthCheckClient(*host)
	if err != nil {
		panic(fmt.Errorf("unsuccessful healthcheck: %w", err))
	}

	status, err := hcClient.HealthCheck()
	if err != nil {
		panic(fmt.Errorf("failed to check health: %w", err))
	}
	if status {
		fmt.Println("Flow running at", *host, "is healthy!")
	} else {
		panic(fmt.Errorf("unsuccessful healthcheck"))
	}
}
