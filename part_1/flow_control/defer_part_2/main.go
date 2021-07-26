package main

import "fmt"

func main() {
	/*
		example 1

		code opens up database connections*
		.
		.
		code uses the database connections
		.
		.
		defer statements called in LIFO order
		A deferred function's arguments are evaluated when the defer statement is evaluated,
		and not when the call executes

		A deferred function's arguments are evaluated when the defer statement is evaluated.
	*/
	i := 1
	defer printNumber(i)
	i++
	defer printNumber(i)
	i++
	defer printNumber(i)

	//fmt.Println("main function exiting")

	/*
		example 2
		Deferred functions may read and assign to the returning function's named return values.
	*/

	fmt.Println("main function exiting")
}
func printNumber(i int) {
	fmt.Println("printNumber() is called with i:", i)
}
