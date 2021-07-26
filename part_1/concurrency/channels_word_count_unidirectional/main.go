package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const FILES_DIR = "./files"

func main() {
	example1()
}

func example1() {
	start := time.Now()
	intChannel := make(chan int)
	files := getFilesInDir(FILES_DIR)
	for _, file := range files {
		go printCountOfWordsWithChannel(file, intChannel)
	}

	waitForFinish(intChannel, len(files))
	elapsed := time.Since(start)
	fmt.Println("time taken:", elapsed)
}

func waitForFinish(intChannel <-chan int, numOfFiles int) {
	for i := 0; i < numOfFiles; i++ {
		<-intChannel
	}
}

func printCountOfWordsWithChannel(file fs.FileInfo, intChannel chan<- int) {

	content, err := os.ReadFile(FILES_DIR + "/" + file.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s has %d words\n", file.Name(), len(strings.Fields(string(content))))
	intChannel <- 0
}

func getFilesInDir(dir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []fs.FileInfo{}
	}
	return files
}
