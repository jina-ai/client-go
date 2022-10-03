## Healthcheck for gRPC Flow

### Start a Flow with gRPC protocol

```bash
jina flow --uses flow.yml
```

### Check if the Flow is healthy

##### Healthy Flow
```bash
go run main.go --host localhost:12345
```
```text
Flow running at localhost:12345 is healthy!
```

##### Unhealthy Flow
```bash
go run main.go --host localhost:12346
```
```text
panic: failed to check health: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing dial tcp 127.0.0.1:12345: connect: connection refused"
```
