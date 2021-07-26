package main

import "fmt"

func main() {

	//deferred statements are executed when panic is called
	defer fmt.Println("I am executed after the main function")

	someFunction()
	fmt.Println("main function exiting")
}

func someFunction() {
	//deferred statements are executed when panic is called
	defer fmt.Println("I am executed after someFunction finishes")

	/*
		The function takes a single argument of arbitrary type—often a string—to be printed as the program dies
		this causes a program exception, have a look at the stacktrace also
	*/
	fmt.Println("Panicking now!!!")
	panic("Stop")
	fmt.Println("will i get executed?") //this line doesn't get executed
}
