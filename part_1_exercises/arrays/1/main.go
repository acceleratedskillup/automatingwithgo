package main

import "fmt"

func main() {
	var days = [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	var reversed [5]string
	for i, day := range days {
		reversed[4-i] = day
	}
	fmt.Println(reversed)

}
