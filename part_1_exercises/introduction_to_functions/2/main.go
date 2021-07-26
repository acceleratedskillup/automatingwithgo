package main

import "fmt"

func swap(a string, b string) (string, string) {
	return b, a
}

func main() {
	first, second := swap("Hello", "World")
	fmt.Println(first, second) // Expected output: World Hello
}
