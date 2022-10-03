## Healthcheck for gRPC Flow

### Start a Flow with gRPC protocol

```bash
jina flow --uses flow.yml
```

### Check if the Flow is healthy

```bash
go run main.go --host localhost:12345
```
```text
Flow running at localhost:12345 is healthy!
```
