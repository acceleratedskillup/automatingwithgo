package main

import "fmt"

func mapsEqual(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}
	return true
}

func main() {
	surveyCityA := map[string]int{"Python": 50, "Go": 30}
	surveyCityB := map[string]int{"Python": 50, "Go": 30}
	fmt.Println(mapsEqual(surveyCityA, surveyCityB)) // true
}
