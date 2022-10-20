package client

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jina-ai/client-go/docarray"
	"github.com/jina-ai/client-go/jina"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shirou/gopsutil/v3/process"
)

func TestClients(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t, "Client Suite", []Reporter{})
}

func execCommand(name string, arg ...string) func() {
	cmd := exec.Command(name, arg...)
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command", err)
	}
	return func() {
		jina := "jina"
		processes, _ := process.Pids()
		for _, element := range processes {
			pro, _ := process.NewProcess(element)
			pro_name, _ := pro.Name()
			if pro_name == jina {
				pro.Kill()
			}
		}
	}
}

func startFlow(path string) func() {
	return execCommand("jina", "flow", "--uses", filepath.Join(curDir(), path))
}

func curDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

// Create a Document
func getDoc(id string) *docarray.DocumentProto {
	return &docarray.DocumentProto{
		Id: id,
		Content: &docarray.DocumentProto_Text{
			Text: "Hello world. This is a test document with id:" + id,
		},
	}
}

// Create a DocumentArray with 3 Documents
func getDocarrays(numDocs int) *docarray.DocumentArrayProto {
	var docs []*docarray.DocumentProto
	for i := 0; i < numDocs; i++ {
		docs = append(docs, getDoc(fmt.Sprint(i)))
	}
	return &docarray.DocumentArrayProto{
		Docs: docs,
	}
}

// Create DataRequest with a DocumentArray
func getDataRequest(numDocs int) *jina.DataRequestProto {
	return &jina.DataRequestProto{
		Data: &jina.DataRequestProto_DataContentProto{
			Documents: &jina.DataRequestProto_DataContentProto_Docs{
				Docs: getDocarrays(numDocs),
			},
		},
	}
}

// Generate a stream of requests
func generateDataRequests(numRequests int) <-chan *jina.DataRequestProto {
	defer GinkgoRecover()
	requests := make(chan *jina.DataRequestProto)
	go func() {
		for i := 0; i < numRequests; i++ {
			requests <- getDataRequest(3)
		}
		defer close(requests)
	}()
	return requests
}

func OnDone(resp *jina.DataRequestProto) {
	defer GinkgoRecover()
	switch docs := resp.Data.Documents.(type) {
	case *jina.DataRequestProto_DataContentProto_Docs:
		By("should have the correct number of documents")
		Expect(len(docs.Docs.Docs)).To(Equal(3))

		for docidx, doc := range docs.Docs.Docs {
			By(fmt.Sprintf("should have the correct id for document %d", docidx))
			Expect(doc.Id).To(Equal(fmt.Sprint(docidx)))

			By(fmt.Sprintf("should have the correct text for document %d", docidx))
			Expect(doc.Chunks).To(HaveLen(2))
			Expect(doc.Chunks[0].Content.(*docarray.DocumentProto_Text).Text).To(Equal("Hello world."))
			Expect(doc.Chunks[1].Content.(*docarray.DocumentProto_Text).Text).To(Equal("This is a test document with id:" + doc.Id))

		}
	}
}

func OnError(resp *jina.DataRequestProto) {
	defer GinkgoRecover()
	fmt.Println("Got an error for request", resp)
}
