package main

import "fmt"

type Coordinates struct {
	X, Y, Z float64
}

func updateCoordinates(coords *Coordinates) {
	coords.X = 5.5
	coords.Y = 6.6
	coords.Z = 7.7
}

func main() {
	roverCoords := Coordinates{1.1, 2.2, 3.3}
	updateCoordinates(&roverCoords)
	fmt.Println("Updated Rover Coordinates:", roverCoords)
}
