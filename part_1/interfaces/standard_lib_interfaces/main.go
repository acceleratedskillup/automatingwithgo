package main

import (
	"fmt"
	"interfaces/car"
)

func main() {
	//example1()
	//example2()
	example3()
}

/*
execute this WITHOUT String()
*/
func example1() {
	car := new(car.Car)
	car.Model = "toyota"

	fmt.Printf("%s\n", car)
}

/*
execute this WITH String()
*/
func example2() {
	car := new(car.Car)
	car.Model = "toyota"

	fmt.Printf("%s\n", car)
}

func example3() {
	car := new(car.Car)
	car.Model = "toyota"

	fmt.Printf("%s\n", car)
	fmt.Println()
	fmt.Fprintf(car, "Tesla")
	fmt.Println()
	fmt.Printf("%s\n", car)
}
