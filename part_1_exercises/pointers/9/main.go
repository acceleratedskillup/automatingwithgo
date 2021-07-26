package main

import "fmt"

func averageTemperature(readings *[5]float64) float64 {
	sum := 0.0
	for _, temp := range readings {
		sum += temp
	}
	return sum / float64(len(readings))
}

func main() {
	temperatures := [5]float64{20.5, 21.3, 19.8, 22.1, 20.9}
	avgTemp := averageTemperature(&temperatures)
	fmt.Println("Average Temperature:", avgTemp)
}
