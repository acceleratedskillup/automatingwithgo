package main

import (
	"fmt"
)

// you got to use "var" outside of functions
var numberOfProcessors = 4

func main() {

	//part 1:
	var height int = 12
	fmt.Println(height)

	//part 2: declaring a variable but initializing later
	var carModel string
	carModel = "toyota"

	//part 2.1:declaring multiple variables of the same data type
	var i, j, k int
	i = 1
	j = 2
	k = 3

	//part 2.2:declaring and initializing multiple variables at the same time
	//var i, j int = 1, 2

	//part 2.3: grouping variables together might indicate some relationship
	var (
		person1 = "mark"
		person2 = "jane"
	)

	//part 2: this Println is a dummy statement to show use the variables declared above or golang will complain about unused variables
	fmt.Println(carModel, i, j, k, person1, person2, numberOfProcessors, i, j)

	//part 2.4:another way of creating int variable
	//int1 := 123
	int1 := int(123)
	//part 2.5:if we want explicitly a signed 64 bit integer
	int2 := int64(123)

	float1 := 3.142
	float2 := float64(3.142)

	fmt.Printf("int1: %T\n", int1)
	fmt.Printf("int2: %T\n", int2)
	fmt.Printf("float1: %T\n", float1)
	fmt.Printf("float2: %T\n", float2)

	//part 2.6:redeclaration of variables only allowed when it is used with at least one other new variable
	int3 := 5
	int3, int4 := 50, 60
	fmt.Println(int3, int4)

	//part 2.7:creating new aliases for existing data types
	type millimeter float32
	//var lengthOfInsect millimeter = 3.15

	//alternative declaration for aliases
	lengthOfInsect := millimeter(3.15)
	fmt.Println(lengthOfInsect)

	//------------------------------------
	//part 3:printing out default values of variables if there are not initialized with an explicit value
	var int5 int
	var truefalse1 bool
	var string2 string
	fmt.Println(int5, truefalse1, ":", string2, ":")

}
