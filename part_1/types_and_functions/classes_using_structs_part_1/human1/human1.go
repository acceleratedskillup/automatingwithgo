package human1

import "fmt"

type Human1 struct {
	name       string //take note of small letters
	heightInCM float64
}

//note that the names are capitalized
func (h1 *Human1) Walk() {
	fmt.Printf("%s is walking\n", h1.name)
}

func (h1 *Human1) Run() {
	fmt.Printf("%s is running\n", h1.name)
}

func (h1 *Human1) GetName() string {
	return h1.name
}

func (h1 *Human1) SetName(newName string) {
	h1.name = newName
}

func (human *Human1) GetHeightInMeters() float64 {
	return human.heightInCM / 100
}

func (h1 *Human1) SetHeightInMeters(newHeightInMeters float64) {
	h1.heightInCM = newHeightInMeters * 100
}

func (h1 *Human1) GetString() string {
	return fmt.Sprintf("name:%s heightInCM:%v", h1.name, h1.heightInCM)
}
