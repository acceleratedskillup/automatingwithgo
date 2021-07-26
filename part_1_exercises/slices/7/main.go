package main

import "fmt"

func main() {
	easternSpices := []string{"Saffron", "Turmeric", "Cardamom"}
	westernHerbs := []string{"Rosemary", "Thyme", "Oregano"}
	allIngredients := append(easternSpices, westernHerbs...)
	fmt.Println(allIngredients)
}
