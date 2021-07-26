package main

import "fmt"

func findLongestWord(sentence string) string {
	longestWord := ""
	word := ""
	for _, char := range sentence {
		if char == ' ' {
			if len(word) > len(longestWord) {
				longestWord = word
			}
			word = ""
		} else {
			word += string(char)
		}
	}
	// Check for the last word in the sentence
	if len(word) > len(longestWord) {
		longestWord = word
	}
	return longestWord
}

func main() {
	fmt.Println(findLongestWord("The quick brown fox"))      // Expected output: "quick"
	fmt.Println(findLongestWord("Jumped over the lazy dog")) // Expected output: "Jumped"
}
