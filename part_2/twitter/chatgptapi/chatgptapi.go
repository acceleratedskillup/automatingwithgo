package chatgptapi

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

const (
	chatgptPrompt = `Summarize the following article for a Twitter post, including the original URL in the tweet: %s. The summary, including the URL, should be concise, engaging, and fit within Twitter's 280-character limit. Highlight the main points and include a call to action that encourages readers to click on the link for more information. Assume the URL will use approximately 23 characters.`
)

// CreateTweetContent sends data to ChatGPT to generate tweet content.
func CreateTweetContent(apiKey, articleUrl string) (string, error) {
	// Prepare the prompt for GPT
	prompt := fmt.Sprintf(chatgptPrompt, articleUrl)
	//fmt.Println("prompt:", prompt)
	// Create a new OpenAI client
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("error in ChatCompletion: %v", err)
	}

	if len(resp.Choices) == 0 || resp.Choices[0].Message.Content == "" {
		return "", fmt.Errorf("no completion found or empty response")
	}

	return resp.Choices[0].Message.Content, nil
}
