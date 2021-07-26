package main

import "fmt"

func main() {
	dataPoints := [3]int{10, 20, 30}
	pData := &dataPoints
	(*pData)[0] = 15
	fmt.Println("Updated Data Points:", dataPoints)
}
