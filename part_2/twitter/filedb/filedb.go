package filedb

import (
	"bufio"
	"fmt"
	"os"
)

// openOrCreateFile opens an existing file or creates a new one if it doesn't exist.
func openOrCreateFile(filePath string) (*os.File, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create file if it does not exist
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		fmt.Println("File created:", filePath)
		return file, nil
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("error opening file %s: %v", filePath, err)
		}
		fmt.Println("File already exists:", filePath)
		return file, nil
	}
}

// IsAlreadyPosted checks if an item has already been posted.
func IsAlreadyPosted(itemID string, filePath string) (bool, error) {
	file, err := openOrCreateFile(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == itemID {
			return true, nil
		}
	}

	return false, scanner.Err()
}

// RecordPostedItem records a posted item.
func RecordPostedItem(itemID string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(itemID + "\n")
	return err
}
