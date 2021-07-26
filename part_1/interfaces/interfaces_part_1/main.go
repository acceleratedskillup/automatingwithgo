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
	example1()
	//example2()
	//example3()
}

/*
assigning a car struct (value receiver)
to an interface var
*/
func example1() {
	car := car.Car{
		Model: "toyota",
	}
	var vehicle Vehicle = car
	fmt.Printf("Concrete type of interface value: %T\n", vehicle)
	fmt.Printf("Concrete value of interface value: %+v\n", vehicle)
	fmt.Printf("Car s-addr:%p\n", &car)
	vehicle.Start()
}

/*
assigning a spacecraft struct (value receiver)
to an interface var

showing same interface but with different implementation
*/
func example2() {
	spacecraft := spacecraft.Spacecraft{
		Model: "USS Enterprise",
	}

	var vehicle Vehicle = spacecraft
	fmt.Printf("Concrete type of interface value: %T\n", vehicle)
	fmt.Printf("Concrete value of interface value: %+v\n", vehicle)
	fmt.Printf("Spacecraft s-addr:%p\n", &spacecraft)
	vehicle.Start()
}

/*
shows multiple interfaces for 1 object
*/

func example3() {

	spacecraft := spacecraft.Spacecraft{
		Model: "USS Enterprise",
	}

	var vehicle Vehicle = spacecraft
	var aircraft Aircraft = spacecraft
	fmt.Printf("Concrete type of Vehicle interface value: %T\n", vehicle)
	fmt.Printf("Concrete value of Vehicle interface value: %+v\n", vehicle)
	fmt.Printf("Spacecraft s-addr:%p\n", &spacecraft)
	vehicle.Start()
	fmt.Printf("Concrete type of Aircraft interface value: %T\n", aircraft)
	fmt.Printf("Concrete value of Aircraft interface value: %+v\n", aircraft)
	fmt.Printf("Spacecraft s-addr:%p\n", &spacecraft)
	aircraft.Fly()

}
