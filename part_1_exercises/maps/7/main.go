package main

import "fmt"

func changeSeat(seating map[string]int, student string, newSeat int) {
	seating[student] = newSeat
}

func main() {
	seatingArrangement := map[string]int{
		"Alice": 1,
		"Bob":   2,
		"Eve":   3,
	}

	fmt.Println("Original Seating:", seatingArrangement)

	changeSeat(seatingArrangement, "Alice", 4)
	fmt.Println("Updated Seating:", seatingArrangement)
}
