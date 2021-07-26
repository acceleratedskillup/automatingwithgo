package utils

import "fmt"

/*
Always remember that variables and functions that start
with a capital letter is an export of the package

*/
var ToPrint bool = false

func PrintDummyString() {
	fmt.Println("in PrintDummyString")
}
