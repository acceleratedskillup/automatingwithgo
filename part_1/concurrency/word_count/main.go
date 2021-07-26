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
	start := time.Now()
	files := getFilesInDir(FILES_DIR)
	for _, file := range files {
		printCountOfWords(file)
	}
	elapsed := time.Since(start)
	fmt.Println("time taken:", elapsed)
}

func printCountOfWords(file fs.FileInfo) {

	content, err := os.ReadFile(FILES_DIR + "/" + file.Name())
	if err != nil {
		fmt.Println(err)
	}
	words := strings.Fields(string(content))
	fmt.Printf("%s has %d words\n", file.Name(), len(words))

}

func getFilesInDir(dir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []fs.FileInfo{}
	}
	return files
}
