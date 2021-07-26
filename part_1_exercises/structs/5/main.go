package main

import "fmt"

type Borrower struct {
	Name         string
	DateBorrowed string
}

type Book struct {
	Title     string
	Author    string
	Borrowers []Borrower
}

func main() {
	book := Book{
		Title:  "A Popular Book",
		Author: "Jane Smith",
		Borrowers: []Borrower{
			{Name: "John Doe", DateBorrowed: "2022-01-10"},
			{Name: "Alice", DateBorrowed: "2022-01-15"},
		},
	}
	fmt.Println(book)
}
