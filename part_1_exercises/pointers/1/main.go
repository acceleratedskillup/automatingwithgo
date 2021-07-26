package main

import "fmt"

func main() {
	satellite1 := 42
	satellite2 := 55
	pSat1 := &satellite1
	pSat2 := &satellite2
	fmt.Println("Data from Satellite 1:", *pSat1)
	fmt.Println("Data from Satellite 2:", *pSat2)
}
