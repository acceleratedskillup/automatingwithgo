package main

import "fmt"

func main() {

	categories := [3]string{"a", "b", "c"}
	for i := 0; i < len(categories); i++ {
		fmt.Println("category:", categories[i])
	}

	for index, value := range categories {
		//value = "x"
		//fmt.Printf("%d) %v\n", index, value)
		fmt.Printf("%d) value:%v category:%v\n", index, value, categories[index])

	}
}
