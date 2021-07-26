package main

import "fmt"

type Car struct {
	Model       string
	PlateNumber string
}

func main() {
	//example1()
	//example1_1()
	//example2()
	//example3()
	//example4()
	example5()

}

func example1() {
	//example 1
	//declaring a int pointer which after declaration will be nil.
	var honda1 *Car
	//take note that we are using honda in the 2nd parameter, not *honda
	fmt.Printf("honda1) type:%T, honda1:%v\n", honda1, honda1)
}

const INDENT = 7

func example1_1() {

	/*
		use the builtin new() function to obtain a pointer to a zero-valued struct
		we covered this before in previous lecture on structs
		but we didn't see that it was a struct pointer then
		we used it just like any other variable
	*/
	var honda1 *Car = new(Car)
	/*
		ALTERNATIVELY
		honda1 := new(Car)
	*/
	fmt.Printf("honda1) type:%T, &honda1:%p, honda1:%p\n", honda1, &honda1, honda1)
	fmt.Printf("%*v *honda1:%#v \n", INDENT, " ", *honda1)

	(*honda1).PlateNumber = "S1234"
	honda1.Model = "Honda1"

	fmt.Println()
	fmt.Printf("honda1) type:%T, &honda1:%p, honda1:%p\n", honda1, &honda1, honda1)
	fmt.Printf("%*v *honda1:%#v \n", INDENT, " ", *honda1)
}

/*
	assignment to struct pointer variable
	from another struct pointer variable
*/
func example2() {

	var honda1 *Car = new(Car)
	//alternatively
	//honda1 := new(Car)
	fmt.Printf("honda1) type:%T, &honda1:%p, honda1:%p\n", honda1, &honda1, honda1)
	fmt.Printf("%*v *honda1:%#v \n", INDENT, " ", *honda1)
	(*honda1).PlateNumber = "S1234"
	honda1.Model = "Honda1"
	fmt.Println()
	fmt.Printf("honda1) type:%T, &honda1:%p, honda1:%p\n", honda1, &honda1, honda1)
	fmt.Printf("%*v *honda1:%#v \n", INDENT, " ", *honda1)

	var honda2 *Car
	honda2 = honda1

	/*
		ALTERNATIVELY
		var honda2 *Car = honda1
	*/

	fmt.Println()
	fmt.Printf("honda2) type:%T, &honda2:%p, honda2:%p\n", honda2, &honda2, honda2)
	fmt.Printf("%*v *honda2:%#v \n", INDENT, " ", *honda2)
}

/*
assigning the address of honda1 struct literal
to struct pointer honda2
*/
func example3() {
	honda1 := Car{
		Model:       "Honda1",
		PlateNumber: "S1234",
	}
	fmt.Printf("honda1) type:%T, &honda1:%p, honda1:%#v\n", honda1, &honda1, honda1)
	/*
		assigning the address of honda1 struct literal
		to honda2
	*/
	var honda2 *Car
	honda2 = &honda1

	/*
		ALTERNATIVELY 1
		var honda2 *Car = &honda1
	*/
	/*
		ALTERNATIVELY 2
		what happens when we create and
		initialize a new variable using &
		special prefix & returns a pointer to the struct value.
		honda2 := &honda1
	*/
	fmt.Println()
	fmt.Printf("honda2) type:%T, &honda2:%p, honda2:%p\n", honda2, &honda2, honda2)
	fmt.Printf("%*v *honda2:%#v \n", INDENT, " ", *honda2)
}

func example4() {
	/*
		can concise the above example
		can straight away init honda2 without honda1
	*/
	honda2 := &Car{
		Model:       "Honda2",
		PlateNumber: "S1234",
	}
	fmt.Printf("honda2) type:%T, &honda2:%p, honda2:%p\n", honda2, &honda2, honda2)
	fmt.Printf("%*v *honda2:%#v \n", INDENT, " ", *honda2)
}

func example5() {
	honda2 := &Car{
		Model:       "Honda2",
		PlateNumber: "S1234",
	}
	fmt.Printf("honda2) type:%T, &honda2:%p, honda2:%p\n", honda2, &honda2, honda2)
	fmt.Printf("%*v *honda2:%#v \n\n", INDENT, " ", *honda2)

	updatePlateNumber(honda2)

	fmt.Println("after update of plate number")
	fmt.Println()
	fmt.Printf("honda2) type:%T, &honda2:%p, honda2:%p\n", honda2, &honda2, honda2)
	fmt.Printf("%*v *honda2:%#v \n", INDENT, " ", *honda2)

}

func updatePlateNumber(input *Car) {
	input.PlateNumber = "A456"
	fmt.Printf("input) type:%T, &input:%p, input:%p\n", input, &input, input)
	fmt.Printf("%*v *input:%#v \n\n", INDENT-1, " ", *input)

}
