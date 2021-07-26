package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

// readCSV reads a CSV file and returns its records.
func readCSV(filename string) ([][]string, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// csvToJSON converts CSV records to JSON format.
func csvToJSON(records [][]string) ([]byte, error) {
	var jsonData []map[string]string
	headers := records[0]
	for _, record := range records[1:] {
		rowData := make(map[string]string)
		for i, header := range headers {
			rowData[header] = record[i]
		}
		jsonData = append(jsonData, rowData)
	}

	// Convert to JSON format
	return json.MarshalIndent(jsonData, "", "  ")
}
