package googlealerts

import (
	"fmt"
	"net/url"

	"github.com/mmcdole/gofeed"
)

// FetchRSSFeed fetches and parses an RSS feed from a given URL.
func FetchRSSFeed(inputUrl string) ([]*gofeed.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(inputUrl)
	if err != nil {
		return nil, err
	}

	return feed.Items, nil
}

// GetURLParam extracts a specific query parameter from a given URL.
func GetURLParam(inputUrl, paramKey string) (string, error) {
	// Parse the URL
	parsedURL, err := url.Parse(inputUrl)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return "", err
	}

	// Extract query parameters
	queryParams := parsedURL.Query()

	// Get the value of the specified parameter
	paramValue := queryParams.Get(paramKey)

	fmt.Printf("Value of %s: %s\n", paramKey, paramValue)
	return paramValue, nil
}
