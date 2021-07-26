package main

import "fmt"

func oldestRelic(ages []int) int {
	oldest := ages[0]
	for _, age := range ages {
		if age > oldest {
			oldest = age
		}
	}
	return oldest
}

func main() {
	relicAges := []int{500, 600, 450, 700, 650}
	fmt.Println("The oldest relic is:", oldestRelic(relicAges), "years old.")
}
