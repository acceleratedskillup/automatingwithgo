package main

import "fmt"

func main() {
	var matrix = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var transposed [3][3]int
	for i, row := range matrix {
		for j, val := range row {
			transposed[j][i] = val
		}
	}
	for _, row := range transposed {
		fmt.Println(row)
	}

}
