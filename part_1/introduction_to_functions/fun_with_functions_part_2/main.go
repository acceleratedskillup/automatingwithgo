package main

import "fmt"

func main() {

	/*
		Function literals in Go are closures
		They may refer to variables defined in an enclosing function.
	*/
	magicNumber := 2
	divideByMagicNumber := func(x int) (quotient, remainder int) {
		quotient = x / magicNumber
		remainder = x % magicNumber
		return
	}
	fmt.Println(divideByMagicNumber(10))

	printFunction := createPrintFunction("mary had a little lamb")
	printFunction()

	multiply := createMultiplyFunction()
	fmt.Println(multiply(3, 2))

}

// a function that returns a function
func createPrintFunction(toPrint string) func() {
	return func() {
		fmt.Println(toPrint)
	}
}

// a function that returns a more complex function
func createMultiplyFunction() func(x, y int) int {
	return func(x, y int) int {
		return x * y
	}
}
