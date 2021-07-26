package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const maxDownloads = 5

func main() {
	rssURL := "https://feeds.simplecast.com/qm_9xx0g" // Replace with the RSS feed URL of the podcast you want to download

	resp, err := http.Get(rssURL)
	if err != nil {
		fmt.Println("Error fetching RSS feed:", err)
		return
	}
	defer resp.Body.Close()

	var rss Rss
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Println("Error decoding RSS feed:", err)
		return
	}
	start := time.Now()
	for index, item := range rss.Channel.Items {
		if index >= maxDownloads {
			break
		}
		downloadPodcast(item.Enclosure.URL, item.Title+".mp3")
	}
	elapsed := time.Since(start)
	fmt.Printf("The code took %s to run.\n", elapsed)
}

func downloadPodcast(url, filename string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading podcast:", err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Downloaded:", filename)
}
