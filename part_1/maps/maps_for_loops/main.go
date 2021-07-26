package main

import "fmt"

func main() {
	//example1()
	example2()
}

func example1() {
	nameToScoreMap := map[string]int{
		"john":  64,
		"may":   84,
		"peter": 70,
	}
	for name, score := range nameToScoreMap {
		fmt.Printf("name:%s score:%d\n", name, score)
	}
	fmt.Println("-------------------------")
	for name, score := range nameToScoreMap {
		fmt.Printf("name:%s score:%d\n", name, score)
	}
	fmt.Println("-------------------------")
	for name, score := range nameToScoreMap {
		fmt.Printf("name:%s score:%d\n", name, score)
	}
}

func example2() {
	nameToScoreMap := map[string]int{
		"john":  64,
		"may":   84,
		"peter": 70,
	}

	fmt.Println("before:", nameToScoreMap)

	for name := range nameToScoreMap {
		delete(nameToScoreMap, name)
	}

	fmt.Println("after:", nameToScoreMap)
}
