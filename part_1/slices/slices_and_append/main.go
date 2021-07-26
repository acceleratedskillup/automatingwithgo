package main

import "fmt"

func main() {

	intSlice1 := []int{1, 2, 3}
	printSlice("intSlice1", intSlice1)

	//this will not work
	//intSlice1[3] = 4
	/*
		The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
		If the backing array of s is too small to fit all the given values a bigger array will be allocated.
		The returned slice will point to the newly allocated array.
	*/
	intSlice1 = append(intSlice1, 4, 5)
	printSlice("intSlice1", intSlice1)

	intSlice1 = append(intSlice1, 6)
	printSlice("intSlice1", intSlice1)

	//another way to use append by using another slice
	intSlice2 := []int{7, 8, 9}
	/*
		"..." means to unpack the slice
		so that it is equals to append(intSlice1, 3, 4)
	*/
	intSlice1 = append(intSlice1, intSlice2...)
	//intSlice1 = append(intSlice1, 7, 8, 9)
	printSlice("intSlice1", intSlice1)

	/*
		var intSlice2 []int
		printSlice("intSlice2", intSlice2)
		intSlice2 = append(intSlice2, intSlice1...)
		printSlice("intSlice2", intSlice2)
	*/

}

func printSlice(prefix string, input []int) {
	fmt.Printf("%s) len():%d cap():%d s-addr:%p a-addr:%p\n", prefix, len(input), cap(input), &input, input)
	fmt.Printf("%*v %v \n", len(prefix)+1, " ", input)
}
