package main

import (
	"dupfilefinder/dirwalker"
	dups "dupfilefinder/duplicates"
	"dupfilefinder/filehash" // Replace with your module path
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var cleanup func() error

const logFilePath = "dupfilefinder.log"

func init() {

	// Open or create the log file
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// Set log output to the file
	log.SetOutput(logFile)

	// Cleanup function to close the file
	cleanup = func() error {
		return logFile.Close()
	}
}
func parseFlags() (int, int, int, string) {
	numOfFileHashingGoRoutines := flag.Int("fh", 10, "Specifies number of File Hashing GoRoutines")
	fileHashingChannelBufferSize := flag.Int("fhcb", 100000, "Specifies File Hashing Channel Buffer Size")
	updateDuplicatesChannelBufferSize := flag.Int("udcb", 100000, "Specifies Update Duplicates Channel Buffer Size")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go --fh=<number-of-goroutines> --fhcb=<buffer-size> --udcb=<buffer-size> <starting-directory>")
		os.Exit(1)
	}

	root := flag.Arg(0)
	return *numOfFileHashingGoRoutines, *fileHashingChannelBufferSize, *updateDuplicatesChannelBufferSize, root
}

func main() {
	defer func() {
		if err := cleanup(); err != nil {
			log.Printf("Error during cleanup: %v\n", err)
			os.Exit(1)
		}
	}()

	numOfFileHashingGoRoutines, fileHashingChannelBufferSize, updateDuplicatesChannelBufferSize, root := parseFlags()

	fmt.Println("Root directory is:", root)
	fmt.Println("Number of File Hashing GoRoutines:", numOfFileHashingGoRoutines)
	fmt.Println("File Hashing Channel Buffer Size:", fileHashingChannelBufferSize)
	fmt.Println("Update Duplicates Channel Buffer Size:", updateDuplicatesChannelBufferSize)

	fileHashingChannel := make(chan string, fileHashingChannelBufferSize)
	updateDuplicatesChannel := make(chan filehash.FileHash, updateDuplicatesChannelBufferSize)
	duplicatesMap := make(map[dups.DuplicatesMapKey][]string)

	log.Println("This is a message from main")
	startTime := time.Now()

	var updateDuplicatesMapWG sync.WaitGroup
	updateDuplicatesMapWG.Add(1)
	go dups.UpdateDuplicatesMap(&updateDuplicatesMapWG,
		updateDuplicatesChannel,
		duplicatesMap,
		startTime)

	var fileHashingWG sync.WaitGroup
	for i := 0; i < numOfFileHashingGoRoutines; i++ {
		fileHashingWG.Add(1)
		go filehash.HashFiles(&fileHashingWG, fileHashingChannel, updateDuplicatesChannel)
	}

	err := dirwalker.Walk(root, fileHashingChannel)
	if err != nil {
		log.Println("Error walking directory")
		os.Exit(1)
	}

	log.Println("finished walking")
	close(fileHashingChannel)

	fileHashingWG.Wait()
	close(updateDuplicatesChannel)

	updateDuplicatesMapWG.Wait()
	err = dups.CheckForDuplicates(duplicatesMap)
	if err != nil {
		fmt.Printf("Error checking for duplicates: %v\n", err)
		os.Exit(1)
	}

}
