package circle

//lint:ignore ST1001 this is safe
import (
	. "classinheritancestructsembedding/shape"
	"fmt"
)

/*
	Class inheritance by embedding
	take note  that using this approach,
	we are exposing all functions of Shape
	this is useful if circle inherits the exact behavior of Shape
*/
type Circle struct {
	Shape
	radius float64
}

func (circle *Circle) GetRadius() float64 {
	return circle.radius
}

func (circle *Circle) SetRadius(radius float64) {
	circle.radius = radius
}

func (circle *Circle) Print() {
	fmt.Println("this is the Print() of Circle")
}
