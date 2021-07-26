package main

import "fmt"

func safeDivision(numerator, denominator int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Division by zero error")
		}
	}()
	if denominator == 0 {
		panic("Denominator cannot be zero")
	}
	return numerator / denominator
}

func main() {
	fmt.Println(safeDivision(10, 2)) // Expected output: 5
	fmt.Println(safeDivision(10, 0)) // Expected output: Division by zero error
}
