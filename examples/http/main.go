package main

import (
	"fmt"

	"github.com/jina-ai/client-go"
	"github.com/jina-ai/client-go/docarray"
	"github.com/jina-ai/client-go/jina"
)

// Create a Document
func getDoc(id string) *docarray.DocumentProto {
	return &docarray.DocumentProto{
		Id: id,
		Content: &docarray.DocumentProto_Text{
			Text: "Hello world. This is a test document with id:" + id,
		},
	}
}

// Create a DocumentArray with 3 Documents
func getDocarrays(numDocs int) *docarray.DocumentArrayProto {
	var docs []*docarray.DocumentProto
	for i := 0; i < numDocs; i++ {
		docs = append(docs, getDoc(fmt.Sprint(i)))
	}
	return &docarray.DocumentArrayProto{
		Docs: docs,
	}
}

// Create DataRequest with a DocumentArray
func getDataRequest(numDocs int) *jina.DataRequestProto {
	return &jina.DataRequestProto{
		Data: &jina.DataRequestProto_DataContentProto{
			Documents: &jina.DataRequestProto_DataContentProto_Docs{
				Docs: getDocarrays(numDocs),
			},
		},
	}
}

// Generate a stream of requests
func generateDataRequests(numRequests int) <-chan *jina.DataRequestProto {
	requests := make(chan *jina.DataRequestProto)
	go func() {
		for i := 0; i < numRequests; i++ {
			requests <- getDataRequest(3)
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
			fmt.Printf("\nDocID: %d, Chunks: %d", docidx, len(doc.Chunks))
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
	HTTPClient, err := client.NewHTTPClient("http://localhost:12345")
	if err != nil {
		panic(err)
	}
	HTTPClient.POST(generateDataRequests(5), OnDone, OnError, nil)
}
