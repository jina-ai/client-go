package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jina-ai/client-go/jina"
)

type WebSocketClient struct {
	Host string
	conn *websocket.Conn
	ctx  context.Context
}

func NewWebSocketClient(host string) (*WebSocketClient, error) {
	var u *url.URL
	u, err := url.Parse(host)
	if err != nil {
		u = &url.URL{Scheme: "ws", Host: host, Path: "/ws"}
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return &WebSocketClient{}, err
	}
	client := &WebSocketClient{
		Host: host,
		conn: conn,
		ctx:  context.Background(),
	}
	return client, nil
}

func (client WebSocketClient) POST(requests <-chan *jina.DataRequestProto, onDone, onError, onAlways CallbackType) error {
	var wg sync.WaitGroup

	handleRequest := func(request *jina.DataRequestProto) {
		reqJSON, err := json.Marshal(request)
		if err != nil {
			if onError != nil {
				onError(request)
			}
			if onAlways != nil {
				onAlways(request)
			}
		}

		err = client.conn.WriteMessage(websocket.TextMessage, reqJSON)
		if err != nil {
			if onError != nil {
				onError(request)
			}
			if onAlways != nil {
				onAlways(request)
			}
		}
	}
	go func() {
		for {
			_, data, err := client.conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
			}
			var res jina.DataRequestProto
			if err := json.Unmarshal(data, &res); err != nil {
				fmt.Println(err)
			}

			if onDone != nil {
				onDone(&res)
			}
			if onAlways != nil {
				onAlways(&res)
			}
			wg.Done()
		}
	}()

	for {
		req, ok := <-requests
		if !ok {
			break
		}
		handleRequest(req)
		wg.Add(1)
	}
	wg.Wait()
	return nil
}
