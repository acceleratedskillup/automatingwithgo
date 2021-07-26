package main

import "fmt"

func main() {
	//example 1
	fmt.Println("simple example of usage of a label")
	goto End
	fmt.Println("will this ever get printed")
End:
	fmt.Println("this is the end")

	//example 2
	fmt.Println("creating a loop")
	x := 0
Loop:
	if x < 5 {
		fmt.Println("x:", x)
		x++
		goto Loop
	} else {
		goto OutOfLoop
	}
OutOfLoop:
	fmt.Println("End")

	//example 3
	//you can use labels and variables of the same name
	fmt.Println("using same names for label and variable")
	Y := 1
Y:
	if Y > 3 {
		fmt.Println("Y is now > 3:", Y)
	} else {
		fmt.Println("Y used to be:", Y)
		Y = 5
		fmt.Println("now it is:", Y)
		fmt.Println("goto Y now")
		goto Y
	}
	/*
		example 4
		this works if your label is declared outside
		a block
	*/
	fmt.Println("\ndeclaring label outside a block")
	x = 0
Start2:
	{
		if x <= 2 {
			x++
			fmt.Println("x:", x)
			goto Start2
		}
		fmt.Println("x is > 2")
	}
	/*
		if your label is declared inside a block
		and goto is outside,it doesnt work
		**code cannot be compiled**
	*/
	/*
		goto Block
		{
		Block:
			fmt.Println("inside block")
		}
	*/

	/*
		example 5
		labels are not block scoped
	*/
	/*
		J:
			{
			J:
				fmt.Println("labels are not block scoped")
			}
	*/
}

/*example 6
labels are valid in the scope of the whole function
where it is declared
invalid code below
*/
/*
func someFunc() {
	goto Start2
}
*/
/*
func someFunc() {
SomeFuncLabel:
	fmt.Println("this is calling from label Start2")
}
*/
