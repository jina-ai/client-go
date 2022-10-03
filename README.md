# Jina Golang Client

### Install

```bash
go get github.com/jina-ai/client-go
```

### Basic Usage

```go
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
func getDocarray() *docarray.DocumentArrayProto {
	return &docarray.DocumentArrayProto{
		Docs: []*docarray.DocumentProto{getDoc("1"), getDoc("2"), getDoc("3")},
	}
}

// Create DataRequest with a DocumentArray
func getDataRequest() *jina.DataRequestProto {
	return &jina.DataRequestProto{
		Data: &jina.DataRequestProto_DataContentProto{
			Documents: &jina.DataRequestProto_DataContentProto_Docs{
				Docs: getDocarray(),
			},
		},
	}
}

// Generate `DataRequest`s with random DocumentArrays
func generateDataRequests() <-chan *jina.DataRequestProto {
	requests := make(chan *jina.DataRequestProto)
	go func() {
		// Generate 10 requests
		for i := 0; i < 10; i++ {
			requests <- getDataRequest()
		}
		defer close(requests)
	}()
	return requests
}

// Custom OnDone callback
func OnDone(resp *jina.DataRequestProto) {
	fmt.Println("Got a successful response!")
}

// Custom OnError callback
func OnError(resp *jina.DataRequestProto) {
	fmt.Println("Got an error in response!")
}

func main() {
    	// Create a HTTP client (expects a Jina Flow with http protocol running on localhost:12345)
	HTTPClient, err := client.NewHTTPClient("http://localhost:12345")
	if err != nil {
		panic(err)
	}
    
    	// Send requests to the Flow
	HTTPClient.POST(generateDataRequests(), OnDone, OnError, nil)
}

```



### Examples


| Example |  |
| :---   | ---:  |
| [gRPC](examples/grpc/) | Stream requests using gRPC Client |
| [HTTP](examples/http/) | Stream requests using HTTP Client |
| [WebSocket](examples/websocket/) | Stream requests using WebSocket Client |
| [gRPC Healthcheck](examples/healthcheck/grpc/) | Check if the gRPC Flow is healthy |
| [HTTP Healthcheck](examples/healthcheck/http/) | Check if the  HTTP Flow is healthy |
| [WebSocket Healthcheck](examples/healthcheck/websocket/) | Check if the WebSocket Flow is healthy |
| DocArray usage | Example usage of DocArray (TODO) |
| Concurrent requests | Send concurrent requests to Jina Gateway (TODO) |


### Gotchas

##### Directory structure

```bash
.
├── protos
│   ├── docarray.proto          # proto file for DocArray
│   └── jina.proto              # proto file for Jina
├── docarray                    # docarray package
│   ├── docarray.pb.go          # generated from docarray.proto  
│   └── json.go                 # custom json (un)marshaler for few fields in docarray.proto
├── jina                        # jina package
│   ├── jina_grpc.pb.go         # generated from jina.proto
│   ├── jina.pb.go              # generated from jina.proto
│   └── json.go                 # custom json (un)marshaler for few fields in jina.proto
├── client.go                   # Client interface
├── grpc.go                     # gRPC client
├── http.go                     # HTTP client
├── websocket.go                # WebSocket client
├── scripts
└   └── protogen.sh             # script to Golang code from proto files
```

- `scripts/protogen.sh` generates the Golang code from the protos. Each proto generates code in a separate package. This is to avoid name clashes.

- `jina/json.go` and `docarray/json.go` are custom json (un)marshalers for few fields in `jina.proto` and `docarray.proto` respectively. 

- `client.go` defines the `Client` & `HealthCheckClient` interface. This is implemented by `grpc.go`, `http.go` and `websocket.go`.


#### Jina/Docarray Version Compatibility 

Current jina version is mentioned in the [requirements.txt](requirements.txt) file. Every time, there's a PR that bumps the jina version, [Version Update](.github/workflows/version-update.yml) workflow 
- Downloads the right protos for jina & docarray.
- Generates the Golang code from the protos.
- Runs integration tests.
- If all tests pass, it commits the latest code into the same branch. 

Once the PR is merged, a release is created with the new Jina version via [Tag & Release](.github/workflows/tag.yml) workflow.

---

Another (better?) approach for keeping all docarray/jina versions compatible would be,

- For all docarray versions, generate the Golang code from the protos under `docarray/v1`, `docarray/v2` packages.
- For all jina versions, generate the Golang code from the protos under `jina/v1`, `jina/v2` packages.
- Skip re-releasing the client-go package for every jina version, rather user can pick the right version of jina/docarray package.
