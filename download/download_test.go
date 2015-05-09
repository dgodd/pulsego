package download_test

import (
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Download", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()
		// client = sprockets.NewClient(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	// Describe("The sprockets client", func() {
	// 	Describe("fetching sprockets", func() {
	// 		BeforeEach(func() {
	// 			server.AppendHandlers(
	// 				ghttp.VerifyRequest("GET", "https://ci.appveyor.com/api/projects/dgodd/pulsego/history?recordsNumber=10"),
	// 			)
	// 		})

	// 		It("should make a request to fetch sprockets", func() {
	// 			project = payload.Project{
	// 				Type:       "appveyor",
	// 				Repository: "dgodd/pulsogo",
	// 			}
	// 			project.Download()
	// 			Ω(server.ReceivedRequests()).Should(HaveLen(1))
	// 		})
	// 	})
	// })
})
