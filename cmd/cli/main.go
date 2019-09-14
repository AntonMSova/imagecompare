package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AntonMSova/imagecompare/pkg/processor"
)

// main sets up the flags and runs image processor
func main() {
	inputFilename := flag.String("f", "examples/input/images.csv", "a csv file in a format of 'aa.png,bb.png'")
	outputFilename := flag.String("o", "examples/output/result.csv", "a csv output file name")
	flag.Parse()

	fmt.Printf("Comparing images from '%s'...\n", *inputFilename)
	err := processor.Run(*inputFilename, *outputFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}