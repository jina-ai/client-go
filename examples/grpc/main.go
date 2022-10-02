package main

import (
	"fmt"

	"github.com/deepankarm/client-go"
	"github.com/deepankarm/client-go/docarray"
	"github.com/deepankarm/client-go/jina"
)

// Generate random DocumentArrays
func docs(numDocs int) *jina.DataRequestProto_DataContentProto_Docs {
	var docs []*docarray.DocumentProto
	for i := 0; i < numDocs; i++ {
		docs = append(docs, &docarray.DocumentProto{
			Id: fmt.Sprint(i),
			Content: &docarray.DocumentProto_Text{
				Text: fmt.Sprintf("Hello world. This is a test document %d", i),
			},
		})
	}
	return &jina.DataRequestProto_DataContentProto_Docs{
		Docs: &docarray.DocumentArrayProto{
			Docs: docs,
		},
	}
}

// Generate a stream of requests
func requestsGen(numRequests int) <-chan *jina.DataRequestProto {
	requests := make(chan *jina.DataRequestProto)
	go func() {
		for i := 0; i < numRequests; i++ {
			requests <- &jina.DataRequestProto{
				Data: &jina.DataRequestProto_DataContentProto{
					Documents: docs(3),
				},
			}
		}
		defer close(requests)
	}()
	return requests
}

func OnDone(resp *jina.DataRequestProto) {
	switch docs := resp.Data.Documents.(type) {
	case *jina.DataRequestProto_DataContentProto_Docs:
		fmt.Printf("\n\nTotal %d docs received.", len(docs.Docs.Docs))
		for docidx, doc := range docs.Docs.Docs {
			fmt.Printf("\nDocID: %d", docidx)
			for i, chunk := range doc.Chunks {
				fmt.Printf("\n\tChunk %d text: %s ", i, chunk.Content.(*docarray.DocumentProto_Text).Text)
			}
		}
	}
}

func OnError(resp *jina.DataRequestProto) {
	fmt.Println("Got an error for request", resp)
}

func main() {
	GRPCClient, err := client.NewGRPCClient("localhost:12345")
	if err != nil {
		panic(err)
	}
	GRPCClient.POST(requestsGen(5), OnDone, OnError, nil)
}
