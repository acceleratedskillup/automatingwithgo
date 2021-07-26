package main

import (
	"fmt"
	"os"
)

const (
	expectedNumOfArgs   = 4
	inputFilenameArgPos = 1
	pagesArgPos         = 2
	outputDirArgPos     = 3
	pdfExtensionLength  = len(".pdf")
)

func main() {
	if len(os.Args) != expectedNumOfArgs {
		fmt.Println("Usage: go run . <input_filename> <pages> <output_directory>")
		fmt.Println("Pages format: 1-4 or 1,3,5 or a combination like 1-3,5,7-9")
		return
	}

	inputFilename := os.Args[inputFilenameArgPos]
	pages := os.Args[pagesArgPos]
	outputDir := os.Args[outputDirArgPos]

	extractPages(inputFilename, pages, outputDir)
}
