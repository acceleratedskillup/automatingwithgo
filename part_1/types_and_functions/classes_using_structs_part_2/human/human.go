package human

import "fmt"

type Human struct {
	name   string //take note of small letters
	height float64
}

//note that the names are capitalized
func (human *Human) Walk() {
	fmt.Printf("%s is walking\n", human.name)
}

func (human *Human) Run() {
	fmt.Printf("%s is running\n", human.name)
}

func (human *Human) GetName() string {
	return human.name
}

func (human *Human) SetName(newName string) {
	human.name = newName
}

func (human *Human) GetHeight() float64 {
	return human.height
}

func (human *Human) SetHeight(newHeight float64) {
	human.height = newHeight
}

func (human *Human) GetString() string {
	return fmt.Sprintf("human name:%s height:%v", human.name, human.height)
}
