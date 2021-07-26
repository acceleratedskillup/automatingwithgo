package main

import "fmt"

func operate(x int, y int, fn func(int, int) int) int {
	return fn(x, y)
}

func main() {
	resultAdd := operate(5, 3, func(a int, b int) int { return a + b })
	resultSub := operate(5, 3, func(a int, b int) int { return a - b })
	fmt.Println(resultAdd) // Expected output: 8
	fmt.Println(resultSub) // Expected output: 2
}
