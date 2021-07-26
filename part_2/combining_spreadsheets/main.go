package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Specify the directory to search for .xlsx files
const dir = "./excel_files" // Current directory

func main() {
	// Create a new Excel file
	targetFile := excelize.NewFile()
	defer func() {
		if err := targetFile.Close(); err != nil {
			fmt.Println("Error closing target file:", err)
		}
	}()

	// Read the directory
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Failed to read directory:", err)
		return
	}

	// Filter for .xlsx files
	var xlsxFiles []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".xlsx") {
			xlsxFiles = append(xlsxFiles, entry.Name())
		}
	}

	// Check if the default sheet exists
	defaultSheet := targetFile.GetSheetName(0)
	firstSheet := true

	for _, srcFile := range xlsxFiles {

		// Determine the sheet name
		sheetName := strings.TrimSuffix(srcFile, ".xlsx")

		// If it's the first sheet and the default sheet exists, rename the default sheet
		if firstSheet && defaultSheet != "" {
			targetFile.SetSheetName(defaultSheet, sheetName)
			firstSheet = false
		} else {
			// Otherwise, create a new sheet
			_, err := targetFile.NewSheet(sheetName)
			if err != nil {
				fmt.Println("Error creating new sheet:", err)
				continue
			}
		}

		err := copyToTargetFile(filepath.Join(dir, srcFile), targetFile, sheetName)
		if err != nil {
			fmt.Printf("Failed to copy from file %s: %v", srcFile, err)
			continue
		}

	}

	// Save the new Excel file
	if err := targetFile.SaveAs("full_inventory.xlsx"); err != nil {
		fmt.Println("Failed to save file:", err)
		return
	}

	fmt.Println("Data successfully copied over to target excel file")
}

// copyToTargetFile copies the content of a source .xlsx file to a target file
func copyToTargetFile(srcPath string, targetFile *excelize.File, sheetName string) error {
	// Open the source Excel file
	srcFile, err := excelize.OpenFile(srcPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := srcFile.Close(); err != nil {
			fmt.Println("Error closing source file:", err)
		}
	}()

	// Get all rows from the source file
	rows, err := srcFile.GetRows(srcFile.GetSheetName(0))
	if err != nil {
		return err
	}

	// Write rows to the target sheet
	for rowIndex, row := range rows {
		for colIndex, cell := range row {
			axis, err := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)
			if err != nil {
				fmt.Println("Error converting coordinates to cell name:", err)
				continue
			}
			targetFile.SetCellValue(sheetName, axis, cell)
		}
	}

	return nil
}
