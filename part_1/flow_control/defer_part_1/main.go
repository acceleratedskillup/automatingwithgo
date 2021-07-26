package main

import (
	"fmt"
	"os"
)

func main() {

	//example 1
	fmt.Println("main function has started.")
	defer fmt.Println("1) I am executed after the main function")
	defer fmt.Println("2) I am also executed after the main function")
	fmt.Println("main function has ended.")
	//fmt.Println("main function exiting")

	//example 2
	file, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		negative example where file is closed at the end of the processing
		but might not get executed due to conditional code or errors happening before
		this line gets executed
		.
		.
		do something with file
		.
		.
		error := true
		if error {
			return
		}
		file.Close()

		better way is to put the defer statement right after file open,
		guaranteeing that the file will be closed
		no matter what happens, right after the function ends
	*/
	defer file.Close()
	/*
		.
		.
		do something with file...
		.
		.
	*/

	fmt.Println("main function exiting")
}
