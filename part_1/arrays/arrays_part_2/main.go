package main

import "fmt"

func main() {

	var weights = [3]int{1, 2, 3}
	fmt.Printf("weights array %v\n", weights)
	//weights := [3]int{1, 2, 3}
	//weights := [...]int{1, 2, 3}
	fmt.Printf("length of weights array %v\n", len(weights))

	//names := [2]string{"mary", "john"}
	names := [10]string{
		"mary",
		"john",
	}
	fmt.Printf("names %q\n", names)

	//weights := [3]int{1, 2}
	//fmt.Printf("weights array %v\n", weights)

	keyedNames := [10]string{
		2: "mary",
		5: "john",
		"may",
	}
	fmt.Printf("keyedNames %q\n", keyedNames)

}
