package main

// An example demonstrating an application with multiple views.
//
// Note that this example was produced before the Bubbles progress component
// was available (github.com/charmbracelet/bubbles/progress) and thus, we're
// implementing a progress bar from scratch here.

import (
	"goprojgen/internal/cli"
	"log"
	"os"
)

const logFilePath = "goprojgen.log"

var cleanup func() error

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

func main() {

	defer func() {
		if err := cleanup(); err != nil {
			log.Printf("Error during cleanup: %v\n", err)
			os.Exit(1)
		}
	}()

	cli.Start()
}
