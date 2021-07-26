package main

import "fmt"

func main() {
	signalStrength := 78
	pSignal := &signalStrength
	*pSignal = 85
	fmt.Println("Updated Signal Strength:", signalStrength)
}
