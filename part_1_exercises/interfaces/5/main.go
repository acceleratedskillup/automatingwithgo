package main

import "fmt"

type Box struct {
	Item interface{}
}

func PlaceInBox(b *Box, item interface{}) {
	b.Item = item
}

func main() {
	b := &Box{}
	PlaceInBox(b, "A shiny gem")
	fmt.Println(b.Item)

	PlaceInBox(b, 42)
	fmt.Println(b.Item)
}
