package main

import "fmt"

func isEven(num int) string {
	if num%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func main() {
	fmt.Println(isEven(4)) // Expected output: even
	fmt.Println(isEven(7)) // Expected output: odd
}
