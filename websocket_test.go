package client

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Websocket Client", Ordered, func() {
	var c *WebSocketClient
	var hc *WebSocketHealthCheckClient
	var err error

	BeforeEach(func() {
		flowFunction := startFlow("examples/websocket/flow.yml")
		time.Sleep(2 * time.Second)
		DeferCleanup(func() {
			*c = WebSocketClient{}
			flowFunction()
		})
	})

	Describe("Create the Client and stream requests", func() {
		It("should create a new WebSocketClient & stream requests", func() {
			Eventually(func() error {
				c, err = NewWebSocketClient("ws://localhost:12345")
				return err
			}, 10*time.Second, 1*time.Second).Should(BeNil())
			Expect(c).NotTo(BeNil())
			c.POST(generateDataRequests(3), OnDone, OnError, nil)
		})
	})

	Describe("Perform healthchecks on the client", func() {
		It("should create a new WebSocketHealthCheckClient & perform a successful healthcheck", func() {
			Eventually(func() error {
				hc, err = NewWebSocketHealthCheckClient("ws://localhost:12345")
				return err
			}, 10*time.Second, 1*time.Second).Should(BeNil())
			Expect(hc).NotTo(BeNil())

			Eventually(func() bool {
				status, _ := hc.HealthCheck()
				return status
			}, 10*time.Second, 1*time.Second).Should(BeTrue())
		})
	})
})
