package alien

import "fmt"

type alien struct {
	name   string //take note of small letters
	height float64
}

func New(name string, height float64) *alien {
	initialize()
	return &alien{name, height}
}

//constructor "overloading"
func NewWithName(name string) *alien {
	return New(name, 0)
}

//constructor "overloading"
func NewWithHeight(height float64) *alien {
	return New("", height)
}

func initialize() {
	fmt.Println("doing some other initialization")
}

//note that the names are capitalized
func (alien *alien) BlastLasers() {
	fmt.Printf("%s is blasting lasers\n", alien.name)
}

func (alien *alien) GetName() string {
	return alien.name
}

func (alien *alien) SetName(newName string) {
	alien.name = newName
}

func (alien *alien) GetHeight() float64 {
	return alien.height
}

func (alien *alien) SetHeight(newHeight float64) {
	alien.height = newHeight
}

func (alien *alien) GetString() string {
	return fmt.Sprintf("alien name:%s height:%v", alien.name, alien.height)
}
