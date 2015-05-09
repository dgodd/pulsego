package payload_test

import (
	"io/ioutil"

	"github.com/dgodd/pulsego/payload"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AppVeyorXmlPayload", func() {
	Describe("Full Fixture", func() {
		var result []payload.ProjectStatus
		var isBuilding bool

		BeforeEach(func() {
			xml, err := ioutil.ReadFile("./fixtures/appveyor.xml")
			Expect(err).NotTo(HaveOccurred())
			result, isBuilding, err = payload.AppVeyorXmlPayload(xml)
			Expect(err).NotTo(HaveOccurred())
		})

		It("Has four records", func() {
			Expect(result).To(HaveLen(4))
		})

		It("Returns IsBuilding", func() {
			Expect(isBuilding).To(BeFalse())
		})

		It("Sets build ids", func() {
			Expect(result[0].BuildId).To(Equal(774311))
			Expect(result[1].BuildId).To(Equal(746574))
			Expect(result[2].BuildId).To(Equal(746571))
			Expect(result[3].BuildId).To(Equal(746474))
		})

		It("Sets success", func() {
			Expect(result[0].Success).To(Equal(false))
			Expect(result[1].Success).To(Equal(false))
			Expect(result[2].Success).To(Equal(false))
			Expect(result[3].Success).To(Equal(false))
		})

		It("Sets Published At", func() {
			Expect(result[0].PublishedAt.Unix()).To(Equal(int64(1430607679)))
			Expect(result[1].PublishedAt.Unix()).To(Equal(int64(1429984683)))
			Expect(result[2].PublishedAt.Unix()).To(Equal(int64(1429984609)))
			Expect(result[3].PublishedAt.Unix()).To(Equal(int64(1429981239)))
		})
	})
})
