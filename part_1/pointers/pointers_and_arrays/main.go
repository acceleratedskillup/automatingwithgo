package main

import "fmt"

const (
	INDENT = 8
	WIDTH  = 0
	HEIGHT = 1
	DEPTH  = 2
)

func main() {
	//example1()
	//example1_1()
	//example2()
	example3()
	//example4()
	//example5()
}

func example1() {
	var dimPtr1 *[3]int = new([3]int)

	fmt.Printf("dimPtr1) type:%T, &dimPtr1:%p, dimPtr1:%p\n", dimPtr1, &dimPtr1, dimPtr1)
	fmt.Printf("%*v *dimPtr1:%#v \n", INDENT, " ", *dimPtr1)
}

func example1_1() {
	var dimPtr1 *[3]int = new([3]int)

	fmt.Printf("dimPtr1) type:%T, &dimPtr1:%p, dimPtr1:%p\n", dimPtr1, &dimPtr1, dimPtr1)
	fmt.Printf("%*v *dimPtr1:%#v \n", INDENT, " ", *dimPtr1)

	(*dimPtr1)[WIDTH] = 100
	dimPtr1[HEIGHT] = 200
	dimPtr1[DEPTH] = 300

	fmt.Println()
	fmt.Printf("dimPtr1) type:%T, &dimPtr1:%p, dimPtr1:%p\n", dimPtr1, &dimPtr1, dimPtr1)
	fmt.Printf("%*v *dimPtr1:%#v \n", INDENT, " ", *dimPtr1)
}

func example2() {
	var dimPtr1 *[3]int = new([3]int)

	fmt.Printf("dimPtr1) type:%T, &dimPtr1:%p, dimPtr1:%p\n", dimPtr1, &dimPtr1, dimPtr1)
	fmt.Printf("%*v *dimPtr1:%#v \n", INDENT, " ", *dimPtr1)

	(*dimPtr1)[WIDTH] = 100
	dimPtr1[HEIGHT] = 200
	dimPtr1[DEPTH] = 300

	fmt.Println()
	fmt.Printf("dimPtr1) type:%T, &dimPtr1:%p, dimPtr1:%p\n", dimPtr1, &dimPtr1, dimPtr1)
	fmt.Printf("%*v *dimPtr1:%#v \n", INDENT, " ", *dimPtr1)

	var dimPtr2 *[3]int = dimPtr1
	fmt.Println()
	fmt.Printf("dimPtr2) type:%T, &dimPtr2:%p, dimPtr2:%p\n", dimPtr2, &dimPtr2, dimPtr2)
	fmt.Printf("%*v *dimPtr2:%#v \n", INDENT, " ", *dimPtr2)
}

func example3() {
	dim := [3]int{100, 200, 300}

	fmt.Printf("dim) type:%T, &dim:%p, dim:%#v\n", dim, &dim, dim)

	var dimPtr1 *[3]int = &dim

	/*
		ALTERNATIVE
		dimPtr := &dim
	*/
	fmt.Println()
	fmt.Printf("dimPtr1) type:%T, &dimPtr1:%p, dimPtr1:%p\n", dimPtr1, &dimPtr1, dimPtr1)
	fmt.Printf("%*v *dimPtr1:%#v \n", INDENT, " ", *dimPtr1)
}

func example4() {

	dimPtr1 := &[3]int{100, 200, 300}

	fmt.Printf("dimPtr1) type:%T, &dimPtr1:%p, dimPtr1:%p\n", dimPtr1, &dimPtr1, dimPtr1)
	fmt.Printf("%*v *dimPtr1:%#v \n", INDENT, " ", *dimPtr1)
}

func example5() {

	dimPtr1 := &[3]int{100, 200, 300}

	fmt.Printf("dimPtr1) type:%T, &dimPtr1:%p, dimPtr1:%p\n", dimPtr1, &dimPtr1, dimPtr1)
	fmt.Printf("%*v *dimPtr1:%#v \n", INDENT, " ", *dimPtr1)

	modifyWidth(dimPtr1)
	fmt.Println()
	fmt.Printf("dimPtr1) type:%T, &dimPtr1:%p, dimPtr1:%p\n", dimPtr1, &dimPtr1, dimPtr1)
	fmt.Printf("%*v *dimPtr1:%#v \n", INDENT, " ", *dimPtr1)

}

func modifyWidth(input *[3]int) {
	input[WIDTH] = 999
	fmt.Println()
	fmt.Printf("input) type:%T, &input:%p, input:%p\n", input, &input, input)
	fmt.Printf("%*v *input:%#v \n", INDENT-2, " ", *input)

}
