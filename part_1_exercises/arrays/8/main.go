package main

import "fmt"

func main() {
	type Point [2]float64
	var points = [3]Point{{1.2, 2.3}, {3.4, 4.5}, {5.6, 6.7}}
	for _, p := range points {
		fmt.Println("X:", p[0], "Y:", p[1])
	}

}
