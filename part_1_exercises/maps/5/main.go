package main

import "fmt"

func updateProfile(profile map[string]string, key, value string) {
	profile[key] = value
}

func main() {
	userProfile := map[string]string{
		"Name":    "John",
		"Email":   "john@example.com",
		"Country": "USA",
	}

	updateProfile(userProfile, "Country", "Canada")
	fmt.Println(userProfile) // map[Name:John Email:john@example.com Country:Canada]
}
