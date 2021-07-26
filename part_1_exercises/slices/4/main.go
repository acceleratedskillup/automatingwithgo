package main

import "fmt"

func averageTemperature(temps []float64) float64 {
	total := 0.0
	for _, temp := range temps {
		total += temp
	}
	return total / float64(len(temps))
}

func main() {
	temperatures := []float64{-5.2, -3.8, -4.6, -2.9, -3.4, -4.1, -5.0}
	fmt.Println("The average temperature is:", averageTemperature(temperatures))
}
