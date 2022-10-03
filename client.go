package client

import (
	"github.com/jina-ai/client-go/jina"
)

type CallbackType func(*jina.DataRequestProto)

type Client interface {
	POST(requests <-chan *jina.DataRequestProto, onDone, onError, onAlways CallbackType) error
}

type HealthCheckClient interface {
	HealthCheck() (bool, error)
}
