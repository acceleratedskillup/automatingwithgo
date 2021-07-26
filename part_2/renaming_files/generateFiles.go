package main

import (
	"fmt"
	"os"
	"time"
)

func createDummyLogFile(filename string, day int) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write dummy log data
	for i := 0; i < 100; i++ {
		file.WriteString(fmt.Sprintf("Log entry %d for %s\n", i, filename))
	}

	// Set a different timestamp for each log file
	timestamp := time.Date(1950, time.September, day, 15, 30, 0, 0, time.UTC)
	//timestamp := time.Now().AddDate(0, 0, -daysAgo)
	fmt.Println(timestamp)
	os.Chtimes(filename, timestamp, timestamp)
}

func main1() {
	// Create dummy log files
	numFiles := 5
	for i := 1; i <= numFiles; i++ {
		createDummyLogFile(fmt.Sprintf("log%d.txt", i), i)
	}
}
