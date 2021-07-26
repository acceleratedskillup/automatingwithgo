package example2

import (
	"fmt"
)

func main() {
	i := writeToFile()
	fmt.Println("i:", i)

	fmt.Println("main() exiting")
}

func writeToFile() (i int) {
	fmt.Println("writeToFile() called")

	defer fmt.Println("does this get printed after recover() gets called?")
	//defer closeFile()

	/*
		example 2
		alternatively, if you want to return a valid value, put in a closure
		so that you can set the return value
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered in closure: \"%s\"\n", err)
			i = 15
		}
	}()

	if true {
		panic("Panic message")
	}
	fmt.Println("createFile() exiting")
	return 10
}
