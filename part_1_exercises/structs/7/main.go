package main

import "fmt"

type Book struct {
	Title   string
	Author  string
	Reviews map[string]string
}

func main() {
	book := Book{
		Title:  "Advanced Go",
		Author: "Jane Smith",
		Reviews: map[string]string{
			"John":  "Great book!",
			"Alice": "Very informative.",
		},
	}
	fmt.Println(book)
}
