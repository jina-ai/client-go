package client

import (
	"github.com/deepankarm/client-go/jina"
)

type CallbackType func(*jina.DataRequestProto)

type Client interface {
	POST(requests <-chan *jina.DataRequestProto, onDone, onError, onAlways CallbackType) error
}
