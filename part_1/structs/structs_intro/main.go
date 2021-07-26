package main

import "fmt"

type Car struct {
	model         string
	plateNumber   string
	numberOfSeats int
}

func main() {

	example1()
	//example2()
	//example3()
	//example4()
	//example5()
	//example6()
	//example7()
	//example8()

}

func example1() {
	models := []string{}
	plateNumbers := []string{}
	numberOfSeats := []int{}

	models = append(models, "Toyota")
	plateNumbers = append(plateNumbers, "S1254")
	numberOfSeats = append(numberOfSeats, 4)

	selectedCarIndex := 0
	fmt.Printf("model: %v\n", models[selectedCarIndex])
	fmt.Printf("plateNumber: %v\n", plateNumbers[selectedCarIndex])
	fmt.Printf("numberOfSeats: %v\n", numberOfSeats[selectedCarIndex])

	numberOfSeats = append(numberOfSeats, 4)

}
func example2() {
	var toyotaCar Car
	toyotaCar.model = "Prius"
	toyotaCar.plateNumber = "S1254"
	toyotaCar.numberOfSeats = 4
	fmt.Printf("toyotaCar: %#v\n", toyotaCar)
}

func example3() {
	toyotaCar := Car{}
	fmt.Printf("before init toyotaCar: %#v\n", toyotaCar)
	toyotaCar.model = "Prius"
	toyotaCar.plateNumber = "S1254"
	toyotaCar.numberOfSeats = 4
	fmt.Printf("after init toyotaCar: %#v\n", toyotaCar)
}
func example4() {
	/*
		we can create and initialize a struct with a struct literal.
		order of the fields are not irrelevant
	*/
	toyotaCar :=
		Car{
			model:         "Prius",
			plateNumber:   "S1254",
			numberOfSeats: 4,
		}
	/*
		when printing structs, the # flag (%#v) adds field names+struct name
	*/
	fmt.Printf("toyotaCar: %#v\n", toyotaCar)
}

func example5() {
	/*
		another way to create and initialize a struct, all fields must be specified
		in the order in which the fields are declared.
	*/
	toyotaCar :=
		Car{
			"Prius",
			"S1254",
			4,
		}
	fmt.Printf("toyotaCar: %#v\n", toyotaCar)
}

func example6() {
	// another way to create and initialize a struct
	toyotaCar :=
		Car{
			model:       "Prius",
			plateNumber: "S1254",
		}
	fmt.Printf("toyotaCar: %#v\n", toyotaCar)
}

func example7() {
	//anonymous struct
	var superCar = struct {
		engineType string
	}{engineType: "V8"}
	fmt.Printf("superCar: %#v\n", superCar)
}

func example8() {
	//anonymous fields struct

	type SuperCar struct {
		string
		int
		bool
	}
	ferrariCar := SuperCar{"ferrari", 4, true}
	fmt.Printf("model: %v\n", ferrariCar.string)
	fmt.Printf("number of seats: %v\n", ferrariCar.int)
	fmt.Printf("has spoiler: %v\n", ferrariCar.bool)

	fmt.Printf("superCar: %#v\n", ferrariCar)
}
