package main

import "fmt"

func main() {
	words := []string{"apple", "banana", "secret1", "grape", "orange", "secret2", "cherry", "blueberry", "secret3"}
	var secretMessage []string

	for i := 2; i < len(words); i += 3 {
		secretMessage = append(secretMessage, words[i])
	}

	fmt.Println("The secret message is:", secretMessage)
}
