package main

import "fmt"

type Tomb struct {
	currentChamber int
	chambers       map[int]string
}

func (t *Tomb) enterChamber(chamberNumber int) {
	riddle, exists := t.chambers[chamberNumber]
	if exists {
		fmt.Println("You've entered chamber", chamberNumber)
		fmt.Println("Riddle:", riddle)
	} else {
		fmt.Println("This chamber doesn't exist!")
	}
}

func (t *Tomb) solveRiddle(chamberNumber int, answer string) bool {
	if chamberNumber == 1 && answer == "time" {
		t.currentChamber++
		return true
	}
	// Additional riddles and answers can be added here
	return false
}

func main() {
	pharaohTomb := &Tomb{
		currentChamber: 0,
		chambers: map[int]string{
			1: "What flies without wings?",
			// More chambers and riddles can be added
		},
	}
	pharaohTomb.enterChamber(1)
	if pharaohTomb.solveRiddle(1, "time") {
		fmt.Println("Correct! You may proceed to the next chamber.")
	} else {
		fmt.Println("Incorrect! The chamber door remains closed.")
	}
}
