package main

import "fmt"

func main() {
	studentGrades := map[string]string{
		"Alice":   "A",
		"Bob":     "B",
		"Charlie": "C",
		"David":   "D",
		"Eve":     "E",
	}

	for student, grade := range studentGrades {
		fmt.Println(student, ":", grade)
	}
}
