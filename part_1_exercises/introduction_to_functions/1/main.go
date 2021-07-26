package main

import "fmt"

func multiply(x int, y int) int {
	return x * y
}

func main() {
	result := multiply(5, 4)
	fmt.Println(result) // Expected output: 20
}
