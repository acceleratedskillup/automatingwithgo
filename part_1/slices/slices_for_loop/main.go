package main

import "fmt"

func main() {
	//example1()
	//example2()
	//example3()
	//example4()
	example5()
}

func example1() {
	//scenarios: append, delete
	intSlice1 := []int{0, 1, 2, 3, 4}

	for index := 0; index < len(intSlice1); index++ {
		intSlice1 = append(intSlice1, index)
		printSlice("intSlice1", intSlice1, index)
	}
}

func example2() {
	//scenarios: append, delete
	intSlice1 := []int{0, 1, 2, 3, 4}

	for index := range intSlice1 {
		intSlice1 = append(intSlice1, index)
		printSlice("intSlice1", intSlice1, index)
	}
}

func example3() {
	intSlice1 := []int{0, 1, 2, 3, 4}
	for index := range intSlice1 {
		intSlice1 = intSlice1[index:]
		printSlice("intSlice1", intSlice1, index)
	}
}

func example4() {
	intSlice1 := []int{0, 1, 2, 3, 4}

	for index := 0; index < len(intSlice1); index++ {
		intSlice1 = intSlice1[index:]
		printSlice("intSlice1", intSlice1, index)
	}
}

func example5() {
	var intSlice1 []int
	fmt.Println("before loop")
	for index := range intSlice1 {
		intSlice1 = intSlice1[index:]
		printSlice("intSlice1", intSlice1, index)
	}

	for index := 0; index < len(intSlice1); index++ {
		intSlice1 = intSlice1[index:]
		printSlice("intSlice1", intSlice1, index)
	}
	fmt.Println("after loop")
}

func printSlice(prefix string, input []int, index int) {
	fmt.Printf("%s) index:%d len():%d cap():%d\n", prefix, index, len(input), cap(input))
	fmt.Printf("%*v %v \n", len(prefix)+1, " ", input)
}
