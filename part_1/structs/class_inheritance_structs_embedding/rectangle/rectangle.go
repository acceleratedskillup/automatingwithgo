package rectangle

//lint:ignore ST1001 this is safe
import . "classinheritancestructsembedding/shape"

//Class inheritance by composition
type Rectangle struct {
	/*
		nested struct, inheritance by composition
		take note also that using this approach,
		we can control which functions of Shape to expose to the user
		this is useful if rectangle inherits some of the behavior of Shape
	*/
	shape  Shape
	width  float64
	height float64
}

//of course, if we expose shape by capitalizing as Shape, there is no need for this function
func (rectangle *Rectangle) GetNumberOfSides() int {
	return rectangle.shape.GetNumberOfSides()
}

func (rectangle *Rectangle) SetNumberOfSides(sides int) {
	rectangle.shape.SetNumberOfSides(sides)
}

func (rectangle *Rectangle) GetWidth() float64 {
	return rectangle.width
}

func (rectangle *Rectangle) SetWidth(width float64) {
	rectangle.width = width
}

func (rectangle *Rectangle) GetHeight() float64 {
	return rectangle.height
}

func (rectangle *Rectangle) SetHeight(height float64) {
	rectangle.height = height
}
