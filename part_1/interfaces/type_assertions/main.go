package main

import (
	"fmt"
	"interfaces/car"
	"interfaces/spacecraft"
)

type Vehicle interface {
	Start()
}

type Aircraft interface {
	Fly()
}

func main() {

	//example1()
	//example2()
	//example3()
	//example4()
	//example4_1()
	example4_2()
	//example5()
	//example6()
	//example7()
}

/*
can only call interface method
so need type assertion
*/
func example1() {
	var vehicle Vehicle = spacecraft.Spacecraft{
		Model: "USS Enterprise",
	}
	v := vehicle.(spacecraft.Spacecraft)
	fmt.Printf("v type:%T value:%+v\n", v, v)
}

/*
try to use a wrong type to convert, which results in runtime
error
*/
func example2() {

	var vehicle Vehicle = spacecraft.Spacecraft{
		Model: "USS Enterprise",
	}

	v := vehicle.(car.Car)
	fmt.Printf("v type:%T value:%+v\n", v, v)
}

/*
use the boolean ok variable to check if type assertion is ok
*/
func example3() {
	var vehicle Vehicle = spacecraft.Spacecraft{
		Model: "USS Enterprise",
	}

	//try switching to spacecraft.Spacecraft for positive example
	v, assertOk := vehicle.(car.Car)
	fmt.Printf("v type:%T value:%+v assertOk:%v\n", v, v, assertOk)
}

/*
use the if else statement with the boolean ok variable
*/
func example4() {
	var vehicle Vehicle = spacecraft.Spacecraft{
		Model: "USS Enterprise",
	}
	/*
		change to interface type to show we can also use that
		besides concrete type

		call v.Fly() after the printf()
	*/
	if v, assertOk := vehicle.(spacecraft.Spacecraft); assertOk {
		fmt.Printf("v type:%T value:%+v assertOk:%v\n", v, v, assertOk)
		//v.Fly()
	} else {
		fmt.Println("unknown type")
	}
}

/*
change to interface type to show we can also use that
besides concrete type
*/
func example4_1() {
	var vehicle Vehicle = spacecraft.Spacecraft{
		Model: "USS Enterprise",
	}
	if v, assertOk := vehicle.(Aircraft); assertOk {
		fmt.Printf("v type:%T value:%+v assertOk:%v\n", v, v, assertOk)
		v.Fly()
	} else {
		fmt.Println("unknown type")
	}
}

/*
using embedded interface type
*/
func example4_2() {
	var vehicle Vehicle = spacecraft.Spacecraft{
		Model: "USS Enterprise",
	}
	if v, assertOk := vehicle.(interface{ Fly() }); assertOk {
		fmt.Printf("v type:%T value:%+v assertOk:%v\n", v, v, assertOk)
		v.Fly()
	} else {
		fmt.Println("unknown type")
	}
}

/*
using multiple if else statements
*/
func example5() {
	var vehicle Vehicle = spacecraft.Spacecraft{
		Model: "USS Enterprise",
	}
	printVehicle(vehicle)

	car := car.Car{
		Model: "toyota",
	}
	printVehicle(car)
}

/*
using switch statement
*/
func example6() {
	var vehicle Vehicle = spacecraft.Spacecraft{
		Model: "USS Enterprise",
	}
	printVehicleUsingSwitch(vehicle)

	car := car.Car{
		Model: "toyota",
	}
	printVehicleUsingSwitch(car)
}

func printVehicle(vehicle Vehicle) {

	if v, assertOk := vehicle.(car.Car); assertOk {
		fmt.Printf("car.Car: type:%T value:%+v\n", v, v)
	} else if v, assertOk := vehicle.(Aircraft); assertOk {
		fmt.Printf("Aircraft: type:%T value:%+v\n", v, v)
	} else {
		fmt.Println("no such vehicle type")
	}
}

func printVehicleUsingSwitch(vehicle Vehicle) {
	switch value := vehicle.(type) {
	case car.Car:
		fmt.Printf("value type:%T value:%+v\n", value, value)
	case Aircraft:
		fmt.Printf("value type:%T value:%+v\n", value, value)
	default:
		fmt.Println("no such vehicle type")
	}
}
