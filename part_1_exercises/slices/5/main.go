package main

import "fmt"

func main() {
	playlist := []string{"Moon's Serenade", "Star's Dance", "Sun's Ballad", "Sky's Symphony"}
	indexToRemove := 2
	playlist = append(playlist[:indexToRemove], playlist[indexToRemove+1:]...)
	fmt.Println(playlist)
}
