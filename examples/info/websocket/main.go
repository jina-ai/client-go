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
	infoClient, err := client.NewWebSocketInfoClient(*host)
	if err != nil {
		panic(fmt.Errorf("unsuccessful info: %w", err))
	}

	info, err := infoClient.InfoJSON()
	if err != nil {
		panic(fmt.Errorf("failed to check info: %w", err))
	}
	fmt.Println(string(info))
}
