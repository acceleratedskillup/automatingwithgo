package main

import (
	"fmt"
)

func main() {

	x := 5
	//simple if example
	if x < 0 {
		fmt.Println("x is < 0")
	} else {
		fmt.Println("x is >= 0")
	}

	/*
		if statement prefixed with a short statement
		Variables declared by the statement are only in scope until the end of the if/else.
	*/
	if y := getY(); y < 0 {
		fmt.Printf("y:%d is < 0\n", y)
	} else {
		fmt.Printf("y:%d is >= 0\n", y)
	}

	//variable y is not accessible here
	//fmt.Printf("y:%d \n", y)

	//nested if statements
	if y := getY(); y < 0 {
		fmt.Printf("y:%d is < 0\n", y)
	} else if y >= 0 && y < 5 {
		fmt.Printf("y:%d is >= to 0 but < 5\n", y)
	} else {
		fmt.Printf("y:%d is > 5\n", y)
	}
}

func getY() int {
	return 4
}
