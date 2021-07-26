package main

import (
	"fmt"
)

func main() {

	//normal usage of switch statement
	i := 1
	switch i {
	case 1:
		fmt.Println("i is 1") //break statement is automatically applied for each case
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default: //will be executed after all cases above this are not executed
		fmt.Println("unknown i value")
	}

	switch i {
	case 1, 2, 3:
		fmt.Println("i is 1,2,3") //break statement is automatically applied for each case
	case 4:
		fmt.Println("i is 4")
	default: //will be executed after all cases above this are not executed
		fmt.Println("unknown i value")
	}

	switch i := 1; true {
	case i == 1:
		fmt.Println("i is 1")
	case i == 2:
		fmt.Println("i is 2")
	case i == 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("unknown value i:", i)
	}

	/*
		take note need to remove i from
		switch expression as the result of
		case expression is boolean instead
		of integer so cannot have i has
		part of switch expression
	*/
	switch {
	case 1 <= i && i <= 5:
		fmt.Println("i is between 1 and 5") //break statement is automatically applied for each case
	case 6 <= i && i <= 10:
		fmt.Println("i is between 6 and 10")
	default: //will be executed after all cases above this are not executed
		fmt.Println("unknown i value")
	}

	/*
		we can use an expression/function at the switch statement and also
		multiple expressions/functions when evaluating cases
	*/

	i = 2
	j := 1

	switch i + j {
	case getValue(1):
		fmt.Println("sum is 1") //break statement is automatically applied for each case
	case 1 + 1:
		fmt.Println("sum is 2")
	case getValue(1) + getValue(2):
		fmt.Println("sum is 3")
	case getValue(4), getValue(5):
		fmt.Println("sum is 4 or 5")
	default: //will be executed after all cases above this are not executed
		fmt.Println("unknown sum value")
	}

	i = 1
	switch i {
	case 1:
		fmt.Println("i is 1")
		fallthrough //need to end of case block
	case 2: // cannot have 2-1 as case number is not unique
		fmt.Println("i is 2")
	default:
		fmt.Println("unknown i value:", i)
	}

Start:
	for i = 0; i < 5; i++ {
		switch i {
		case 0:
			fmt.Println("i is 1")
			break
		case 3:
			break Start
		default:
			fmt.Println("i:", i)
		}
	}
	fmt.Println("loop is finished")
}

func getValue(input int) int {
	fmt.Println("getValue input:", input)
	return input
}
