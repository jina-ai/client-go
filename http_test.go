package client

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTTP Client", Ordered, func() {
	var c *HTTPClient
	var hc *HTTPHealthCheckClient
	var err error

	BeforeEach(func() {
		cleanUp := startFlow("examples/http/flow.yml")
		time.Sleep(2 * time.Second)
		DeferCleanup(func() {
			*c = HTTPClient{}
			cleanUp()
		})
	})

	Describe("Create the Client and stream requests", func() {
		It("should create a new HTTPClient & stream requests", func() {
			Eventually(func() error {
				c, err = NewHTTPClient("http://localhost:12345")
				return err
			}, 10*time.Second, 1*time.Second).Should(BeNil())
			Expect(c).NotTo(BeNil())
			c.POST(generateDataRequests(3), OnDone, OnError, nil)
		})
	})

	Describe("Perform healthchecks on the client", func() {
		It("should create a new HTTPHealthCheckClient & perform a successful healthcheck", func() {
			Eventually(func() error {
				hc, err = NewHTTPHealthCheckClient("http://localhost:12345")
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
