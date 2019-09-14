package processor_test

import (
	"github.com/AntonMSova/imagecompare/pkg/csvparser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"image"
	"image/color"
	"strconv"

	. "github.com/AntonMSova/imagecompare/pkg/processor"
)

type customColor struct {
	a uint32
	b uint32
	g uint32
	r uint32
}

func (c customColor) RGBA() (r, g, b, a uint32) {
	return c.r, c.g, c.b, c.a
}

type rgba struct {
	pix []uint8
	stride int
	rect image.Rectangle
	color customColor
}

func (p rgba) ColorModel() color.Model { return color.RGBAModel }
func (p rgba) Bounds() image.Rectangle { return p.rect }
func (p rgba) At(x, y int) color.Color { return p.color }

var _ = Describe("Service", func() {

	Describe("Run", func() {
		Context("when input csv exists", func() {
			It("creates an output csv file", func() {
				input := "../../examples/input/fixtures/images_test.csv"
				output := "../../examples/output/tests/result_test.csv"
				err := Run(input, output)
				Expect(err).NotTo(HaveOccurred())
				Expect(input).To(BeAnExistingFile())
				Expect(output).To(BeAnExistingFile())
				Expect(input).NotTo(BeEmpty())
				Expect(output).NotTo(BeEmpty())
			})
		})

		Context("when an error occurs", func() {
			It("returns does not create an output file", func() {
				input := "../../examples/input/fixtures/images_test2.csv"
				output := "../../examples/output/tests/result_test2.csv"
				err := Run(input, output)
				Expect(err).To(HaveOccurred())
				Expect(input).To(BeAnExistingFile())
				Expect(input).NotTo(BeEmpty())
				Expect(output).NotTo(BeAnExistingFile())
			})
		})
	})

	Describe("GenerateReports", func() {
		Context("when an image doesn't exist", func() {
			It("returns an error", func() {
				img1 := "../../examples/input/images/test.png"
				img2 := "qwerrt"
				images := []csvparser.Images{
					{
						Img1: img1,
						Img2: img2,
					},
				}

				reports, err := GenerateReports(images)
				Expect(err).To(HaveOccurred())
				Expect(img1).To(BeAnExistingFile())
				Expect(img1).NotTo(BeEmpty())
				Expect(img2).NotTo(BeAnExistingFile())
				Expect(reports).To(HaveLen(0))
			})
		})

		Context("when both images exist but have different types", func() {
			It("returns a ratio of 1", func() {
				img1 := "../../examples/input/images/test.png"
				img2 := "../../examples/input/images/test6.jpeg"
				images := []csvparser.Images{
					{
						Img1: img1,
						Img2: img2,
					},
				}

				reports, err := GenerateReports(images)
				Expect(err).NotTo(HaveOccurred())
				Expect(img1).To(BeAnExistingFile())
				Expect(img1).NotTo(BeEmpty())
				Expect(img2).To(BeAnExistingFile())
				Expect(img2).NotTo(BeEmpty())
				Expect(reports).To(HaveLen(1))
				Expect(reports[0]).To(HaveLen(4))

				f, err := strconv.ParseFloat(reports[0][2], 64)
				Expect(err).NotTo(HaveOccurred())
				Expect(f).Should(BeNumerically("==", 1))

				f, err = strconv.ParseFloat(reports[0][3], 64)
				Expect(err).NotTo(HaveOccurred())
				Expect(f).Should(BeNumerically("==", 0))
			})
		})

		Context("when both images exist but have different types", func() {
			It("returns a ratio between 0 and 1", func() {
				img1 := "../../examples/input/images/test.png"
				img2 := "../../examples/input/images/test3.png"
				images := []csvparser.Images{
					{
						Img1: img1,
						Img2: img2,
					},
				}

				reports, err := GenerateReports(images)
				Expect(err).NotTo(HaveOccurred())
				Expect(img1).To(BeAnExistingFile())
				Expect(img1).NotTo(BeEmpty())
				Expect(img2).To(BeAnExistingFile())
				Expect(img2).NotTo(BeEmpty())
				Expect(reports).To(HaveLen(1))
				Expect(reports[0]).To(HaveLen(4))

				f, err := strconv.ParseFloat(reports[0][2], 64)
				Expect(err).NotTo(HaveOccurred())
				Expect(f).Should(BeNumerically(">", 0))
				Expect(f).Should(BeNumerically("<", 1))

				f, err = strconv.ParseFloat(reports[0][3], 64)
				Expect(err).NotTo(HaveOccurred())
				Expect(f).Should(BeNumerically(">", 0))
			})
		})
	})

	Describe("Load", func() {
		Context("when a file doesn't exist", func() {
			It("returns an error", func() {
				f := "../../examples/input/images/doesnt_exist.txt"
				_, _, err := Load(f)
				Expect(err).To(HaveOccurred())
				Expect(f).NotTo(BeAnExistingFile())
			})
		})

		Context("when a file exists but has a wrong format", func() {
			It("returns an error", func() {
				f := "../../examples/input/images/test5.txt"
				_, _, err := Load(f)
				Expect(err).To(HaveOccurred())
				Expect(f).To(BeAnExistingFile())
			})
		})

		Context("when a file exists but has a correct format", func() {
			It("returns an image and image type", func() {
				f := "../../examples/input/images/test.png"
				_, imgType, err := Load(f)
				Expect(err).NotTo(HaveOccurred())
				Expect(f).To(BeAnExistingFile())
				Expect(imgType).To(Equal("png"))
			})
		})
	})

	Describe("CompareImages", func() {
		a := image.Rectangle{
			Min: image.Point{
				X: 100,
				Y: 100,
			},
			Max: image.Point{
				X: 200,
				Y: 200,
			},
		}

		Context("when two images have different bounds", func() {
			It("returns 1", func() {
				b := image.Rectangle{
					Min: image.Point{
						X: 110,
						Y: 100,
					},
					Max: image.Point{
						X: 200,
						Y: 200,
					},
				}

				p1 := rgba{ rect: a }
				p2 := rgba{ rect: b }

				dif := CompareImages(p1, p2)
				Expect(dif).Should(BeNumerically("==", 1))
			})
		})

		Context("when two images have the same bounds and same colors", func() {
			It("returns 0", func() {
				b := image.Rectangle{
					Min: image.Point{
						X: 100,
						Y: 100,
					},
					Max: image.Point{
						X: 200,
						Y: 200,
					},
				}

				c1 := customColor{10, 20, 30, 40}
				c2 := customColor{10, 20, 30, 40}

				p1 := rgba{ rect: a, color: c1 }
				p2 := rgba{ rect: b, color: c2 }

				dif := CompareImages(p1, p2)
				Expect(dif).To(BeZero())
			})
		})

		Context("when two images have the same bounds but different colors", func() {
			It("calculates difference", func() {
				b := image.Rectangle{
					Min: image.Point{
						X: 100,
						Y: 100,
					},
					Max: image.Point{
						X: 200,
						Y: 200,
					},
				}

				c1 := customColor{10, 20, 30, 40}
				c2 := customColor{100, 200, 300, 400}

				p1 := rgba{ rect: a, color: c1 }
				p2 := rgba{ rect: b, color: c2 }

				dif := CompareImages(p1, p2)
				Expect(dif).Should(BeNumerically(">", 0))
				Expect(dif).Should(BeNumerically("<", 1))
			})
		})
	})

	Describe("CompareColor", func() {
		Context("given tow colors", func() {
			It("correctly calculates the difference between them", func() {
				c1 := customColor{10, 20, 30, 40}
				c2 := customColor{15, 30, 50, 100}

				dif := CompareColor(c1, c2)
				Expect(dif).To(Equal(int64(90)))
			})
		})
	})

	Describe("Diff", func() {
		a := uint32(5)
		b := uint32(15)

		Context("when 'a' is passed as a first parameter", func() {
			It("the difference should be positive", func() {
				dif := Differ(a, b)
				Expect(dif).To(Equal(int64(10)))
			})
		})

		Context("when 'a' is passed as a second parameter", func() {
			It("the difference should be positive", func() {
				dif := Differ(b, a)
				Expect(dif).To(Equal(int64(10)))
			})
		})
	})

	Describe("BoundsMatch", func() {
		a := image.Rectangle{
			Min: image.Point{
				X: 100,
				Y: 100,
			},
			Max: image.Point{
				X: 200,
				Y: 200,
			},
		}

		Context("when two rectangles are different ", func() {
			It("should return false", func() {
				b := image.Rectangle{
					Min: image.Point{
						X: 110,
						Y: 100,
					},
					Max: image.Point{
						X: 200,
						Y: 200,
					},
				}

				matched := BoundsMatch(a, b)
				Expect(matched).To(BeFalse())
			})
		})

		Context("when two rectangles are the same", func() {
			It("should return true", func() {
				b := image.Rectangle{
					Min: image.Point{
						X: 100,
						Y: 100,
					},
					Max: image.Point{
						X: 200,
						Y: 200,
					},
				}

				matched := BoundsMatch(a, b)
				Expect(matched).To(BeTrue())
			})
		})
	})
})
