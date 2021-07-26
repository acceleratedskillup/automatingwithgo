package main

import (
	"fmt"
)

type Operation struct {
	A, B     int
	Operator string
}

func calculator(op Operation, c chan int) {
	switch op.Operator {
	case "+":
		c <- op.A + op.B
	case "-":
		c <- op.A - op.B
	case "*":
		c <- op.A * op.B
	case "/":
		c <- op.A / op.B
	default:
		c <- 0
	}
}

func main() {
	c := make(chan int)
	operations := []Operation{
		{A: 5, B: 3, Operator: "+"},
		{A: 8, B: 2, Operator: "-"},
		{A: 6, B: 3, Operator: "*"},
		{A: 8, B: 4, Operator: "/"},
	}

	for _, op := range operations {
		go calculator(op, c)
	}

	for i := 0; i < len(operations); i++ {
		result := <-c
		fmt.Println("Result:", result)
	}
}
