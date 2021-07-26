package main

import (
	"fmt"
	"sync"
	"time"
)

func joinClass(studentName string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(studentName, "has joined the class.")
	time.Sleep(1 * time.Second)
}

func main() {
	var wg sync.WaitGroup
	students := []string{"Alice", "Bob", "Charlie", "David"}

	for _, student := range students {
		wg.Add(1)
		go joinClass(student, &wg)
	}

	wg.Wait()
	fmt.Println("All students have joined. Class is starting!")
}
