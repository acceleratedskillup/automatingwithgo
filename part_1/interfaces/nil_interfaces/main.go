package main

import (
	"fmt"
	//lint:ignore ST1001 this is safe
	car "interfaces/car"
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

func example1() {
	car := car.Car{
		Model: "toyota",
	}

	var vehicle Vehicle = &car
	fmt.Printf("Concrete type of interface value: %T\n", vehicle)
	fmt.Printf("Concrete value of interface value: %+v\n", vehicle)
	fmt.Printf("Car s-addr:%p\n", &car)
	vehicle.Start()
}

func example2() {

	var vehicle Vehicle
	fmt.Printf("Concrete type of interface value: %T\n", vehicle)
	fmt.Printf("Concrete value of interface value: %+v\n", vehicle)
	fmt.Println("is vehicle nil?", (vehicle == nil))

}

func example3() {

	var vehicle Vehicle
	fmt.Printf("Concrete type of interface value: %T\n", vehicle)
	fmt.Printf("Concrete value of interface value: %+v\n", vehicle)
	fmt.Println("is vehicle nil?", (vehicle == nil))
	//will throw error
	vehicle.Start()
}

func example4() {
	var car *car.Car

	var vehicle Vehicle = car
	fmt.Printf("Concrete type of interface value: %T\n", vehicle)
	fmt.Printf("Concrete value of interface value: %+v\n", vehicle)
	fmt.Println("is vehicle nil?", (vehicle == nil))
	vehicle.Start()
}
