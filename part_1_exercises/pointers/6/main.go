package main

import "fmt"

func main() {
	var setting *string
	if setting == nil {
		fmt.Println("Setting is not initialized!")
	} else {
		fmt.Println("Setting:", *setting)
	}
}
