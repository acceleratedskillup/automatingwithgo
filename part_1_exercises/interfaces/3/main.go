package main

import "fmt"

type Color struct {
	R, G, B int
}

func (c Color) String() string {
	return fmt.Sprintf("R: %d, G: %d, B: %d", c.R, c.G, c.B)
}

func main() {
	red := Color{R: 255, G: 0, B: 0}
	fmt.Println(red)
}
