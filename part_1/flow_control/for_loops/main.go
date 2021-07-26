package main

import "fmt"

func main() {

	/*
		3 components to a for loop
		the init statement: executed before the first iteration
		the condition expression: evaluated before every iteration
		the post statement: executed at the end of every iteration
		i is only available inside the loop
	*/
	sum := 0
	fmt.Println("using all 3 components of for loop")
	for i := 0; i < 10; i++ {
		sum = sum + 2
	}
	fmt.Println("sum:", sum)

	fmt.Println("another way to write a for loop")
	sum = 0
	i := 0 // new variable, different from the loop above
	//for ;i < 10;i++ {
	for i < 10 {
		sum = sum + 2
		i++
	}
	fmt.Println("sum:", sum)

	fmt.Println("another way of writing for loops using break")
	sum = 0
	i = 0
	for {
		//need to have a condition to break out or the loop will run forever
		if i >= 10 {
			fmt.Println("breaking out of loop")
			break
		}
		sum = sum + 2
		i++
	}
	fmt.Println("sum:", sum)

	fmt.Println("using multiple variables in for loop")
	for i, c := 0, 'a'; i < 5; i, c = i+1, c+1 {
		fmt.Printf("i:%d c:%c\n", i, c)
	}
	fmt.Println()
	/*
		you can use continue keyword to jump to next iteration of the loop
		bypassing all code below that point
		below we are printing out odd numbers only
	*/

	//printing only odd numbers
	fmt.Println("another way of writing for loops using continue")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	//nested for loop
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("i:%d j:%d\n", i, j)
		}
	}

	fmt.Println("classic way to break out of nested loops")
	breakLoops := false
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("i:%d j:%d\n", i, j)
			if j == 1 {
				breakLoops = true
				break
			}
		}
		if breakLoops {
			break
		}
	}

	fmt.Println("breaking out of nested loops using labels")
Start:
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("i:%d j:%d\n", i, j)
			if j == 1 {
				break Start //continue, break
			}
		}
		fmt.Printf("is this ever executed?")
	}

	fmt.Printf("nested for loop has finished")
}
