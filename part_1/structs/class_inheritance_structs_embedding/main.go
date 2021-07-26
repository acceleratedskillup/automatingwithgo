package main

import (
	//lint:ignore ST1001 this is safe
	. "classinheritancestructsembedding/circle"
	//lint:ignore ST1001 this is safe
	. "classinheritancestructsembedding/rectangle"
	"fmt"
)

func main() {

	rectangle := Rectangle{}
	rectangle.SetNumberOfSides(4)
	rectangle.SetWidth(10)
	rectangle.SetHeight(20)
	//rectangle.setColor() //function not available

	fmt.Printf("%+v\n", rectangle)

	circle := Circle{}
	circle.SetNumberOfSides(1)
	circle.SetRadius(1.23)
	circle.SetColor(1)
	fmt.Printf("%+v\n", circle)
	circle.Print()
	circle.Shape.Print()

}
