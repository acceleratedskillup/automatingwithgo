package main

import "fmt"

func orderDish(dishOrders map[string]int, dish string) {
	dishOrders[dish]++
}

func main() {
	dishes := map[string]int{
		"Pasta":  0,
		"Burger": 0,
		"Salad":  0,
	}

	orderDish(dishes, "Pasta")
	orderDish(dishes, "Burger")
	orderDish(dishes, "Pasta")

	fmt.Println("Dish Popularity:", dishes)
}
