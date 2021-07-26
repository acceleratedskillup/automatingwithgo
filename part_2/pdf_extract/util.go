package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func extractPages(inputFilename, pages, outputDir string) {

	pageSelection, err := processPageSelection(pages)
	if err != nil {
		fmt.Printf("Error processing page selection: %v\n", err)
		return
	}

	// Check if outputDir exists, if not, create it
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			return
		}
	}

	// Open the input PDF file
	file, err := os.Open(inputFilename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	// Use the input filename as a prefix for the output files
	prefix := inputFilename[:len(inputFilename)-pdfExtensionLength] // remove ".pdf" extension

	conf := model.NewDefaultConfiguration()
	if err := api.ExtractPages(file, outputDir, prefix, pageSelection, conf); err != nil {
		fmt.Println("Error extracting pages:", err)
		return
	}

	fmt.Println("Pages extracted successfully!")
}

func processPageSelection(pages string) ([]string, error) {
	var selection []string
	parts := strings.Split(pages, ",")

	for _, part := range parts {
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return nil, fmt.Errorf("invalid start page in range %s: %v", part, err)
			}
			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid end page in range %s: %v", part, err)
			}

			for i := start; i <= end; i++ {
				selection = append(selection, strconv.Itoa(i))
			}
		} else {
			_, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid page number %s: %v", part, err)
			}
			selection = append(selection, part)
		}
	}

	return selection, nil
}
