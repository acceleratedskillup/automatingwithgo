package main

import "fmt"

func findRank(warriors []string, name string) int {
	for i, w := range warriors {
		if w == name {
			return i + 1
		}
	}
	return -1
}

func main() {
	warriors := []string{"Blade", "Arrow", "Spear", "Shield"}
	fmt.Println("Rank of Spear is:", findRank(warriors, "Spear"))
}
