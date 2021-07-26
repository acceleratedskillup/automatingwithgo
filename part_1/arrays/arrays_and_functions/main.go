package main

import "fmt"

func main() {
	intArray := [3]int{1, 2, 3}

	returnedArray := printArray(intArray)
	fmt.Println("intArray:", intArray)
	fmt.Printf("&returnedArray:%p\n", &returnedArray)
	fmt.Println("returnedArray:", returnedArray)
}

func printArray(inputArray [3]int) [3]int {
	inputArray[2] = 5000
	fmt.Println("inputArray:", inputArray)
	fmt.Printf("&inputArray:%p\n", &inputArray)
	return inputArray
}
