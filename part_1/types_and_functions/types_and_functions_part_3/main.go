package main

import "fmt"

type Car struct {
	Model         string
	PlateNumber   string
	HasAircon     bool
	NumberOfSeats int
}

func main() {

	//example 1
	toyotaCar :=
		Car{
			Model:         "Prius",
			PlateNumber:   "S1254",
			HasAircon:     true,
			NumberOfSeats: 4,
		}

	toyotaCar.Print()
	fmt.Printf("car:%p %+v\n", &toyotaCar, toyotaCar)
	toyotaCar.PrintValue()
	fmt.Printf("car:%+v\n", toyotaCar)

	//example 2 just to show we can use the same functions with pointers
	teslaCar := new(Car)
	teslaCar.Model = "Tesla123"
	teslaCar.PlateNumber = "T100"
	teslaCar.HasAircon = true
	teslaCar.NumberOfSeats = 4

	teslaCar.Print()
	fmt.Printf("teslaCar: %+v\n", *teslaCar)
	teslaCar.PrintValue()
	fmt.Printf("teslaCar: %+v\n", *teslaCar)
}

func (inputCar Car) Print() {
	inputCar.NumberOfSeats = 3
	fmt.Printf("Print() car:%+v\n", inputCar)
}

func (inputCar *Car) PrintValue() {
	inputCar.NumberOfSeats = 2
	fmt.Printf("PrintValue() %p car:%+v\n", inputCar, *inputCar)
}
