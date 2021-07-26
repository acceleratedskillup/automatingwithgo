package main

import (
	"fmt"
	"interfaces/car"
	"interfaces/spacecraft"
)

type Vehicle interface {
	Start()
}

func main() {

	//example1()
	//example2()
	//example3()
	example4()
}

/*
comparing 2 interface variables with each other,
when types and values are same
*/

func example1() {
	var vehicle1 Vehicle = car.Car{Model: "toyota"}
	var vehicle2 Vehicle = car.Car{Model: "toyota"}
	fmt.Printf("vehicle1==vehicle2: %t\n", vehicle1 == vehicle2)
}

/*
comparing 2 interface variables with each other,
when types are same but values are different
*/
func example2() {
	var vehicle1 Vehicle = car.Car{Model: "toyota"}
	var vehicle2 Vehicle = car.Car{Model: "honda"}
	fmt.Printf("vehicle1==vehicle2: %t\n", vehicle1 == vehicle2)
}

/*
comparing 2 interface variables with each other,
when types are different but values are same(nil)
*/
func example3() {
	var carPtr *car.Car
	var vehicle1 Vehicle = carPtr

	var spacecraftPtr *spacecraft.Spacecraft
	var vehicle2 Vehicle = spacecraftPtr

	fmt.Printf("vehicle1: type:%T value:%v\n", vehicle1, vehicle1)
	fmt.Printf("vehicle2: type:%T value:%v\n", vehicle2, vehicle2)
	fmt.Printf("vehicle1==vehicle2: %t\n", vehicle1 == vehicle2)
}

/*
comparing interface value with non interface value
*/

func example4() {

	var carVehicle Vehicle = car.Car{Model: "toyota"}
	car := car.Car{Model: "toyota"}

	fmt.Printf("carVehicle: type:%T value:%+v\n", carVehicle, carVehicle)
	fmt.Printf("car: type:%T value:%+v\n", car, car)
	fmt.Printf("carVehicle==car: %t\n", carVehicle == car)
}
