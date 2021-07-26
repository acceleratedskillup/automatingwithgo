package main

import "fmt"

func PrintAnything(item interface{}) {
	fmt.Printf("Item: %v (Type: %T)\n", item, item)
}

func main() {
	PrintAnything("Hello, World!")
	PrintAnything(12345)
	PrintAnything(3.14159)
}
