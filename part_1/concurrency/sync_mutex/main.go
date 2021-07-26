package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const FILES_DIR = "./files"

type AtomicTotalWordCount struct {
	lock  sync.Mutex // A lock than can be held by one goroutine at a time.
	count int
}

func (a *AtomicTotalWordCount) Add(filename string, fileWordCount int) {
	a.lock.Lock()         // Wait for the lock to be free and then take it.
	defer a.lock.Unlock() // Release the lock.

	tempTotalWordCount := a.count + fileWordCount
	fmt.Printf("%s) tempTotalWordCount:%d\n", filename, tempTotalWordCount)
	time.Sleep(500 * time.Millisecond)
	a.count = tempTotalWordCount
	fmt.Printf("%s) totalWordCount:%d\n", filename, a.count)
}

var totalWordCount int

func main() {
	example1()
	//example2()
	//example3()
}

/*
without using lock
*/
func example1() {
	start := time.Now()
	files := getFilesInDir(FILES_DIR)
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(files))
	for _, file := range files {
		go printCountOfWordsWithWaitGroup(file, &waitGroup)
	}

	waitGroup.Wait()
	fmt.Println("totalWordCount:", totalWordCount)
	elapsed := time.Since(start)
	fmt.Println("time taken:", elapsed)
}

func printCountOfWordsWithWaitGroup(file fs.FileInfo, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	content, err := os.ReadFile(FILES_DIR + "/" + file.Name())
	if err != nil {
		log.Fatal(err)
		return
	}

	words := len(strings.Fields(string(content)))
	fmt.Printf("%s has %d words\n", file.Name(), words)
	tempTotalWordCount := totalWordCount + words
	fmt.Printf("%s) tempTotalWordCount:%d\n", file.Name(), tempTotalWordCount)
	time.Sleep(500 * time.Millisecond)
	totalWordCount = tempTotalWordCount
	fmt.Printf("%s) totalWordCount:%d\n", file.Name(), totalWordCount)
}

/*
using lock
*/
func example2() {
	start := time.Now()
	files := getFilesInDir(FILES_DIR)
	var waitGroup sync.WaitGroup
	var lock sync.Mutex
	waitGroup.Add(len(files))
	for _, file := range files {
		go printCountOfWordsWithWaitGroupAndLock(file, &waitGroup, &lock)
	}

	waitGroup.Wait()
	fmt.Println("totalWordCount:", totalWordCount)
	elapsed := time.Since(start)
	fmt.Println("time taken:", elapsed)
}

func printCountOfWordsWithWaitGroupAndLock(file fs.FileInfo, waitGroup *sync.WaitGroup, lock *sync.Mutex) {
	defer waitGroup.Done()
	content, err := os.ReadFile(FILES_DIR + "/" + file.Name())
	if err != nil {
		log.Fatal(err)
		return
	}

	words := len(strings.Fields(string(content)))
	fmt.Printf("%s has %d words\n", file.Name(), words)
	lock.Lock()
	defer lock.Unlock()
	tempTotalWordCount := totalWordCount + words
	fmt.Printf("%s) tempTotalWordCount:%d\n", file.Name(), tempTotalWordCount)
	time.Sleep(500 * time.Millisecond)
	totalWordCount = tempTotalWordCount
	fmt.Printf("%s) totalWordCount:%d\n", file.Name(), totalWordCount)
}

/*
using lock struct
*/
func example3() {
	start := time.Now()
	files := getFilesInDir(FILES_DIR)
	var waitGroup sync.WaitGroup
	var totalWordCount AtomicTotalWordCount
	waitGroup.Add(len(files))
	for _, file := range files {
		go printCountOfWordsWithWaitGroupAndAtomicAdd(file, &waitGroup, &totalWordCount)
	}

	waitGroup.Wait()
	fmt.Println("totalWordCount:", totalWordCount.count)
	elapsed := time.Since(start)
	fmt.Println("time taken:", elapsed)
}
func printCountOfWordsWithWaitGroupAndAtomicAdd(file fs.FileInfo, waitGroup *sync.WaitGroup, totalWordCount *AtomicTotalWordCount) {
	defer waitGroup.Done()
	content, err := os.ReadFile(FILES_DIR + "/" + file.Name())
	if err != nil {
		log.Fatal(err)
		return
	}

	words := len(strings.Fields(string(content)))
	fmt.Printf("%s has %d words\n", file.Name(), words)
	totalWordCount.Add(file.Name(), words)
}
func getFilesInDir(dir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []fs.FileInfo{}
	}
	return files
}
