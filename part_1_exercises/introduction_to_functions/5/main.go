package main

import "fmt"

func sumFirstThree(nums ...int) int {
	return nums[0] + nums[1] + nums[2]
}

func main() {
	result := sumFirstThree(1, 2, 3, 4, 5)
	fmt.Println(result) // Expected output: 6
}
