package human2

import (
	"fmt"
	"strings"
)

type human2 struct {
	name       string //take note of small letters
	heightInCM float64
}

func New(name string, heightInMeters float64) *human2 {
	initialize(&name)
	return &human2{name, heightInMeters * 100}
}

//constructor "overloading"
func NewWithName(name string) *human2 {
	return New(name, 1.50)
}

//constructor "overloading"
func NewWithHeight(heightInMeters float64) *human2 {
	return New("Human", heightInMeters)
}

func initialize(name *string) {
	*name = strings.ToUpper(*name)
	fmt.Println("doing some other initialization")
}

func (h2 *human2) Run() {
	fmt.Printf("%s is running\n", h2.name)
}

func (h2 *human2) Walk() {
	fmt.Printf("%s is walking\n", h2.name)
}

func (h2 *human2) GetName() string {
	return h2.name
}

func (h2 *human2) SetName(newName string) {
	h2.name = newName
}

func (h2 *human2) GetHeightInMeters() float64 {
	return h2.heightInCM / 100
}

func (h2 *human2) SetHeightInMeters(newHeightInMeters float64) {
	h2.heightInCM = newHeightInMeters * 100
}

func (h2 *human2) GetString() string {
	return fmt.Sprintf("name:%s heightInCM:%v", h2.name, h2.heightInCM)
}
