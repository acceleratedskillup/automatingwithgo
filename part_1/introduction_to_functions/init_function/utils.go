package main

import "fmt"

func init() {
	fmt.Println("in init() of utils.go varA:", varA)
	fmt.Println("in init() of utils.go 1")

}

func init() {
	fmt.Println("in init() of utils.go 2")
}

var varA = 1
