package main

import "fmt"

func calculate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func main() {
	fmt.Println(calculate(5, 3, "+")) // Expected output: 8
	fmt.Println(calculate(5, 3, "-")) // Expected output: 2
}
