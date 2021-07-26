package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	sourceFolder      = `.`
	destinationFolder = `logbackups` // You can change this to where you want to save the ZIP files
	zipPrefix         = "serverlogs_"
)

func main() {

	// Check if destinationFolder exists, if not create it
	if _, err := os.Stat(destinationFolder); os.IsNotExist(err) {
		os.MkdirAll(destinationFolder, os.ModePerm)
	}

	// Directories to exclude from zipping
	excludedDirs := []string{destinationFolder} // Add directory names to exclude

	// Find the next available ZIP filename
	nextZipFilename := getNextZipFilename(destinationFolder, zipPrefix)

	// Create the ZIP file
	err := zipLogFiles(sourceFolder, filepath.Join(destinationFolder, nextZipFilename), excludedDirs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Successfully created:", nextZipFilename)
}
