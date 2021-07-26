package main

import "fmt"

func main() {
	var numbers = [3]int{1, 2, 3}
	fmt.Printf("Array address: %p\n", &numbers)
	fmt.Printf("First element address: %p\n", &numbers[0])

}
