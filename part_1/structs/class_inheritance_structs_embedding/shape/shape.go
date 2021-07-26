package shape

import "fmt"

type Shape struct {
	numberOfSides int
	color         int //lets just assume different values of int mean different colors
}

func (shape *Shape) GetNumberOfSides() int {
	return shape.numberOfSides
}

func (shape *Shape) SetNumberOfSides(sides int) {
	shape.numberOfSides = sides
}

func (shape *Shape) GetColor() int {
	return shape.color
}

func (shape *Shape) SetColor(color int) {
	shape.color = color
}

func (shape *Shape) Print() {
	fmt.Println("this is the Print() of Shape")
}
