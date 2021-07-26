package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	book1 := Book{"Go Basics", "John Doe", 200}
	book2 := Book{"Go Basics", "John Doe", 200}

	isSame := book1 == book2
	fmt.Println("Are the books the same?", isSame)
}
