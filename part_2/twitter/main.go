package main

import (
	"fmt"
	"os"
	"twitter/chatgptapi"
	"twitter/config"
	"twitter/filedb"
	"twitter/googlealerts"
	"twitter/twitter"

	"github.com/mmcdole/gofeed"
)

const (
	rssURL = "https://www.google.com/alerts/feeds/03498191070354501431/621352783961875540"
)

var (
	postedItemsFilePath = "posted_items.txt"
	iniFile             = "config.ini"
)

func main() {
	// Allow overriding file paths from environment variables
	if envFilePath := os.Getenv("POSTED_ITEMS_FILE_PATH"); envFilePath != "" {
		postedItemsFilePath = envFilePath
	}
	if envIniFile := os.Getenv("INI_FILE_PATH"); envIniFile != "" {
		iniFile = envIniFile
	}

	cfg, err := config.ReadConfig(iniFile)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}

	items, err := googlealerts.FetchRSSFeed(rssURL)
	if err != nil {
		fmt.Println("Error fetching RSS feed:", err)
		return
	}

	if len(items) > 0 {
		processRSSItem(items[0], cfg)
	}
}

func processRSSItem(item *gofeed.Item, cfg config.Config) {
	alreadyPosted, err := filedb.IsAlreadyPosted(item.Title, postedItemsFilePath)
	if err != nil {
		fmt.Printf("Error checking if item is already posted:%v\n", err)
		return
	}

	if alreadyPosted {
		fmt.Println("Item already posted:", item.Title)
		return
	}

	originalUrl, err := googlealerts.GetURLParam(item.Link, "url")
	if err != nil {
		fmt.Printf("Error getting original url:%v\n", err)
		return
	}

	tweetContent, err := chatgptapi.CreateTweetContent(cfg.OpenAiApiKey, originalUrl)
	if err != nil {
		fmt.Println("Error creating tweet content:", err)
		return
	}
	fmt.Println("tweetContent:", tweetContent)
	err = twitter.Tweet(tweetContent, cfg)
	if err != nil {
		fmt.Printf("Error posting tweet:%v\n", err)
		return
	}
	fmt.Println("Article tweeted successfully")

	err = filedb.RecordPostedItem(item.Title, postedItemsFilePath)
	if err != nil {
		fmt.Printf("Error recording posted tweet:%v\n", err)
		return
	}
	fmt.Println("Item recorded in filedb successfully")
}
