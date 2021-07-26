package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"

	"regexp"
)

func main() {

	// Read text from clipboard
	clipText, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Failed to read from clipboard:", err)
		return
	}

	// Regular expression for phone numbers (US format)
	phoneRegex := regexp.MustCompile(`\(?[0-9]{3}\)?[-. ]?[0-9]{3}[-. ]?[0-9]{4}`)

	// Regular expression for email addresses
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Find phone numbers and email addresses
	phoneNumbers := phoneRegex.FindAllString(clipText, -1)
	emailAddresses := emailRegex.FindAllString(clipText, -1)

	// Combine results
	allResults := append(phoneNumbers, emailAddresses...)

	// Replace clipboard content with found phone numbers and email addresses
	if len(allResults) > 0 {
		newClipText := strings.Join(allResults, "\n")
		err := clipboard.WriteAll(newClipText)
		if err != nil {
			fmt.Println("Failed to write to clipboard:", err)
			return
		}
		fmt.Println("Clipboard updated with phone numbers and email addresses.")
	} else {
		fmt.Println("No phone numbers or email addresses found.")
	}
}
