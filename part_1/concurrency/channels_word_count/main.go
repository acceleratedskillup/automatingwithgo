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
	//example1()
	example2()
}

func example1() {
	start := time.Now()
	files := getFilesInDir(FILES_DIR)
	for _, file := range files {
		go printCountOfWords(file)
	}
	time.Sleep(time.Second)
	elapsed := time.Since(start)
	fmt.Println("time taken:", elapsed)
}

func example2() {
	start := time.Now()
	intChannel := make(chan int)
	files := getFilesInDir(FILES_DIR)
	for _, file := range files {
		go printCountOfWordsWithChannel(file, intChannel)
	}

	for i := 0; i < len(files); i++ {
		<-intChannel
	}
	elapsed := time.Since(start)
	fmt.Println("time taken:", elapsed)
}

func printCountOfWordsWithChannel(file fs.FileInfo, intChannel chan int) {

	content, err := os.ReadFile(FILES_DIR + "/" + file.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s has %d words\n", file.Name(), len(strings.Fields(string(content))))
	intChannel <- 0
}

func printCountOfWords(file fs.FileInfo) {

	content, err := os.ReadFile(FILES_DIR + "/" + file.Name())
	if err != nil {
		fmt.Println(err)
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
