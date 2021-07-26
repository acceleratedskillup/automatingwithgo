package main

import (
	"fmt"
	"os"
)

func main() {
	records, err := readCSV("sales_data.csv")
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	jsonOutput, err := csvToJSON(records)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Write the JSON data to a file
	err = os.WriteFile("output.json", jsonOutput, os.ModePerm)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data written to output.json")
}
