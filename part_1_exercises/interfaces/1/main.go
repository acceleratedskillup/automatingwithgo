package main

import "fmt"

type SpaceVehicle interface {
	Travel()
}

type Rocket struct{}
type Shuttle struct{}

func (r Rocket) Travel() {
	fmt.Println("Rocket is traveling through space!")
}

func (s Shuttle) Travel() {
	fmt.Println("Shuttle is orbiting the planet!")
}

func main() {
	var vehicle1 SpaceVehicle = Rocket{}
	var vehicle2 SpaceVehicle = Shuttle{}

	vehicle1.Travel()
	vehicle2.Travel()
}
