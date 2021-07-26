package main

import "fmt"

func main() {
	example1()
	//example2()
	//example3()
	//example4()
	//example5()
	//example6and7()
}

func example1() {
	intArray := [6]int16{0, 1, 2, 3, 4, 5}
	fmt.Printf("intArray) a-addr:%p\n", &intArray)
	intSlice1 := intArray[0:6]
	printSlice("intSlice1", intSlice1)
	intSlice1 = append(intSlice1, 6)
	printSlice("intSlice1", intSlice1)

	name := "john smith"
	firstName := name[:4]
	fmt.Println("firstName:", firstName)
}

func example2() {
	intArray := [6]int16{0, 1, 2, 3, 4, 5}
	fmt.Printf("intArray) a-addr:%p\n", &intArray)
	intSlice1 := intArray[0:6]
	intSlice1 = append(intSlice1, 4)
	printSlice("intSlice1", intSlice1)

	intSlice2 := intSlice1[2:]
	printSlice("[2:]", intSlice2)

	intSlice2 = intSlice1[:2]
	printSlice("[:2]", intSlice2)

	//increase again length of slice
	intSlice2 = intSlice1[:4]
	printSlice("[:4]", intSlice2)

	//increase again length of slice
	intSlice2 = intSlice1[:100]

}

func example3() {
	intArray := [6]int16{0, 1, 2, 3, 4, 5}
	fmt.Printf("intArray) a-addr:%p\n", &intArray)
	intSlice1 := intArray[0:6]
	printSlice("intSlice1", intSlice1)

	intSlice2 := intSlice1[:2]
	printSlice("intSlice2", intSlice2)

	intSlice2 = append(intSlice2, 100)
	printSlice("intSlice1", intSlice1)
	printSlice("intSlice2", intSlice2)
}

func example4() {
	intArray := [6]int16{0, 1, 2, 3, 4, 5}
	fmt.Printf("intArray) a-addr:%p\n", &intArray)
	intSlice1 := intArray[0:6]
	printSlice("intSlice1", intSlice1)

	intSlice2 := intSlice1[3:]
	printSlice("intSlice2", intSlice2)
	intSlice2 = append(intSlice2, 100)
	fmt.Println("after appending")
	printSlice("intSlice1", intSlice1)
	printSlice("intSlice2", intSlice2)
}

func example5and6() {
	intArray := [6]int16{0, 1, 2, 3, 4, 5}
	fmt.Printf("intArray) a-addr:%p\n", &intArray)
	intSlice1 := intArray[0:6]
	printSlice("intSlice1", intSlice1)

	intSlice2 := intSlice1[3:]
	fmt.Println("intSlice1[3]:", intSlice1[3])
	fmt.Println("intSlice2[0]:", intSlice2[0])

}

func printSlice(prefix string, input []int16) {
	fmt.Printf("%s) len():%d cap():%d s-addr:%p a-addr:%p\n", prefix, len(input), cap(input), &input, input)
	fmt.Printf("%*v %v \n", len(prefix)+1, " ", input)
}
