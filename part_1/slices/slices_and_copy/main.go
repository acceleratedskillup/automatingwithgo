package main

import "fmt"

func main() {

	example1()
	//example2()
	//example3()
	//example4()
	//example5()
}

//Copy from one slice to another
func example1() {
	intSlice1 := []int{0, 1, 2, 3, 4}
	intSlice2 := make([]int, len(intSlice1))
	elementsCopied := copy(intSlice2, intSlice1)
	fmt.Println("elementsCopied:", elementsCopied)
	printSlice("intSlice1", intSlice1)
	printSlice("intSlice2", intSlice2)
}

//when the destination slice is smaller
func example2() {
	intSlice1 := []int{0, 1, 2, 3, 4}
	intSlice2 := make([]int, 1)
	elementsCopied := copy(intSlice2, intSlice1)
	fmt.Println("elementsCopied:", elementsCopied)
	printSlice("intSlice1", intSlice1)
	printSlice("intSlice2", intSlice2)
}

//when the src slice is smaller
func example3() {
	intSlice1 := []int{0, 1, 2, 3, 4}
	intSlice2 := make([]int, 10)
	elementsCopied := copy(intSlice2, intSlice1)
	fmt.Println("elementsCopied:", elementsCopied)
	printSlice("intSlice1", intSlice1)
	printSlice("intSlice2", intSlice2)
}

//Copy from a slice to itself
func example4() {
	intSlice1 := []int{0, 1, 2, 3, 4}
	fmt.Println("before")
	printSlice("intSlice1", intSlice1)
	elementsCopied := copy(intSlice1, intSlice1[2:])
	fmt.Println("after")
	fmt.Println("elementsCopied:", elementsCopied)
	printSlice("intSlice1", intSlice1)
}

func example5() {
	name := "john"
	byteSlice := make([]byte, len(name))
	elementsCopied := copy(byteSlice, name)
	fmt.Println("elementsCopied:", elementsCopied)
	printByteSlice("byteSlice", byteSlice)
}

func printSlice(prefix string, input []int) {
	fmt.Printf("%s) len():%d cap():%d s-addr:%p a-addr:%p\n", prefix, len(input), cap(input), &input, input)
	fmt.Printf("%*v %v \n", len(prefix)+1, " ", input)
}
func printByteSlice(prefix string, input []byte) {
	fmt.Printf("%s) len():%d cap():%d s-addr:%p a-addr:%p\n", prefix, len(input), cap(input), &input, input)
	fmt.Printf("%*v %v \n", len(prefix)+1, " ", input)
}
