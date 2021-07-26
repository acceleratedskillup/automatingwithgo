package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

const (
	pdfDir         = "./pdfs"
	outputFilename = "merged_output.pdf"
)

func main() {
	inputFilenames := getPdfFilesFromDirectory(pdfDir)

	if len(inputFilenames) == 0 {
		fmt.Println("No PDF files found in the directory.")
		return
	}

	mergePDFs(outputFilename, inputFilenames)
}

func getPdfFilesFromDirectory(directory string) []string {
	entries, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	var pdfFiles []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".pdf") {
			pdfFiles = append(pdfFiles, directory+"/"+entry.Name())
		}
	}

	return pdfFiles
}

func mergePdfs(outputFilename string, inputFilenames []string) {
	config := model.NewDefaultConfiguration()
	if err := api.MergeCreateFile(inputFilenames, outputFilename, config); err != nil {
		fmt.Println("Error merge creating file:", err)
		return
	}

	if err := api.ValidateFile(outputFilename, config); err != nil {
		fmt.Println("Error validating file:", err)
		return
	}

	fmt.Println("PDFs merged successfully!")
}
