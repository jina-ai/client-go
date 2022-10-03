package client

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GRPC Client", Ordered, func() {
	var c *GRPCClient
	var hc *GRPCHealthCheckClient
	var err error

	BeforeEach(func() {
		flowFunction := startFlow("examples/grpc/flow.yml")
		time.Sleep(2 * time.Second)
		DeferCleanup(func() {
			*c = GRPCClient{}
			flowFunction()
		})
	})

	Describe("Create the Client and stream requests", func() {
		It("should create a new GRPCClient & stream requests", func() {
			Eventually(func() error {
				c, err = NewGRPCClient("grpc://localhost:12345")
				return err
			}, 10*time.Second, 1*time.Second).Should(BeNil())
			Expect(c).NotTo(BeNil())
			c.POST(generateDataRequests(3), OnDone, OnError, nil)
		})
	})

	Describe("Perform healthchecks on the client", func() {
		It("should create a new GRPCHealthCheckClient & perform a successful healthcheck", func() {
			Eventually(func() error {
				hc, err = NewGRPCHealthCheckClient("grpc://localhost:12345")
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
