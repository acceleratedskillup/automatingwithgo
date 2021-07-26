package main

import "fmt"

func arraysAreEqual(a, b [5]int) bool {
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arraysAreEqual(arr1, arr2))
}
