package main

import "fmt"

func doubleValues(arr [5]int) [5]int {
	for i, val := range arr {
		arr[i] = val * 2
	}
	return arr
}
func main() {

	var original = [5]int{1, 2, 3, 4, 5}
	doubled := doubleValues(original)
	fmt.Println("Original array:", original)
	fmt.Println("Doubled array:", doubled)

}
