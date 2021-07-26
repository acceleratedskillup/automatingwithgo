package main

import "fmt"

func main() {
	var numbers = [7]int{1, 2, 3, 4, 5, 6, 7}
	middle := numbers[2:5]
	fmt.Println(middle)

}
