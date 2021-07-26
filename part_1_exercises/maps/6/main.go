package main

import "fmt"

func main() {
	inventory := map[string]int{
		"Laptop": 5,
		"Mouse":  10,
	}

	product := "Laptop"
	if _, exists := inventory[product]; exists {
		fmt.Println(product, "is available!")
	} else {
		fmt.Println(product, "is not available!")
	}
}
