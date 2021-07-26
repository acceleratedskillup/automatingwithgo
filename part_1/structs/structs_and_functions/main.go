package main

import "fmt"

type Car struct {
	plateNumber   string
	numberOfSeats int
}

func main() {
	//example1()
	example2()
}

func example1() {
	toyotaCar :=
		Car{
			plateNumber:   "S1254",
			numberOfSeats: 4,
		}
	fmt.Printf("toyotaCar address:%p\n", &toyotaCar)
	fmt.Printf("toyotaCar before:%#v\n", toyotaCar)
	modifyCar(toyotaCar)
	fmt.Printf("toyotaCar address:%p\n", &toyotaCar)
	fmt.Printf("toyotaCar after:%#v\n", toyotaCar)
}

func example2() {
	toyotaCar :=
		Car{
			plateNumber:   "S1254",
			numberOfSeats: 4,
		}
	fmt.Printf("toyotaCar address:%p\n", &toyotaCar)
	fmt.Printf("toyotaCar before:%#v\n", toyotaCar)
	toyotaCar = modifyCarWithReturn(toyotaCar)
	fmt.Printf("toyotaCar address:%p\n", &toyotaCar)
	fmt.Printf("toyotaCar after:%#v\n", toyotaCar)
}

// function arguments are passed by value
func modifyCar(inputCar Car) {
	fmt.Printf("inputCar address:%p\n", &inputCar)
	inputCar.numberOfSeats = 1
	fmt.Printf("inputCar: %#v\n", inputCar)
}

func modifyCarWithReturn(inputCar Car) Car {
	fmt.Printf("inputCar address:%p\n", &inputCar)
	inputCar.numberOfSeats = 1
	fmt.Printf("inputCar: %#v\n", inputCar)
	return inputCar
}
