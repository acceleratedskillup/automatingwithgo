package main

import (
	"fmt"
	"initfunctions/helloworld"
)

func init() {
	//setting up database connection pooling
	fmt.Println("in init() of main.go 1")
}

func init() {
	//reading configuration files
	fmt.Println("in init() of main.go 2")
}

func init() {
	//initializing global variables
	fmt.Println("in init() of main.go 3")
}

func main() {
	fmt.Println("Hello World")
	helloworld.Run()
}
