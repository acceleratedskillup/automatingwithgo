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

func main() {
	start := time.Now()
	files := getFilesInDir(FILES_DIR)
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(files))
	for _, file := range files {
		go printCountOfWords(file, &waitGroup)
	}

	waitGroup.Wait()
	elapsed := time.Since(start)
	fmt.Println("time taken:", elapsed)
}

func printCountOfWords(file fs.FileInfo, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	content, err := os.ReadFile(FILES_DIR + "/" + file.Name())
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s has %d words\n", file.Name(), len(strings.Fields(string(content))))
}

func getFilesInDir(dir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []fs.FileInfo{}
	}
	return files
}
