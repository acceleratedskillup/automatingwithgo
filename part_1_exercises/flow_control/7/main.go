package main

import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("radar")) // Expected output: true
	fmt.Println(isPalindrome("hello")) // Expected output: false
}
