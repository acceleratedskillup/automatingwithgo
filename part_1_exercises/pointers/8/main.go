package main

import "fmt"

func main() {
	sequence := [5]int{10, 20, 30, 40, 50}
	pSeq := &sequence[0]
	fmt.Println("First Element:", *pSeq)
	*pSeq++
	fmt.Println("Second Element:", *pSeq)
}
