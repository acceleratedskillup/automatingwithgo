package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting download for:", filename)
	time.Sleep(2 * time.Second)
	fmt.Println("Completed download for:", filename)
}

func main() {
	var wg sync.WaitGroup
	files := []string{"file1.txt", "file2.txt", "file3.txt"}

	for _, file := range files {
		wg.Add(1)
		go downloadFile(file, &wg)
	}

	wg.Wait()
	fmt.Println("All downloads completed!")
}
