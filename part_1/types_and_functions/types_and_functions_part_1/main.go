package main

import "fmt"

type bigint int64
type smallint int32

func main() {
	example2()
	//example3()

}

func example1() {
	myBI := bigint(5)
	mySI := smallint(5)

	fmt.Println("myBI:", myBI)
	fmt.Println("mySI:", mySI)
}

//example1
/*
func Print(bi bigint) {

}

func Print(si smallint) {

}
*/
func (si smallint) PrintByCopy() {
	si++
	fmt.Printf("%-16s %10s %v %10s %p\n", "PrintByCopy()", "bi:", si, "&bi:", &si)
}

func (bi bigint) PrintByCopy() {
	bi++
	fmt.Printf("%-16s %10s %v %10s %p\n", "PrintByCopy()", "bi:", bi, "&bi:", &bi)
}

func example2() {
	myBI := bigint(5)
	mySI := smallint(1)
	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "myBI:", myBI, "&myBI:", &myBI)
	myBI.PrintByCopy()
	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "myBI:", myBI, "&myBI:", &myBI)

	fmt.Println()

	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "mySI:", mySI, "&mySI:", &mySI)
	mySI.PrintByCopy()
	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "mySI:", mySI, "&mySI:", &mySI)

}

type Building struct {
	Width  int
	Height int
	Depth  int
}

func (b Building) PrintByCopy() {
	b.Width = 999
	fmt.Printf("%-16s %10s %v %10s %p\n", "PrintByCopy()", "b:", b, "&b:", &b)
}

func example3() {
	myB := Building{
		100, 200, 300,
	}
	//fmt.Printf("buildingA addr:%p\n", &buildingA)
	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "myB:", myB, "&myB:", &myB)
	myB.PrintByCopy()
	fmt.Printf("%-16s %10s %v %10s %p\n", "main()", "myB:", myB, "&myB:", &myB)
}
