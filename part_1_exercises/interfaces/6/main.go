package main

import "fmt"

type Tool interface{}

type Magnifier struct{}
type Notebook struct{}

func main() {
	var tool1 Tool = Magnifier{}
	var tool2 Tool = Notebook{}

	fmt.Println(tool1 == tool2)
}
