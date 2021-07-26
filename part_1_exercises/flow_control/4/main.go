package main

import (
	"fmt"
)

func loopingCalculator(numbers []float64, operations []string) ([]float64, string) {
	message := ""
ExitLoop:
	for _, operation := range operations {
		switch operation {
		case "add":
			for i := range numbers {
				numbers[i] += numbers[i]
			}
		case "subtract":
			for i := range numbers {
				numbers[i] -= numbers[i]
			}
		case "multiply":
			for i := range numbers {
				numbers[i] *= numbers[i]
			}
		case "divide":
			for i := range numbers {
				if numbers[i] == 0 {
					message = "Division by zero detected. Skipping."
					continue
				}
				numbers[i] /= numbers[i]
			}
		case "exit":
			message = "Calculator exited"
			break ExitLoop
		default:
			message = "Unknown operation detected. Skipping."
		}
	}
	return numbers, message
}

func main() {
	numbers := []float64{1, 2, 3, 4}
	operations := []string{"add", "subtract", "exit", "multiply"}
	result, msg := loopingCalculator(numbers, operations)
	fmt.Println("Result:", result)
	fmt.Println("Message:", msg)
}
