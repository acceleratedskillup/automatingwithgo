package main

import "fmt"

func reverseSequence(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {
	sequence := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(reverseSequence(sequence))
}
