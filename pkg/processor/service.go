package processor

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"time"

	"github.com/AntonMSova/imagecompare/pkg/csvparser"
)

// Run reads input file, compares images and generates csv containing
// pairs of original images, their difference ratio and process time
func Run(inputFileName, outputFileName string) error {
	// get a slice of image pairs from csv
	allImages, err := csvparser.Read(inputFileName)
	if err != nil {
		return err
	}

	// go over every pair and Load images of each pair
	reports, err := GenerateReports(allImages)
	if err != nil {
		return err
	}

	// Save result in a csv file
	return csvparser.Write(outputFileName, reports)
}

// GenerateReports accepts a slice of images and returns a slice of reports
func GenerateReports(images []csvparser.Images) ([]csvparser.Report, error) {
	var reports []csvparser.Report
	// go over every pair and Load images of each pair
	for _, i := range images {
		img1, imgType1, err := Load(i.Img1)
		if err != nil {
			return nil, err
		}

		img2, imgType2, err := Load(i.Img2)
		if err != nil {
			return nil, err
		}

		// if both images have different types, set ratio to 1 and skip comparing
		if imgType1 != imgType2 {
			r := csvparser.Report{i.Img1, i.Img2, "1", "0.00"}
			reports = append(reports, r)
			continue
		}

		// start timer
		start := time.Now()
		// run comparison algorithm
		diff := CompareImages(img1, img2)
		// stop timer
		elapsed := fmt.Sprintf("%.2f", time.Since(start).Seconds())

		// if difference ratio is 0.00 or 1.00, remove trailing zeros
		d := fmt.Sprintf("%.2f", diff)
		if diff == 0.0 || diff == 1.0 {
			d = fmt.Sprintf("%.0f", diff)
		}

		r := csvparser.Report{i.Img1, i.Img2, d, elapsed}
		reports = append(reports, r)
	}

	return reports, nil
}

// Load opens file and decodes it
func Load(filePath string) (image.Image, string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	// Decode opened file
	img, imgType, err := image.Decode(f)
	if err != nil {
		if err == image.ErrFormat {
			return nil, "", fmt.Errorf("file '%s' has wrong format", f.Name())
		}

		return nil,"", err
	}

	return img, imgType, nil
}

// CompareImages calculates difference ration between two images
func CompareImages(img1, img2 image.Image) float64 {
	var diff int64

	b1 := img1.Bounds()
	b2 := img2.Bounds()

	if !BoundsMatch(b1, b2) {
		// Image sizes don't match
		return 1
	}

	for y := b1.Min.Y; y < b1.Max.Y; y++ {
		for x := b1.Min.X; x < b1.Max.X; x++ {
			diff += CompareColor(img1.At(x, y), img2.At(x, y))
		}
	}

	nPixels := (b1.Max.X - b1.Min.X) * (b1.Max.Y - b1.Min.Y)

	return float64(diff) / (float64(nPixels) * 0xffff * 3)
}

// CompareColor returns difference between RGBA colors
func CompareColor(c1, c2 color.Color) (diff int64) {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	diff += Differ(r1, r2)
	diff += Differ(g1, g2)
	diff += Differ(b1, b2)

	return diff
}

// Differ returns difference between two color values
func Differ(a, b uint32) int64 {
	if a > b {
		return int64(a - b)
	}
	return int64(b - a)
}

// BoundsMatch checks if bounds are different
func BoundsMatch(a, b image.Rectangle) bool {
	return a.Min.X == b.Min.X && a.Min.Y == b.Min.Y && a.Max.X == b.Max.X && a.Max.Y == b.Max.Y
}
