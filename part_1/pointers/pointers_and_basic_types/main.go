package main

import "fmt"

func main() {

	//example1()
	example1_1()
	//example1_2()
	//example2()
	//example3()
	//example4()

	//the above examples also apply to other basic types

}

func example1() {
	//declaring a int pointer which after declaration will be nil.
	var i *int
	//take note that we are using i in the 2nd parameter, not *i
	fmt.Printf("type:%T, &i:%p, i:%p, i:%v\n", i, &i, i, i)

	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is NOT nil")
	}
	//error statement
	//fmt.Printf("type:%T, &i:%p, i:%p, *i:%v\n", i, &i, i, *i)
}

func example1_1() {
	//use the builtin new() function to obtain a pointer to a zero-valued int
	var i *int = new(int)
	fmt.Printf("type:%T, &i:%p, i:%p, *i:%v\n", i, &i, i, *i)
	//then you can assign a value
	*i = 123
	fmt.Println("after *i = 123")
	fmt.Printf("type:%T, &i:%p, i:%p, *i:%v\n", i, &i, i, *i)
}

//more concise version
func example1_2() {
	i := new(int)
	fmt.Printf("type:%T, &i:%p, i:%p, *i:%v\n", i, &i, i, *i)
	//then you can assign a value
	*i = 123
	fmt.Println("after *i = 123")
	fmt.Printf("type:%T, &i:%p, i:%p, *i:%v\n", i, &i, i, *i)
}

/*
assignment to pointer variable from
another pointer variable
*/
func example2() {

	i := new(int)
	*i = 123
	fmt.Printf("type:%T, &i:%p, i:%p, *i:%v\n", i, &i, i, *i)

	x := new(int)
	*x = 456
	fmt.Printf("type:%T, &x:%p, x:%p, *x:%v\n", x, &x, x, *x)
	x = i
	fmt.Println("after x = i")
	fmt.Printf("type:%T, &x:%p, x:%p, *x:%v\n", x, &x, x, *x)
}

/*
assignment of normal variable to pointer variable
*/
func example3() {

	i := new(int)
	*i = 123
	fmt.Printf("type:%T, &i:%p, i:%p, *i:%v\n", i, &i, i, *i)

	x := 456
	fmt.Printf("type:%T, &x:%p, x:%v\n", x, &x, x)
	i = &x //assigning the address of x to i
	fmt.Println("after i = &x")
	fmt.Printf("type:%T, &i:%p, i:%p, *i:%v\n", i, &i, i, *i)

}

/*
what happens when we create and
initialize a new variable using &
*/
func example4() {
	x := 456
	fmt.Printf("type:%T, &x:%p, x:%v\n", x, &x, x)

	//special prefix & returns a pointer to the int value.
	i := &x
	fmt.Println("after i := &x")
	fmt.Printf("type:%T, &i:%p, i:%p, *i:%v\n", i, &i, i, *i)

}
