package main

import "fmt"

type Publisher struct {
	Name string
}

type Book struct {
	Publisher
	Title  string
	Author string
}

func main() {
	book := Book{
		Publisher: Publisher{Name: "TechBooks"},
		Title:     "Go Best Practices",
		Author:    "Robert",
	}
	fmt.Println(book)
}
