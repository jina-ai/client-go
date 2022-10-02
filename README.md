# Jina Golang Client

### Install

```bash
go get github.com/deepankarm/client-go
```

### Examples


| Example |  |
| :---   | ---:  |
| [gRPC](examples/grpc/README.md) | Stream requests using gRPC Client |
| [HTTP](examples/http/README.md) | Stream requests using HTTP Client |
| [WebSocket](examples/websocket/README.md) | Stream requests using WebSocket Client |
| DocArray usage | Example usage of DocArray (TODO) |


### Gotchas

##### Directory structure

```bash
.
├── client.go                   # Client interface
├── docarray                    # docarray package
│   ├── docarray.pb.go          # generated from docarray.proto  
│   └── json.go                 # custom json (un)marshaler for few fields in docarray.proto
├── grpc.go                     # gRPC client
├── http.go                     # HTTP client
├── jina                        # jina package
│   ├── jina_grpc.pb.go         # generated from jina.proto
│   ├── jina.pb.go              # generated from jina.proto
│   └── json.go                 # custom json (un)marshaler for few fields in jina.proto
├── protos
│   ├── docarray.proto          # proto file for DocArray
│   └── jina.proto              # proto file for Jina
├── scripts
│   └── protogen.sh             # script to Golang code from proto files
└── websocket.go                # WebSocket client
```

- `scripts/protogen.sh` generates the Golang code from the protos. Each proto generates code in a separate package. This is to avoid name clashes.

- `jina/json.go` and `docarray/json.go` are custom json (un)marshalers for few fields in `jina.proto` and `docarray.proto` respectively. 

- `client.go` defines the `Client` interface. This is implemented by `grpc.go`, `http.go` and `websocket.go`.


#### Jina version 

TODO

