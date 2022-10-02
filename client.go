package client

import (
	"github.com/deepankarm/client-go/jina"
)

type CallbackType func(*jina.DataRequestProto)
