package main

import "fmt"

func main() {
	inventory := map[string]int{
		"The Great Gatsby": 5,
		"Moby Dick":        3,
	}

	inventory["The Great Gatsby"] = 10
	if inventory["The Great Gatsby"] == 10 {
		fmt.Println("Inventory updated successfully!")
	} else {
		fmt.Println("Failed to update inventory!")
	}
}
