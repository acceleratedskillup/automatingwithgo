package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(minus(3, 1))
	dummyFunc("testString", 1, 2, 3.142)
	fmt.Println(divide1(17, 5))
	fmt.Println(divide2(17, 5))
	print("an series of ints", 1, 2, 3, 4)
	print("an series of ints", []int{1, 2, 3, 4}...)
	print("an series of ints")

}

//one way to declare a function
func add(x int, y int) int {
	return x + y
}

//another way to declare a function if the input parameters are of the same type
func minus(x, y int) int {
	return x - y
}

//another way to declare a function if the input parameters are of the same type
func dummyFunc(s string, x, y int, z float32) {
	fmt.Println(s, x, y, z)
}

//returning multiple values
func divide1(x, y int) (int, int) {
	return x / y, x % y
}

/*
	you can name your return variables and use them in the function like
	any other declared variable
*/

func divide2(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}

/*
	If the last parameter of a function has type ...T,
	it can be called with a trailing, variable number of arguments of type T
	this is called a variadic function

*/
func print(description string, inputInt ...int) {
	//fmt.Println(description)
	fmt.Println(description, inputInt)
	for _, v := range inputInt {
		fmt.Printf("%d,", v)
	}
	fmt.Println()
}
