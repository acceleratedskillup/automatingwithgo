package main

import "fmt"

type Tree struct {
	age     int
	season  string
	secrets []string
}

func (t *Tree) whisperSecret() string {
	return t.secrets[0] // For simplicity, returning the first secret
}

func (t *Tree) changeAppearance() string {
	switch t.season {
	case "Spring":
		return "The tree is full of blossoms."
	case "Summer":
		return "The tree provides shade with its dense leaves."
	case "Autumn":
		return "The tree's leaves turn golden and fall."
	case "Winter":
		return "The tree is bare and covered in snow."
	default:
		return "The tree stands tall, unaffected by time."
	}
}

func main() {
	willow := &Tree{
		age:     1000,
		season:  "Autumn",
		secrets: []string{"I've seen the rise and fall of empires.", "Magic is real.", "The forest protects its own."},
	}
	fmt.Println(willow.whisperSecret())
	fmt.Println(willow.changeAppearance())
}
