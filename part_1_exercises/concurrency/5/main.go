package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func sendHeadlines(c chan string) {
	headlines := []string{
		"Market hits all-time high!",
		"New tech breakthrough promises faster computers.",
		"Sports team wins championship after dramatic game.",
		"Tech stocks soar as market rallies.",
	}
	for i := 0; i < 10; i++ {
		headline := headlines[rand.Intn(len(headlines))]
		c <- headline
		time.Sleep(2 * time.Second)
	}
	close(c)
}

func main() {
	c := make(chan string)
	go sendHeadlines(c)

	keyword := "tech"
	for headline := range c {
		if strings.Contains(strings.ToLower(headline), keyword) {
			fmt.Println("Filtered News:", headline)
		}
	}
}
