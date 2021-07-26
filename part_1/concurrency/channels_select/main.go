package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type StockPrice struct {
	symbol string
	price  float64
}

const BUFFER_SIZE = 10

func main() {
	stockPricesChannel := make(chan StockPrice, BUFFER_SIZE)
	newsChannel := make(chan string, BUFFER_SIZE)
	rand.Seed(time.Now().UnixNano())
	symbols := []string{"msft", "tsla", "aapl", "goog"}
	news := readNewsFile("./news.txt")
	go GenerateStockPrices(symbols, stockPricesChannel)
	go GenerateNews(news, newsChannel)
	for {
		select {
		case stockPrice, ok := <-stockPricesChannel:
			if !ok {
				stockPricesChannel = nil
			} else {
				fmt.Printf("symbol:%s price:%.2f\n", stockPrice.symbol, stockPrice.price)
			}

		case news, ok := <-newsChannel:
			if !ok {
				newsChannel = nil
			} else {
				fmt.Printf("breaking news:%s\n", news)
			}
		}
		if stockPricesChannel == nil && newsChannel == nil {
			fmt.Println("all channels are closed")
			break
		}
	}
}

func GenerateStockPrices(symbols []string, stockPricesChannel chan<- StockPrice) {
	defer close(stockPricesChannel)
	for i := 0; i < 10; i++ {
		randomStockIndex := getRandomInt(0, len(symbols)-1)
		stockPrice := StockPrice{
			symbol: symbols[randomStockIndex],
			price:  getRandomPrice(100, 300),
		}
		stockPricesChannel <- stockPrice
		time.Sleep(time.Second)
	}
}

func GenerateNews(news []string, newsChannel chan<- string) {
	defer close(newsChannel)
	for _, newsLine := range news {
		newsChannel <- newsLine
		time.Sleep(time.Second)
	}
}

func readNewsFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return []string{}
	}
	defer file.Close()
	var news []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		news = append(news, scanner.Text())
	}

	return news
}

func getRandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

func getRandomPrice(min float64, max float64) float64 {
	return (min + rand.Float64()*(max-min))
}
