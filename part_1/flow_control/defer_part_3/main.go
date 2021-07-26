package main

import "fmt"

func main() {

	/*
		Deferred functions may read and assign to the returning function's named return values.
	*/
	fmt.Println("return value from outerFunction():", outerFunction())

}

func outerFunction() (i int) {
	defer func() { i = 5 }() //don't forget to call the function!
	return 1
}
