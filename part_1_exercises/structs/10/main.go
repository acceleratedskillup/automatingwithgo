package main

import "fmt"

type Edition struct {
	Number int
	Year   int
}

type Book struct {
	Title  string
	Author string
	Ed     Edition
}

func main() {
	book := Book{
		Title:  "Go Mastery",
		Author: "Lucas",
		Ed:     Edition{Number: 2, Year: 2023},
	}
	fmt.Println(book)
}
