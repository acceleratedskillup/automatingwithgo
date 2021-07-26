package main

import (
	"fmt"
	//lint:ignore ST1001 this is safe
	car "interfaces/car"
)

type Vehicle interface {
	Start()
	Stop()
}

func main() {
	//example1()
	example2()
	//example3()
}

/*
example 1, assigning a car struct (pointer receiver)
to an interface var
*/
func example1() {
	car := car.Car{
		Model: "toyota",
	}
	//line below will have compile error
	//var vehicle Vehicle = car
	var vehicle Vehicle = &car
	fmt.Printf("Concrete type of interface value: %T\n", vehicle)
	fmt.Printf("Concrete value of interface value: %+v\n", vehicle)
	fmt.Printf("Car s-addr:%p\n", &car)
	vehicle.Start()
}

/*
example 3, assigning a car struct
(combination of pointer + value receivers)
to an interface var

Todo:need to change Stop method to use value receivers
*/
func example2() {
	car := car.Car{
		Model: "toyota",
	}
	var vehicle Vehicle = &car
	fmt.Printf("Concrete type of interface value: %T\n", vehicle)
	fmt.Printf("Concrete value of interface value: %+v\n", vehicle)
	fmt.Printf("Car s-addr:%p\n", &car)
	vehicle.Start()
	vehicle.Stop()
}

/*
example 3, assigning a car struct pointer
to an interface
*/
func example3() {
	car := new(car.Car)
	var vehicle Vehicle = car
	fmt.Printf("car: %T %+v %p\n", car, car, car)
	fmt.Printf("vehicle: %T %+v %p\n", vehicle, vehicle, &vehicle)
	vehicle.Start()
	vehicle.Stop()
}
