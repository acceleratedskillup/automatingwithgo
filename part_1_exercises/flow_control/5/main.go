package main

import "fmt"

func processFile() {
	fmt.Println("File opened")
	defer fmt.Println("File closed")
	fmt.Println("Processing file")
}

func main() {
	processFile()
}
