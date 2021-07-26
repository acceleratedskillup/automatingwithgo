package twitter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"twitter/config"

	"github.com/dghubble/oauth1"
)

const tweetEndpoint = "https://api.twitter.com/2/tweets"

// TweetData represents the data structure for tweet request
type TweetData struct {
	Text string `json:"text"`
}

// Tweet posts a tweet using the Twitter API
func Tweet(content string, cfg config.Config) error {
	oauthConfig := oauth1.NewConfig(cfg.TwitterConsumerKey, cfg.TwitterConsumerSecret)
	token := oauth1.NewToken(cfg.TwitterAccessToken, cfg.TwitterAccessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := oauthConfig.Client(context.Background(), token)

	return postTweet(httpClient, content)
}

// postTweet sends a tweet request to the Twitter API
func postTweet(httpClient *http.Client, content string) error {
	data := TweetData{Text: content}
	tweetBody, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to encode tweet: %v", err)
	}

	req, err := http.NewRequest("POST", tweetEndpoint, bytes.NewBuffer(tweetBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to post tweet: %v, response: %s", resp.StatusCode, body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	fmt.Printf("Successfully tweeted: %s", body)
	return nil
}
