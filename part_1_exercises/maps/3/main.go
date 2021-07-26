package main

import "fmt"

func main() {
	attendees := map[string]bool{
		"Alice": true,
		"Bob":   true,
	}

	delete(attendees, "Alice")
	if _, exists := attendees["Alice"]; !exists {
		fmt.Println("Alice was removed from the attendee list!")
	} else {
		fmt.Println("Alice is still in the attendee list!")
	}
}
