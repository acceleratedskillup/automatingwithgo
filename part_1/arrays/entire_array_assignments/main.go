package main

import "fmt"

func main() {

	categories := [3]string{"a", "b", "c"}
	codes := [3]string{}
	//codes := categories
	//codes = categories
	for index, value := range categories {
		codes[index] = value
	}
	fmt.Println("categories:", categories)
	fmt.Println("codes:", codes)

	codes[2] = "x"
	fmt.Println("categories:", categories)
	fmt.Println("codes:", codes)

	//arrays are of different length
	//if codes is larger
	for index, value := range categories {
		codes[index] = value
	}
	//if codes is smaller
	for i := 0; i < len(codes); i++ {
		codes[i] = categories[i]
	}
}
