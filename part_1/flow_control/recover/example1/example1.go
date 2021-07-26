package example1

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
	defer closeFile()

	if true {
		panic("Panic message")
	}
	fmt.Println("createFile() exiting")
	return 10
}

func closeFile() {
	fmt.Println("closeFile() called")

	//if we comment out this entire section, the program will terminate
	/*
		if err := recover(); err != nil {
			fmt.Printf("Recovered in closeFile(): \"%s\"\n", err)
		}
	*/
	fmt.Println("closeFile() exiting")
}
