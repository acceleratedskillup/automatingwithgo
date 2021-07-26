package main

import "fmt"

func main() {

	var source = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenNumbers [5]int
	j := 0
	for _, num := range source {
		if num%2 == 0 {
			evenNumbers[j] = num
			j++
		}
	}
	fmt.Println(evenNumbers)

}
