package main

import (
	"fmt"
	"sync"
)

type Bank struct {
	balance int
	mu      sync.Mutex
}

func (b *Bank) Deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}

func (b *Bank) Withdraw(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance -= amount
}

func main() {
	b := &Bank{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b.Deposit(10)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			b.Withdraw(5)
		}()
	}

	wg.Wait()
	fmt.Println("Final balance:", b.balance)
}
