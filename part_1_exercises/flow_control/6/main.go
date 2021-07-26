package main

import "fmt"

func causePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("This will cause a panic")
	panic("Panic simulation")
}

func main() {
	causePanic()
	fmt.Println("Program continues after panic")
}
