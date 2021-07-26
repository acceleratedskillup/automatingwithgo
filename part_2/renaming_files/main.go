package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dirName := "." // Current directory

	// Read the directory contents
	files, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Filter and rename files that match the criteria
	for _, file := range files {
		fileName := file.Name()
		if !file.IsDir() && strings.HasPrefix(fileName, "log") && strings.HasSuffix(fileName, ".txt") {
			fileInfo, err := file.Info()
			if err != nil {
				fmt.Println("Error getting FileInfo:", err)
				continue
			}

			fileNameWithoutExtension := strings.TrimSuffix(fileName, ".txt")
			fmt.Println("fileNameWithoutExtension:", fileNameWithoutExtension)
			fmt.Println("currentName:", fileName)
			fmt.Println("Modification Time:", fileInfo.ModTime())
			timestamp := fileInfo.ModTime().Format("20060102_150405")

			// Construct the new filename with the timestamp
			newName := fmt.Sprintf("%s_%s.txt", fileNameWithoutExtension, timestamp)
			fmt.Println("newName:", newName)

			// Rename the file
			err = os.Rename(fileName, newName)
			if err != nil {
				fmt.Printf("Error renaming %s: %v\n", fileName, err)
				continue
			}

			fmt.Printf("Renamed %s to %s\n", fileName, newName)
		}
	}
}
