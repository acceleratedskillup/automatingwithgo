package car

import "fmt"

type Car struct {
	Model string
}

func (cPtr *Car) Start() {
	fmt.Printf("car s-addr:%p is moving\n", cPtr)
}

func (c Car) Stop() {
	fmt.Printf("car s-addr:%p has stopped\n", &c)
}
