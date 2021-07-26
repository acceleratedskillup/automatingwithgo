package main

import (
	"fmt"
	"math/rand"
	"time"
)

type StockPrice struct {
	symbol string
	price  float64
}

func main() {
	stockPricesChannel := make(chan StockPrice, 10)
	rand.Seed(time.Now().UnixNano())

	go func(symbol string, stockPricesChannel chan StockPrice) {
		for i := 0; i < 100; i++ {
			stockPrice := StockPrice{
				symbol: symbol,
				price:  GetRandomPrice(100, 300),
			}
			fmt.Printf("sent: symbol:%s price:%.2f addr:%p\n", stockPrice.symbol, stockPrice.price, &stockPrice)
			stockPricesChannel <- stockPrice
		}
		close(stockPricesChannel)
	}("msft", stockPricesChannel)

	example1(stockPricesChannel)
	//example2(stockPricesChannel)
}

func example1(stockPricesChannel chan StockPrice) {
	for {
		stockPrice, isThereMore := <-stockPricesChannel
		if isThereMore {
			fmt.Printf("received: symbol:%s price:%.2f addr:%p\n", stockPrice.symbol, stockPrice.price, &stockPrice)
		} else {
			fmt.Println("received all stock prices ")
			break
		}
	}
}

func example2(stockPricesChannel chan StockPrice) {
	for stockPrice := range stockPricesChannel {
		fmt.Printf("stock:%s price:%.2f\n", stockPrice.symbol, stockPrice.price)
	}
	fmt.Println("received all stock prices ")
}

func GetRandomPrice(min float64, max float64) float64 {
	return (min + rand.Float64()*(max-min))
}
