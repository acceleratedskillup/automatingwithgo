package main

import "fmt"

type Dragon struct {
	strength       int
	treasureAmount int
	mood           string
}

func (d *Dragon) approach() {
	if d.mood == "aggressive" {
		fmt.Println("The dragon roars fiercely!")
	} else {
		fmt.Println("The dragon seems calm.")
	}
}

func (d *Dragon) takeTreasure() {
	if d.strength > 50 {
		fmt.Println("The dragon defends its treasure and you flee!")
	} else {
		d.treasureAmount -= 10
		fmt.Println("You manage to steal some of the dragon's treasure!")
	}
}

func (d *Dragon) befriend() {
	if d.mood == "calm" {
		fmt.Println("The dragon seems to accept your gesture of friendship!")
	} else {
		fmt.Println("The dragon doesn't trust you!")
	}
}

func main() {
	drogon := &Dragon{
		strength:       70,
		treasureAmount: 100,
		mood:           "aggressive",
	}
	drogon.approach()
	drogon.takeTreasure()
}
