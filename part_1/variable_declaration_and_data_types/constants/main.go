package main

import "fmt"

//you can declare constants at package level
//defining a single typed constant
const MinuteInSeconds int = 60

func main() {

	//defining a single constant with an expression
	const HourInSeconds = MinuteInSeconds * 60

	const DayInSeconds = HourInSeconds * 24

	//another way to declare or group related constants
	/*
		const (
			Red   = 0
			Blue  = 1
			Green = 2
		)
	*/
	//using iota to initialize constants, iota starts at 0, increments after that
	const (
		Red   = iota // 0
		Blue         //1
		Green        //2
	)

	fmt.Println("Red:", Red, "Blue:", Blue, "Green:", Green)

	//iota resets to 0 when it encounters a const keyword
	const (
		Orange = iota + 1 //1
		_                 //skipped
		Purple            // 3
		Black             //4
	)

	fmt.Println("Orange:", Orange, "Purple:", Purple, "Black:", Black)
}
