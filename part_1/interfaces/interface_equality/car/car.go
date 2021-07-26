package car

import "fmt"

type Car struct {
	Model string
}

func (c Car) Start() {
	fmt.Printf("car s-addr:%p has started\n", &c)
}
