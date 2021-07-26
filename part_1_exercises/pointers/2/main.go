package main

import "fmt"

type Astronaut struct {
	Name string
	Age  int
}

func main() {
	astronautPointer := &Astronaut{}
	astronautPointer.Name = "John Doe"
	astronautPointer.Age = 32
	fmt.Println("Astronaut Details:", *astronautPointer)
}
