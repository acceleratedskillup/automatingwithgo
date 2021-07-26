package main

import "fmt"

// Global variables to track database status and query message
var databaseStatus string
var queryMessage string

// init function to simulate database connection setup
func init() {
	databaseStatus = "Database connected!"
	queryMessage = "Querying the database..."
}

// Function to display the database status
func displayStatus() {
	fmt.Println(databaseStatus)
}

// Function to simulate a database query
func queryDatabase() {
	fmt.Println(queryMessage)
}

func main() {
	displayStatus()
	queryDatabase()
}
