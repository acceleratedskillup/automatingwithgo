package main

import (
	"fmt"
	"math/rand"
	s "packages/stack/utils"
	u "packages/utils"

	// always specify the current module's name first before the internal package name

	//lint:ignore ST1001 just ignore for time being
	. "packages/views"

	// an example of using an alias if it helps with long package names

	_ "packages/initialize"
	//"github.com/atotto/clipboard" unused so commented out first
)

func main() {
	fmt.Println(rand.Int31n(10))

	u.PrintDummyString()
	u.ToPrint = true

	s.PrintDummyString()
	GenerateTextOutput()
}
