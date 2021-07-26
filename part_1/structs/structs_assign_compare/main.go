package main

import (
	"fmt"
	"reflect"
)

func main() {

	//example1()
	//example2()
	//example3()
	example4()
}

//assignment is making a copy
func example1() {

	type Building struct {
		height float64
		name   string
	}

	buildingA := Building{
		height: 150,
		name:   "Building A",
	}

	buildingB := buildingA
	buildingB.name = "Building B"

	fmt.Printf("buildingA:%#v\n", buildingA)
	fmt.Printf("address buildingA:%p\n", &buildingA)
	fmt.Printf("buildingB:%#v\n", buildingB)
	fmt.Printf("address buildingB:%p\n", &buildingB)
}

//using the equality operator
func example2() {

	type Building struct {
		height float64
		name   string
	}

	type House struct {
		height float64
		name   string
	}

	buildingA := Building{
		height: 150,
		name:   "Building A",
	}

	/*
		buildingB := House{
			height: 150,
			name:   "Building A",
		}
	*/

	buildingB := buildingA
	buildingB.name = "Building B"

	//alternatively
	fmt.Println("equal?", (buildingA == buildingB))

	if (buildingA.height == buildingB.height) &&
		(buildingA.name == buildingB.name) {
		fmt.Println("they are equal")
	} else {
		fmt.Println("they are not equal")
	}
}

//what if we are using slices
func example3() {
	type Building struct {
		height      float64
		name        string
		unitNumbers []string
	}

	buildingA := Building{
		height:      150,
		name:        "Building A",
		unitNumbers: []string{"#1-01", "#1-02"},
	}

	buildingB := buildingA
	//buildingB.unitNumbers[0] = "#2-01"
	//example 3.1
	buildingB.unitNumbers = []string{"#2-01", "#2-02"}

	fmt.Printf("buildingA:%#v\n", buildingA)
	fmt.Printf("buildingB:%#v\n", buildingB)

	//example 3.2
	//fmt.Println("equal?", (buildingA == buildingB))

	if buildingA.height == buildingB.height &&
		buildingA.name == buildingB.name &&
		len(buildingA.unitNumbers) == len(buildingB.unitNumbers) {
		unitNumbersAreSame := true
		for index, unitNumber := range buildingA.unitNumbers {
			if buildingB.unitNumbers[index] != unitNumber {
				unitNumbersAreSame = false
				break
			}
		}
		if unitNumbersAreSame {
			fmt.Println("they are equal")
		} else {
			fmt.Println("they are not equal")
		}
	} else {
		fmt.Println("they are not equal")
	}
}

//use of reflect.DeepEqual()
func example4() {
	type Building struct {
		height      float64
		name        string
		unitNumbers []string
	}

	buildingA := Building{
		height:      150,
		name:        "Building A",
		unitNumbers: []string{"#1-01", "#1-02"},
	}

	buildingB := buildingA
	buildingB.unitNumbers[0] = "#2-01"
	//example 3.1
	//buildingB.unitNumbers = []string{"#2-01", "#2-02"}

	fmt.Printf("buildingA:%#v\n", buildingA)
	fmt.Printf("buildingB:%#v\n", buildingB)
	fmt.Println("equal?", reflect.DeepEqual(buildingA, buildingB))

}
