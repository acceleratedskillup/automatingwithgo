package main

import "fmt"

func main() {

	/*
		we cannot assign an array to another array of DIFFERENT length

		when we assign an array to another array of same length,
		we copy the contents over to the new array.

		Take special note of this if your array is very large. Assigning large arrays
		can be very expensive
	*/
	type (
		categoryType      string
		categoryArrayType [3]categoryType
		codeArrayType     [3]string
	)

	var categories = [3]string{"a", "b", "c"}
	var codes = [3]string{"a", "b", "c"}
	fmt.Printf("categories: %#v\n", categories)
	fmt.Printf("codes: %#v\n", codes)
	fmt.Printf("equal?: %v\n", categories == codes)

}
