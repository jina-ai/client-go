## gRPC Client

#### Start a Flow with gRPC protocol

```bash
jina flow --uses flow.yml
```

#### Send `DataRequest`s to the Flow

```bash
go run main.go
```
```text
Total 3 docs received.
DocID: 0
        Chunk 0 text: Hello world. 
        Chunk 1 text: This is a test document with id:0 
DocID: 1
        Chunk 0 text: Hello world. 
        Chunk 1 text: This is a test document with id:1 
DocID: 2
        Chunk 0 text: Hello world. 
        Chunk 1 text: This is a test document with id:2
```
