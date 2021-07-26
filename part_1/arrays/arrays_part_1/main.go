package main

import "fmt"

const NUM_OF_WEIGHTS = 3

func main() {

	var weights [3]byte
	//var weightsLongerArray [10]byte
	fmt.Printf("weight array: %v\n", weights)

	var cities [3]string
	fmt.Printf("cities array: %q\n", cities)

	weightIndex := 0
	fmt.Printf("weights[0]: %v\n", weights[weightIndex])
	weights[weightIndex] = 10
	fmt.Printf("length of weights: %d\n", len(weights))

}
