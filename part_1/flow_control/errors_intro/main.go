package main

import (
	"fmt"
	"os"
)

func main() {

	//how go throws errors and how to check them
	_, err := os.Open("temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	//how to throw your own error
	returnValue, err := functionThatReturnsError(1)
	fmt.Println(returnValue)
	fmt.Println(err)
}

func functionThatReturnsError(input int) (int, error) {
	//one way of creating an error
	//return errors.New("error from function")

	//another way of creating an error
	return input, fmt.Errorf("formatted %s from function", "error")
}
