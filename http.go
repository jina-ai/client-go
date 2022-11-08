package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/jina-ai/client-go/jina"
)

type HTTPClient struct {
	Host string
	ctx  context.Context
}

func NewHTTPClient(host string) (*HTTPClient, error) {
	if !strings.HasSuffix(host, "/post") {
		host = host + "/post"
	}
	return &HTTPClient{
		Host: host,
		ctx:  context.Background(),
	}, nil
}

func (c HTTPClient) POST(requests <-chan *jina.DataRequestProto, onDone, onError, onAlways CallbackType) error {
	var wg sync.WaitGroup

	handleRequest := func(request *jina.DataRequestProto) {
		reqJSON, err := json.Marshal(request)
		if err != nil {
			fmt.Println("error marshalling request", err)
			if onError != nil {
				onError(request)
			}
			if onAlways != nil {
				onAlways(request)
			}
		}

		req, err := http.NewRequest("POST", c.Host, bytes.NewBuffer(reqJSON))
		if err != nil {
			fmt.Println("error creating request", err)
		}
		req.Header.Set("Content-Type", "application/json")

		httpResp, err := httpClient.Do(req)
		if err != nil {
			fmt.Println("error sending request", err)
			if onError != nil {
				onError(request)
			}
			if onAlways != nil {
				onAlways(request)
			}
		}

		if httpResp != nil {
			defer httpResp.Body.Close()

			if httpResp.StatusCode != http.StatusOK {
				fmt.Println("Got non 200 status code", httpResp.StatusCode)
				if onError != nil {
					onError(request)
				}
				if onAlways != nil {
					onAlways(request)
				}
			}

			body, err := io.ReadAll(httpResp.Body)
			if err != nil {
				fmt.Println("error reading response body", err)
				if onError != nil {
					onError(request)
				}
				if onAlways != nil {
					onAlways(request)
				}
			}

			var res jina.DataRequestProto
			if err := json.Unmarshal(body, &res); err != nil {
				fmt.Println("error unmarshalling response", err)
				if onError != nil {
					onError(request)
				}
				if onAlways != nil {
					onAlways(request)
				}
			} else {
				if onDone != nil {
					onDone(&res)
				}
				if onAlways != nil {
					onAlways(&res)
				}
			}
		}
		wg.Done()
	}

	for {
		req, ok := <-requests
		if !ok {
			break
		}
		go handleRequest(req)
		wg.Add(1)
	}
	wg.Wait()
	return nil
}

type HTTPHealthCheckClient struct {
	Host string
	ctx  context.Context
}

func NewHTTPHealthCheckClient(host string) (*HTTPHealthCheckClient, error) {
	if !strings.HasPrefix(host, "http") {
		host = "http://" + host
	}

	if !strings.HasSuffix(host, "/dry_run") {
		host = host + "/dry_run"
	}
	return &HTTPHealthCheckClient{
		Host: host,
		ctx:  context.Background(),
	}, nil
}

func (c HTTPHealthCheckClient) HealthCheck() (bool, error) {
	req, err := http.NewRequest("GET", c.Host, nil)
	if err != nil {
		fmt.Println("Failed to create HTTP request", "host", c.Host, "err", err)
		return false, err
	}

	httpResp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}
	if httpResp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("got non 200 status code %d", httpResp.StatusCode)
	}
	defer httpResp.Body.Close()

	var resp map[string]interface{}
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return false, fmt.Errorf("error decoding response: %w", err)
	}

	code := int(resp["code"].(float64))
	if code != 0 {
		return false, fmt.Errorf("got non 0 code %s", resp["description"])
	} else {
		return true, nil
	}
}

type HTTPInfoClient struct {
	Host string
	ctx  context.Context
}

func NewHTTPInfoClient(host string) (HTTPInfoClient, error) {
	if !strings.HasPrefix(host, "http") {
		host = "http://" + host
	}

	if !strings.HasSuffix(host, "/status") {
		host = host + "/status"
	}
	return HTTPInfoClient{
		Host: host,
		ctx:  context.Background(),
	}, nil
}

func (c HTTPInfoClient) InfoJSON() ([]byte, error) {
	req, err := http.NewRequest("GET", c.Host, nil)
	if err != nil {
		fmt.Println("Failed to create HTTP request", "host", c.Host, "err", err)
		return nil, err
	}

	httpResp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got non 200 status code %d", httpResp.StatusCode)
	}
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, body, "", "  "); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c HTTPInfoClient) Info() (*jina.JinaInfoProto, error) {
	body, err := c.InfoJSON()
	if err != nil {
		return nil, err
	}

	var res jina.JinaInfoProto
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
