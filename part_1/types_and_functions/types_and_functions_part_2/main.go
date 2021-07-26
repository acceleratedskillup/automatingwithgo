package main

import "fmt"

type bigint int64

func main() {

	//example1()
	example2()

}

/*
	using a POINTER variable
	show how we call a pointer receiver function
	show how we can also call a value receiver function
*/

func (biPtr *bigint) PrintByPointer() {
	*biPtr++
	//(*biPtr)++
	fmt.Printf("%-16s %10s %v %10s %p\n", "PrintByPointer()", "*biPtr:", *biPtr, "biPtr:", biPtr)
}

func example1() {
	myBIPtr := new(bigint)
	*myBIPtr = 10
	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "*myBIPtr:", *myBIPtr, "myBIPtr:", myBIPtr)
	myBIPtr.PrintByPointer()
	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "*myBIPtr:", *myBIPtr, "myBIPtr:", myBIPtr)
	myBIPtr.PrintByCopy()
	/*
		alternatively:
		(*myBIPtr).PrintByCopy()

		call by type
		(*bigint).PrintByPointer(myBIPtr)
		(*bigint).PrintByCopy(myBIPtr)
		bigint.PrintByCopy(*myBIPtr)
	*/
	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "*myBIPtr:", *myBIPtr, "myBIPtr:", myBIPtr)

}

/*
	using a NORMAL variable
	show how we call a pointer receiver function
	show how we can also call a value receiver function
*/
func example2() {
	myBI := bigint(10)
	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "myBI:", myBI, "&myBI:", &myBI)
	myBI.PrintByPointer()
	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "myBI:", myBI, "&myBI:", &myBI)
	myBI.PrintByCopy()

	/*
		alternatively
		(&myBI).PrintByPointer()

		call by type
		(*bigint).PrintByPointer(&myBI)
		(*bigint).PrintByCopy(&myBI)
		bigint.PrintByCopy(myBI)
	*/
	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "myBI:", myBI, "&myBI:", &myBI)
}

//the value of i is a copied value
func (bi bigint) PrintByCopy() {
	bi++
	fmt.Printf("%-16s %10s %v %10s %p\n", "PrintByCopy()", "bi:", bi, "&bi:", &bi)
}

/*
	example 1
   	negative example
	func (i int) print() {

	}
*/
