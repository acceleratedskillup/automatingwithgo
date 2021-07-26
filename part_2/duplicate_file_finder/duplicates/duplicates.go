package duplicates

import (
	"bufio"
	"dupfilefinder/filehash"
	"dupfilefinder/progressupdate"
	"fmt"
	"os"
	"sort"
	"sync"
	"time"
)

const duplicatesFileName = "duplicates.txt"

type DuplicatesMapKey struct {
	Size int64
	Hash [32]byte
}

func UpdateDuplicatesMap(updateDuplicatesMapWG *sync.WaitGroup,
	updateDuplicatesChannel chan filehash.FileHash,
	duplicatesMap map[DuplicatesMapKey][]string,
	startTime time.Time) {
	defer updateDuplicatesMapWG.Done()
	fileCount := 0
	for fileHash := range updateDuplicatesChannel {
		fileCount++
		progressupdate.PrintProgress(fileCount, startTime)

		key := DuplicatesMapKey{Size: fileHash.Size, Hash: fileHash.Hash}
		duplicatesMap[key] = append(duplicatesMap[key], fileHash.Path)
	}
}

func CheckForDuplicates(duplicatesMap map[DuplicatesMapKey][]string) error {
	// Create or open the file
	file, err := os.Create(duplicatesFileName)
	if err != nil {
		return err
	}
	defer file.Close()
	// Create a buffered writer from the file
	writer := bufio.NewWriter(file)
	sortedKeys := getSortedByFileSize(duplicatesMap)

	duplicateFilesCount := 0
	for _, key := range sortedKeys {
		paths := duplicatesMap[key]
		if len(paths) > 1 {
			duplicateFilesCount += len(paths)

			for _, path := range paths {
				err := writeLine(writer, path+" size:"+fmt.Sprintf("%d", key.Size)+" bytes")
				if err != nil {
					return err
				}
			}
			err := writeLine(writer, "")
			if err != nil {
				return err
			}
		}
	}
	// Flush the buffered writer
	err = writer.Flush()
	if err != nil {
		return err
	}

	fmt.Println("\nnumber of duplicate files:", duplicateFilesCount)
	if duplicateFilesCount > 0 {
		fmt.Println("details in:", duplicatesFileName)
	}

	return nil
}

// Function to write a single line to the file
func writeLine(writer *bufio.Writer, line string) error {
	_, err := writer.WriteString(line + "\n")
	if err != nil {
		return err
	}
	return nil
}

func getSortedByFileSize(duplicatesTracker map[DuplicatesMapKey][]string) []DuplicatesMapKey {
	keys := make([]DuplicatesMapKey, 0, len(duplicatesTracker))
	for key := range duplicatesTracker {
		keys = append(keys, key)
	}

	// Sort the keys slice
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Size > keys[j].Size // Sort in descending order by Size
	})
	return keys

}
