package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) Describe() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d\n", b.Title, b.Author, b.Pages)
}

func main() {
	book := Book{"Go Patterns", "Emily", 250}
	book.Describe()
}
