package main

import (
	"fmt"
	"strings"
)

//define a new type
type validateFunc func(string) bool

func main() {

	/*
		you can assign a function literal to a variable
		A function literal is a function without a name,
		also called anonymous function
	*/
	validate := func(s string) bool {
		return strings.HasPrefix(s, "h")
	}

	//you can pass a function to another function
	fmt.Println(validateThreeStrings(
		"happy", "hospital", "handle", validate))

	fmt.Println(validateThreeStrings(
		"happy", "apple", "handle", validate))

	validateWithPrefixH := createValidateFunc("h")
	fmt.Println(validateThreeStrings(
		"happy", "hospital", "handle", validateWithPrefixH))

}

func validateThreeStrings(s1 string, s2 string, s3 string, validate func(s string) bool) bool {
	return validate(s1) && validate(s2) && validate(s3)
}

func createValidateFunc(validateString string) validateFunc {
	return func(s string) bool {
		return strings.HasPrefix(s, validateString)
	}
}
