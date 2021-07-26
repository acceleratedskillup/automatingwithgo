package main

import "fmt"

func filterEvenPetals(flowers []int) []int {
	var evenPetals []int
	for _, petals := range flowers {
		if petals%2 == 0 {
			evenPetals = append(evenPetals, petals)
		}
	}
	return evenPetals
}

func main() {
	flowers := []int{5, 6, 7, 8, 9, 10}
	fmt.Println("Flowers with even-numbered petals:", filterEvenPetals(flowers))
}
