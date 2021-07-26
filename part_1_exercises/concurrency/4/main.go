package main

import (
	"fmt"
	"math/rand"
	"time"
)

type StockPrice struct {
	Company string
	Price   float64
}

func sendStockPrices(c chan StockPrice) {
	companies := []string{"Apple", "Google", "Amazon"}
	for i := 0; i < 10; i++ {
		company := companies[rand.Intn(len(companies))]
		price := rand.Float64() * 1000
		c <- StockPrice{Company: company, Price: price}
		time.Sleep(time.Second)
	}
	close(c)
}

func main() {
	c := make(chan StockPrice)
	go sendStockPrices(c)

	var maxPrice, minPrice float64 = -1, 1e9
	for stock := range c {
		if stock.Price > maxPrice {
			maxPrice = stock.Price
		}
		if stock.Price < minPrice {
			minPrice = stock.Price
		}
	}

	fmt.Println("Highest Price:", maxPrice)
	fmt.Println("Lowest Price:", minPrice)
}
