package main

import "fmt"

func main() {
	var identity [4][4]int
	for i := 0; i < 4; i++ {
		identity[i][i] = 1
	}
	for _, row := range identity {
		fmt.Println(row)
	}

}
