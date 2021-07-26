package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	// Top-left corner of the rectangle
	x1 = 1000
	y1 = 500
)

func main() {
	// Set the number of seconds for the delay
	var mouseMoveInterval int
	fmt.Print("Enter the number of seconds for each mouse move: ")
	_, err := fmt.Scanln(&mouseMoveInterval)
	if err != nil {
		fmt.Println("Error reading from cli")
		return
	}

	// Get the screen size
	screenWidth, screenHeight := robotgo.GetScreenSize()

	// Define the confined area within the screen
	// Example: A rectangle in the upper-left quarter of the screen

	fmt.Println("screenWidth", screenWidth, "screenHeight:", screenHeight)

	// Seed the random number generator
	seed := time.Now().UnixNano()
	newRand := rand.New(rand.NewSource(seed))

	for {
		// Generate random X and Y coordinates within the confined area

		/*
			if you are using the full screen for
			mouse moves
		*/
		//randX := newRand.Intn(screenWidth)
		//randY := newRand.Intn(screenWidth)

		/*
			if you want to constrain the mouse moves
			to a certain area on the screen
		*/
		randX := x1 + newRand.Intn(screenWidth-x1+1)
		randY := y1 + newRand.Intn(screenHeight-y1+1)

		// Move the mouse to the random position
		robotgo.Move(randX, randY)

		// Wait for 'x' seconds
		time.Sleep(time.Duration(mouseMoveInterval) * time.Second)
	}
}
