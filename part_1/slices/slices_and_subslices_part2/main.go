package main

import "fmt"

func main() {

	intArray := [10]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("intArray) a-addr:%p\n", &intArray)

	//example1()
	//example2()
	//example3()
	//example4()
}

/*
show capacity of returned slice has capacity set from low
cap(input) — low index, which is default behavior
*/
func example1() {
	intArray := [10]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("intArray) a-addr:%p\n", &intArray)
	intSlice1 := intArray[2:]
	printSlice("intSlice1", intSlice1)
}

/*
example of full slice and change in capacity of returned
slice
*/
func example2() {
	intArray := [10]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("intArray) a-addr:%p\n", &intArray)
	intSlice1 := intArray[:2:3]
	printSlice("intSlice1", intSlice1)
}

/*
show that capacity value limited by the input slice
not backing array
*/
func example3() {
	intArray := [10]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("intArray) a-addr:%p\n", &intArray)
	intSlice1 := intArray[:2:3]
	//intSlice1 := intArray[:2:20]
	printSlice("intSlice1", intSlice1)
	intSlice2 := intSlice1[:2:5] // will throw error
	printSlice("intSlice2", intSlice2)
}

/*
show when appending to returned slice
will create new backing array
*/
func example4() {
	intArray := [10]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("intArray) a-addr:%p\n", &intArray)
	intSlice1 := intArray[:2:3]
	printSlice("intSlice1", intSlice1)
	fmt.Println("before appending 1")
	intSlice1 = append(intSlice1, 100)
	printSlice("intSlice1", intSlice1)
	fmt.Println("before appending 2")
	intSlice1 = append(intSlice1, 200)
	printSlice("intSlice1", intSlice1)
}
func printSlice(prefix string, input []int16) {
	fmt.Printf("%s) len():%d cap():%d s-addr:%p a-addr:%p\n", prefix, len(input), cap(input), &input, input)
	fmt.Printf("%*v %v \n", len(prefix)+1, " ", input)
}
