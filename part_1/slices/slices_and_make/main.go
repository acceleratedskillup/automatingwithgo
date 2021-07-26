package main

import "fmt"

func main() {
	example1()
	//example2()
	//example3()
}

/*
use all 3 arguments
the capacity specified must be no smaller than the
length
*/
func example1() {
	intSlice1 := make([]int16, 2, 10)
	printSlice("intSlice1", intSlice1)
	intSlice1[0] = 100
	intSlice1[1] = 200
	//intSlice1[2] = 300
	intSlice1 = append(intSlice1, 300)
	printSlice("intSlice1", intSlice1)

}

/*
use first 2 arguments, The capacity of the slice is
equal to its length
*/
func example2() {
	stringSlice1 := make([]string, 2)
	printStringSlice("stringSlice1", stringSlice1)
}

//slice with 0 length = empty slice
func example3() {
	stringSlice1 := make([]string, 0)
	printStringSlice("stringSlice1", stringSlice1)
	stringSlice1 = append(stringSlice1, "john")
	printStringSlice("stringSlice1", stringSlice1)
}

func printSlice(prefix string, input []int16) {
	fmt.Printf("%s) len():%d cap():%d s-addr:%p a-addr:%p\n", prefix, len(input), cap(input), &input, input)
	fmt.Printf("%*v %v \n", len(prefix)+1, " ", input)
}

func printStringSlice(prefix string, input []string) {
	fmt.Printf("%s) len():%d cap():%d s-addr:%p a-addr:%p\n", prefix, len(input), cap(input), &input, input)
	fmt.Printf("%*v %#v \n", len(prefix)+1, " ", input)
}
