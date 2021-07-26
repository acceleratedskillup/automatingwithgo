package main

import (
	"fmt"
	"interfaces/car"
	"strconv"
)

func main() {

	//example1()
	//example2()
	//example3()
	//example4()
	example5()
	//example6()
	//example7()

}

/*
usage of type empty interface{}
*/
func example1() {

	type empty interface{}
	var y empty = 5
	fmt.Printf("y concrete type:%T concrete value:%v\n", y, y)
}

/*
usage of interface{} directly
*/
func example2() {
	var y interface{} = 10
	fmt.Printf("y concrete type:%T concrete value:%v\n", y, y)
	y = "hello world"
	fmt.Printf("y concrete type:%T concrete value:%v\n", y, y)
}

/*
cannot call any Car method at all
*/
func example3() {

	var y interface{} = car.Car{
		Model: "toyota",
	}
	fmt.Printf("y concrete type:%T concrete value:%+v\n", y, y)
	//cannot call any Car method at all
	//y.
}

/*
need to use type assertion
*/
func example4() {

	var y interface{} = car.Car{
		Model: "toyota",
	}
	fmt.Printf("y concrete type:%T concrete value:%+v\n", y, y)

	y.(car.Car).Start()
}

func example5() {
	fmt.Printf("concrete type:%T\n", ToString(10))
	fmt.Printf("concrete type:%T\n", ToString("interfaces are cool"))
	fmt.Printf("concrete type:%T\n", ToString(float32(1.23)))
	fmt.Printf("concrete type:%T\n", ToString(1.23))
}

func ToString(input interface{}) interface{} {
	switch value := input.(type) {
	case int:
		return strconv.Itoa(value)
	case string:
		return value
	case float32:
		return fmt.Sprintf("%f", value)
	default:
		return input
	}
}

func example6() {
	printAnything(1, 1.23, "hello")
}

func printAnything(input ...interface{}) {
	for _, value := range input {
		fmt.Printf("printAnything(): concrete type:%T concrete value:%v\n", value, value)
	}
}

/*
slice of interfaces is not an empty interface
*/
func example7() {
	var y []interface{}
	names := []string{"mary", "john", "karl"}
	//y = names

	for _, value := range names {
		y = append(y, value)
	}
	fmt.Printf("y concrete type:%T concrete value:%v\n", y, y)
}
