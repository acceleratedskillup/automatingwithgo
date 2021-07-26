package main

import (
	"fmt"
)

func main() {
	i := writeToFile()
	fmt.Println("i:", i)

	//example 4
	//fmt.Println("calling recover outside defer functions:", recover())

	fmt.Println("main() exiting")
}

func writeToFile() (i int) {
	fmt.Println("writeToFile() called")

	defer fmt.Println("does this get printed after recover() gets called?")
	//writes to file
	/*
		example 1
		if setting the return value is not important
	*/
	defer closeFile()

	/*
		example 2
		alternatively, if you want to return a valid value, put in a closure
		so that you can set the return value
	*/
	/*
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Recovered in closure: \"%s\"\n", err)
				i = 2
			}
		}()
	*/
	if true {
		panic("Panic message")
	}
	fmt.Println("createFile() exiting")
	return 10
}

func closeFile() {
	fmt.Println("closeFile() called")

	//if we comment out this entire section, the program will terminate
	if err := recover(); err != nil {
		fmt.Printf("Recovered in closeFile(): \"%s\"\n", err)
	}

	/*
		example 3
		what happens when we have panic again after calling recover?
	*/
	/*
		if true{
			panic("Panic message again")
		}
	*/
	fmt.Println("closeFile() exiting")
}
