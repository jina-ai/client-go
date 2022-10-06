## Get Info for gRPC Flow

### Start a Flow with gRPC protocol

```bash
jina flow --uses flow.yml
```

### Fetch Info

```bash
go run main.go --host grpc://localhost:12345
```
```json
{
  "jina": {
    "jina": "3.10.0",
    "docarray": "0.16.5",
    "jcloud": "0.0.37",
    "jina-hubble-sdk": "0.18.0",
    "jina-proto": "0.1.13",
    "protobuf": "4.21.6",
    "proto-backend": "upb",
    "grpcio": "1.46.0",
    "pyyaml": "5.4.1",
    "python": "3.7.2",
    "platform": "Linux",
    "platform-release": "5.15.0-48-generic",
    "platform-version": "#54~20.04.1-Ubuntu SMP Thu Sep 1 16:17:26 UTC 2022",
    "architecture": "x86_64",
    "processor": "x86_64",
    "uid": "153855726454644",
    "session-id": "734b24c4-4560-11ed-973f-8bee53ec0774",
    "uptime": "2022-10-06T15:47:55.642813",
    "ci-vendor": "(unset)",
    "internal": "False"
  },
  "envs": {
    "JINA_DEFAULT_HOST": "(unset)",
    "JINA_DEFAULT_TIMEOUT_CTRL": "(unset)",
    "JINA_DEPLOYMENT_NAME": "(unset)",
    "JINA_DISABLE_UVLOOP": "(unset)",
    "JINA_EARLY_STOP": "(unset)",
    "JINA_FULL_CLI": "(unset)",
    "JINA_GATEWAY_IMAGE": "(unset)",
    "JINA_GRPC_RECV_BYTES": "(unset)",
    "JINA_GRPC_SEND_BYTES": "(unset)",
    "JINA_HUB_NO_IMAGE_REBUILD": "(unset)",
    "JINA_LOG_CONFIG": "(unset)",
    "JINA_LOG_LEVEL": "(unset)",
    "JINA_LOG_NO_COLOR": "(unset)",
    "JINA_MP_START_METHOD": "(unset)",
    "JINA_OPTOUT_TELEMETRY": "(unset)",
    "JINA_RANDOM_PORT_MAX": "(unset)",
    "JINA_RANDOM_PORT_MIN": "(unset)"
  }
}
```
