package main

import (
	"fmt"
	"io"
)

type Library struct {
	Books []string
}

func (l *Library) Write(p []byte) (n int, err error) {
	book := string(p)
	l.Books = append(l.Books, book)
	return len(p), nil
}

func main() {
	lib := &Library{}
	io.WriteString(lib, "The Great Gatsby")
	fmt.Println(lib.Books)
}
