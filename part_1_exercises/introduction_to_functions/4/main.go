package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	countFunc := counter()
	fmt.Println(countFunc()) // Expected output: 1
	fmt.Println(countFunc()) // Expected output: 2
	fmt.Println(countFunc()) // Expected output: 3
}
