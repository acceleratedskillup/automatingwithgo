package main

import "fmt"

func main() {
	fmt.Println("helloworld", 10, 16.2)

	/*
		printf does not print carriage return
		need to put in '\n' at the end of the line
	*/
	fmt.Printf("student %s has scored %d marks\n", "mark", 80)
	var message = fmt.Sprintf("movie %s has scored %.2f million at the box office\n", "Transformers", 709.7)
	fmt.Println(message)
}
