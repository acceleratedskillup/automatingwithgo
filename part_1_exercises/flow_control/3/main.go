package main

import "fmt"

func dayCategory(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "weekend"
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "weekday"
	default:
		return "invalid day"
	}
}

func main() {
	fmt.Println(dayCategory("Monday"))   // Expected output: weekday
	fmt.Println(dayCategory("Saturday")) // Expected output: weekend
}
