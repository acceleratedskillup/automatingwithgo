package main

import (
	"fmt"
)

func main() {

	//example 4
	fmt.Println("calling recover outside defer functions:", recover())

	fmt.Println("main() exiting")
}
