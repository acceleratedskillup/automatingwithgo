package main

import "fmt"

type DetectiveTool interface{}

type MagnifyingGlass struct{}

func main() {
	var tool DetectiveTool
	var glass *MagnifyingGlass

	fmt.Println(tool == nil)   // true
	fmt.Println(glass == nil)  // true
	fmt.Println(tool == glass) // false
}
