package helloworld

import "fmt"

func init() {
	fmt.Println("in init() of helloworld.go 1")
}

func init() {
	fmt.Println("in init() of helloworld.go 2")
}

func Run() {
	fmt.Println("in Run() of helloworld.go")
}
