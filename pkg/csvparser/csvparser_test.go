package csvparser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/AntonMSova/imagecompare/pkg/csvparser"
)

var _ = Describe("Csvparser", func() {

	Describe("Read", func() {
		Context("when a file doesn't exist", func() {
			It("returns an error", func() {
				f := "qwert.txt"
				_, err := Read(f)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when a file less than two fields", func() {
			It("returns an error", func() {
				f := "../../examples/input/fixtures/images_test3.csv"
				_, err := Read(f)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when a file more than two fields", func() {
			It("returns an error", func() {
				f := "../../examples/input/fixtures/images_test4.csv"
				_, err := Read(f)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when a file more two fields", func() {
			It("returns a list of images", func() {
				f := "../../examples/input/fixtures/images_test.csv"
				images, err := Read(f)
				Expect(err).NotTo(HaveOccurred())
				Expect(images).To(HaveLen(6))
			})
		})
	})
})
