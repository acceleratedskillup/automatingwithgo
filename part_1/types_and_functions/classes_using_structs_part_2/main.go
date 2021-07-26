package main

import (
	"classesusingstructs/alien"
	"classesusingstructs/human"
	"classesusingstructs/omnientity"
	"fmt"
)

func main() {

	//example 1
	humans()
	//example 2
	aliens()
	//example 3
	omnientities()

}

func humans() {
	john := new(human.Human)
	john.SetName("john")
	john.SetHeight(1.88)

	john.Walk()
	john.Run()

	fmt.Println(john.GetName()) //note that you cannot access the struct variables directly
	fmt.Println("height of human:", john.GetHeight())
	john.SetHeight(2.00)
	fmt.Println("height of human:", john.GetHeight())

	/*
		can use struct literals but cannot init fields as they are exported
		might want to do an example of changing the capitalization of the Human struct
		and initialize from there to show that we can init it from there but we lose the point
		of having getters and setters
	*/
	mary := human.Human{}
	fmt.Println("mary:", mary.GetString())
}

func aliens() {

	//cannot use new keyword to create struct as the struct is not exported from package
	borg := alien.New("borg", 5.00)
	fmt.Println(borg.GetString())
	borg.SetHeight(2.00)
	fmt.Println(borg.GetString())
	borg.BlastLasers()

	predator := alien.NewWithName("predator")
	fmt.Println(predator.GetString())
}

func omnientities() {
	omnientity := omnientity.GetInstance()
	fmt.Println(omnientity.GetString())
}
