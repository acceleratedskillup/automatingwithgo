package main

import "fmt"

type Book struct {
	title  string
	author string
}

type Library struct {
	books []Book
}

func (l *Library) addBook(book Book) {
	l.books = append(l.books, book)
}

func (l *Library) displayCollection() {
	fmt.Println("Library Collection:")
	for _, book := range l.books {
		fmt.Println("-", book.title, "by", book.author)
	}
}

func main() {
	library := &Library{}
	library.addBook(Book{title: "Moby Dick", author: "Herman Melville"})
	library.addBook(Book{title: "1984", author: "George Orwell"})

	library.displayCollection()
}
