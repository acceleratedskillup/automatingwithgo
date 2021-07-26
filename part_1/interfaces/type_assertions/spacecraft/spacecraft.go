package spacecraft

import "fmt"

type Spacecraft struct {
	Model string
}

func (s Spacecraft) Start() {
	fmt.Printf("Spacecraft s-addr:%p has started\n", &s)
}

func (s Spacecraft) Fly() {
	fmt.Printf("Spacecraft s-addr:%p is flying\n", &s)
}
