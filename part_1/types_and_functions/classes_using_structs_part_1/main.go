package main

import (
	"classespart1/human1"
	"classespart1/human2"
	"fmt"
)

func main() {
	//example1()
	//example2()
	//example3()
	//example4()
	example5()
}

/*
 To show how to create a class
*/
func example1() {
	john := new(human1.Human1)
	fmt.Println("before:", john.GetString())
	john.SetName("john")
	john.SetHeightInMeters(1.88)
	fmt.Println("after:", john.GetString())

	//fmt.Println("Name:", john.GetName()) //note that you cannot access the struct variables directly

	john.Walk()
	john.Run()

}

func example2() {
	/*
		can use struct literals but cannot init fields as they are exported
		might want to do an example of changing the capitalization of the Human struct
		and initialize from there to show that we can init it from there but we lose the point
		of having getters and setters
	*/
	john := human1.Human1{}
	fmt.Println("before:", john.GetString())
	john.SetName("john")
	john.SetHeightInMeters(1.88)
	fmt.Println("after:", john.GetString())

	//fmt.Println("Name:", john.GetName()) //note that you cannot access the struct variables directly

	john.Walk()
	john.Run()
}

/*
	to show how to force users to use our constructors
	+ "overloaded" constructors
	instead of using new built-in function or
	using struct literal
*/
func example3() {

	/*
		cannot use new keyword to create struct
		as the struct is not exported from package

		borg := new(alien.alien)
	*/
	john := human2.New("John", 5.00)
	fmt.Println(john.GetString())

}

func example4() {
	john := human2.NewWithName("John")
	fmt.Println(john.GetString())
}

func example5() {
	john := human2.NewWithHeight(1.60)
	fmt.Println(john.GetString())
}
