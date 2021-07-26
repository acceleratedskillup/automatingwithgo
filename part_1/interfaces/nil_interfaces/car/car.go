package car

import "fmt"

type Car struct {
	Model string
}

func (cPtr *Car) Start() {
	/*
		if cPtr == nil {
			fmt.Println("cPtr is nil, please fix")
			return
		}
	*/
	fmt.Printf("car s-addr:%p is moving\n", cPtr)
}
