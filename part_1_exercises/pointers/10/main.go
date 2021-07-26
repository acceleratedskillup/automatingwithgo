package main

import "fmt"

func main() {
	a, b := 5, 10
	pA, pB := &a, &b
	sum := *pA + *pB
	fmt.Println("Sum:", sum)
}
