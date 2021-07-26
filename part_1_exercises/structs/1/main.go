package main

import "fmt"

type Book struct {
	Title  string
	Author string
	ISBN   string
	Pages  int
}

func main() {
	book := Book{
		Title:  "Go Programming",
		Author: "John Doe",
		ISBN:   "1234567890",
		Pages:  350,
	}
	fmt.Println(book)
}
