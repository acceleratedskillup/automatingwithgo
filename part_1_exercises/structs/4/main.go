package main

import "fmt"

func main() {
	rareBook := struct {
		Title          string
		Author         string
		EstimatedValue int
	}{
		Title:          "Rare Book Title",
		Author:         "Unknown",
		EstimatedValue: 10000,
	}
	fmt.Println(rareBook)
}
