package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AntonMSova/imagecompare/pkg/processor"
)

// main sets up the flags and runs image processor
func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <input file in a format of 'aa.png,bb.png'>\n\nOptions:\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	outputFilename := flag.String("o", "result.csv", "a csv output file name")
	flag.Parse()

	// Make sure we have input paths.
	if flag.NArg() == 0 {
		fmt.Print("Error: missing <input file>\n\n")
		flag.Usage()

		os.Exit(1)
	}

	inputFilename := flag.Arg(0)

	fmt.Printf("Comparing images from '%s'...\n\n", inputFilename)
	err := processor.Run(inputFilename, *outputFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("All done.\nTHe output file is: '%s'\n", *outputFilename)
}