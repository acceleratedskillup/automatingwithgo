package main

import "fmt"

func lendBook(library map[string]string, book, borrower string) {
	library[book] = borrower
}

func returnBook(library map[string]string, book string) {
	library[book] = ""
}

func main() {
	library := map[string]string{
		"1984":            "",
		"Brave New World": "",
		"Dune":            "",
	}

	lendBook(library, "1984", "John")
	fmt.Println("After lending '1984':", library)

	returnBook(library, "1984")
	fmt.Println("After returning '1984':", library)
}
