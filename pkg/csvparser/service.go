package csvparser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// Report represents one row in csv file
type Report []string

// Images represents a pair of images
type Images struct {
	Img1 string
	Img2 string
}

// Read parses csv file into a slice of Images
func Read(fileName string) ([]Images, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("filed to open the CSV file %s: %v\n", fileName, err)
	}

	defer file.Close()

	csvr := csv.NewReader(file)

	var allImages []Images
	for {
		row, err := csvr.Read()
		if err != nil {
			// we reached the end of file, time to exit loop
			if err == io.EOF {
				break
			}

			return nil, fmt.Errorf("error occurred while reading CSV file %s: %v\n", fileName, err)
		}

		// Expected file should have two images to compare. If not return a file format error
		if len(row) != 2 {
			return nil, fmt.Errorf("CSV file %s has wrong format. It must be in a format of 'aa.png,bb.png'\n", fileName)
		}

		p := Images{
			Img1: strings.TrimSpace(row[0]),
			Img2: strings.TrimSpace(row[1]),
		}

		allImages = append(allImages, p)
	}

	return allImages, nil
}

// Write creates a csv file with provided data
func Write(fileName string, data []Report) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return err
		}
	}

	return nil
}