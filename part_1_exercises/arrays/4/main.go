package main

import "fmt"

func sumAndAverage(arr [5]float64) (float64, float64) {
	sum := 0.0
	for _, num := range arr {
		sum += num
	}
	avg := sum / float64(len(arr))
	return sum, avg
}

func main() {
	numbers := [5]float64{10, 20, 30, 40, 50}
	s, a := sumAndAverage(numbers)
	fmt.Println("Sum:", s, "Average:", a)
}
