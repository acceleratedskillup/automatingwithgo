package main

import (
	"fmt"
	"math"
)

type TimeMachine struct {
	destinationYear int
	currentYear     int
	fuel            float64
}

func (tm *TimeMachine) setDestination(year int) {
	tm.destinationYear = year
}

func (tm *TimeMachine) calculateFuel() float64 {
	return math.Abs(float64(tm.destinationYear-tm.currentYear)) / 10.0 // Consumes 1 fuel unit for every 10 years
}

func (tm *TimeMachine) timeTravel() {
	requiredFuel := tm.calculateFuel()
	if tm.fuel < requiredFuel {
		fmt.Println("Not enough fuel!")
		return
	}
	tm.fuel -= requiredFuel
	tm.currentYear = tm.destinationYear
	fmt.Printf("Welcome to the year %d!\n", tm.currentYear)
}

func main() {
	machine := &TimeMachine{
		currentYear: 2023,
		fuel:        50.0,
	}
	machine.setDestination(3030)
	machine.timeTravel()
}
