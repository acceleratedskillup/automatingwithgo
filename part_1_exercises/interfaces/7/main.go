package main

import "fmt"

type Spacecraft interface{}

type Rocket struct{}
type Satellite struct{}

func Journey(s Spacecraft) {
	switch craft := s.(type) {
	case Rocket:
		fmt.Println("Rocket is launching to space!")
	case Satellite:
		fmt.Println("Satellite is orbiting the Earth!")
	default:
		fmt.Printf("Unknown spacecraft: %T\n", craft)
	}
}

func main() {
	Journey(Rocket{})
	Journey(Satellite{})
}
