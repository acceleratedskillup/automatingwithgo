package main

import "fmt"

func main() {
	example1()
	//example2()
}

func example1() {
	intSlice1 := make([]int, 2)

	fmt.Printf("intSlice1) len():%d cap():%d s-addr:%p a-addr:%p\n", len(intSlice1), cap(intSlice1), &intSlice1, intSlice1)
	fmt.Printf("%v \n", intSlice1)
	mutateSlice(intSlice1)
	fmt.Printf("intSlice1) len():%d cap():%d s-addr:%p a-addr:%p\n", len(intSlice1), cap(intSlice1), &intSlice1, intSlice1)
	fmt.Printf("%v \n", intSlice1)
}

func example2() {
	intSlice1 := make([]int, 2)
	fmt.Printf("intSlice1) len():%d cap():%d s-addr:%p a-addr:%p\n", len(intSlice1), cap(intSlice1), &intSlice1, intSlice1)
	fmt.Printf("%v \n", intSlice1)
	intSlice1 = mutateSliceWithReturn(intSlice1)
	fmt.Printf("intSlice1) len():%d cap():%d s-addr:%p a-addr:%p\n", len(intSlice1), cap(intSlice1), &intSlice1, intSlice1)
	fmt.Printf("%v \n", intSlice1)
}

func mutateSlice(input []int) {
	fmt.Println("---mutateSlice start---")
	fmt.Printf("input) len():%d cap():%d s-addr:%p a-addr:%p\n", len(input), cap(input), &input, input)
	input[0] = 12
	fmt.Printf("%v \n", input)
	fmt.Println("---mutateSlice end---")
}

func mutateSliceWithReturn(input []int) []int {
	fmt.Println("---mutateSlice start---")
	fmt.Printf("input) len():%d cap():%d s-addr:%p a-addr:%p\n", len(input), cap(input), &input, input)
	fmt.Printf("%v \n", input)
	input = append(input, 1000)
	fmt.Printf("input) len():%d cap():%d s-addr:%p a-addr:%p\n", len(input), cap(input), &input, input)
	fmt.Printf("%v \n", input)
	fmt.Println("---mutateSlice end---")
	return input
}

func printSlice(prefix string, input []int) {
	fmt.Printf("%s) len():%d cap():%d s-addr:%p a-addr:%p\n", prefix, len(input), cap(input), &input, input)
	fmt.Printf("%*v %v \n", len(prefix)+1, " ", input)
}
