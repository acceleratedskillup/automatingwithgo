package main

import "fmt"

type Genre struct {
	Name string
}

type Book struct {
	Genre
	Title  string
	Author string
}

func main() {
	book := Book{
		Genre:  Genre{Name: "Science Fiction"},
		Title:  "Dune",
		Author: "Frank Herbert",
	}
	fmt.Println(book)
}
