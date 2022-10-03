## Healthcheck for HTTP Flow

### Start a Flow with HTTP protocol

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
panic: failed to check health: Get "http://localhost:12346": dial tcp 127.0.0.1:12346: connect: connection refused
```
