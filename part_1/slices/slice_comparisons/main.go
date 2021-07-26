package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {

	//example1()
	//example2()
	//example3()
	//example4()
	//example5()
	example6()

}

/*
compare equal length slices using for loops
with no length check
*/
func example1() {
	intArray1 := [3]int{1, 2, 3}
	intArray2 := [3]int{1, 2, 3}
	fmt.Println("areArraysEqual:", intArray1 == intArray2)

	intSlice1 := []int{1, 2, 3}
	intSlice2 := []int{1, 2, 3}

	areSlicesEquals := true
	for index, value := range intSlice1 {
		if value != intSlice2[index] {
			areSlicesEquals = false
		}
	}
	fmt.Println("areSlicesEquals:", areSlicesEquals)
}

/*
negative example:
compare UNequal length slices using for loops
with no length check
*/
func example2() {
	intSlice1 := []int{1, 2}
	//intSlice1 := []int{} //empty slice
	//var intSlice1 []int // nil slice
	intSlice2 := []int{1, 2, 3}

	areSlicesEquals := true
	for index, value := range intSlice1 {
		if value != intSlice2[index] {
			areSlicesEquals = false
			break
		}
	}
	fmt.Println("areSlicesEquals:", areSlicesEquals)
}

/*
positive example to fix the above negative example:
compare UNequal length slices using for loops
with length check
*/
func example3() {
	intSlice1 := []int{1, 2}
	//intSlice1 := []int{} //empty slice
	//var intSlice1 []int // nil slice
	intSlice2 := []int{1, 2, 3}

	areSlicesEquals := true

	if len(intSlice1) != len(intSlice2) {
		areSlicesEquals = false
	} else {
		for index, value := range intSlice1 {
			if value != intSlice2[index] {
				areSlicesEquals = false
			}
		}
	}
	fmt.Println("areSlicesEquals:", areSlicesEquals)
}

//bytes.Equal
func example4() {

	//Using shorthand declaration to initializing byte slice
	byteSlice1 := []byte("hello")
	byteSlice2 := []byte("hello world")

	//example 4.1
	//var byteSlice2 []byte //nil slice

	fmt.Println("byteSlice1", byteSlice1)
	fmt.Println("byteSlice2", byteSlice2)
	areSlicesEquals := bytes.Equal(byteSlice1, byteSlice2)
	fmt.Println("areSlicesEquals:", areSlicesEquals)
}

/*
bytes.Compare
Compare returns an integer comparing two byte slices
lexicographically
"Lexicographical order is the dictionary order
or preferably the order in which words appear in the dictonary"
*/
func example5() {

	//byteSlice1 < byteSlice2
	byteSlice1 := []byte("hello")
	byteSlice2 := []byte("iello")

	//example 5.1
	//byteSlice1 > byteSlice2
	/*
		byteSlice1 := []byte("hello")
		byteSlice2 := []byte("hbllo")
	*/

	//example 5.2
	//byteSlice1 < byteSlice2
	/*
		byteSlice1 := []byte("hello")
		byteSlice2 := []byte("hello world")
	*/

	//example 5.3
	//byteSlice1 > byteSlice2
	/*
		byteSlice1 := []byte("hello")
		byteSlice2 := []byte("borld123")
	*/

	//example 5.4
	//byteSlice1 > byteSlice2
	//var byteSlice2 []byte //nil slice

	areSlicesEquals := bytes.Compare(byteSlice1, byteSlice2)

	if areSlicesEquals == 0 {
		fmt.Println("slices are equal")
	} else if areSlicesEquals > 0 {
		fmt.Println("byteSlice1 > byteSlice2")
	} else {
		fmt.Println("byteSlice1 < byteSlice2")
	}

}

/*
reflect.DeepEqual()
performance is worse than a for loop
*/
func example6() {

	intSlice1 := []int{1, 2, 3}
	intSlice2 := []int{1, 2, 3}

	//example 6.1
	/*
		intSlice1 := []int{1, 2, 3}
		byteSlice2 := []byte{1, 2, 3}
	*/

	//example 6.2
	/*
		intSlice1 := []int{1, 2, 3}
		intSlice2 := intSlice1[1:]
	*/

	areSlicesEquals := reflect.DeepEqual(intSlice1, intSlice2)

	fmt.Println("areSlicesEquals:", areSlicesEquals)

}
