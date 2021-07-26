package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	printSlice("x", x)

	x[0] = 100
	fmt.Println("x[0]:", x[0])

	//an empty slice is an initialized slice
	y := []int{}
	printSlice("y", y)

	//an empty slice is an initialized slice
	y1 := []int{}
	printSlice("y1", y1)

	y = x
	printSlice("x", x)
	printSlice("y", y)

	x[1] = 200
	printSlice("x", x)
	printSlice("y", y)

	/*
		nil slice
		a nil slice is an uninitialized slice
		as compared to array, its already initialized
	*/
	var z []int
	printSlice("z", z)
	var zArray [3]int
	fmt.Println("zArray:", zArray)

	// print address of 0th element of the slice
	//fmt.Printf("Address of 0th element of the slice: %p\n", &x[0])
	/*
		y[0]=0
	*/
	fmt.Println("is nil:", z == nil)

	/*
		take note that you do not need to specify the number of elements
		note that this is a nil slice which is an empty slice with 0 capacity
		and has no underlying array.
	*/
}

func printSlice(prefix string, input []int) {
	//s-addr: memory address of slice
	//a-addr: memory address of backing array
	/*
		cannot use &slice[0] to print out array address
		as it cannot be used by empty slice
	*/
	fmt.Printf("%s) len():%d cap():%d s-addr:%p a-addr:%p\n", prefix, len(input), cap(input), &input, input)
	fmt.Printf("%*v %v \n", len(prefix)+1, " ", input)
}
