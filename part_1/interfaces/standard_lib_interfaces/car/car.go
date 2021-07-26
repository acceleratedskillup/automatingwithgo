package car

import "fmt"

type Car struct {
	Model string
}

func (car *Car) Start() {
	fmt.Printf("car s-addr:%p has started\n", car)
}

func (car *Car) String() string {
	fmt.Println("String() is called")
	return fmt.Sprintf("car model is %s\n", car.Model)
}

func (car *Car) Write(data []byte) (n int, err error) {
	fmt.Println("Write() is called")
	car.Model = string(data)
	return len(data), nil
}
