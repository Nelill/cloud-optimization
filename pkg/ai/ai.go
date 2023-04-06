package ai

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	client *openai.Client
}

func (c *OpenAIClient) Configure(token string) error {
	client := openai.NewClient(token)
	if client == nil {
		return errors.New("error creating openai client.")
	}
	c.client = client
	return nil
}

func (c *OpenAIClient) GetCompletion(ctx context.Context, prompt string) (string, error) {
	resp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: "user",
				Content: fmt.Sprintf("Analyze the following Google Cloud Compute instance and suggest cost optimizations: %s", prompt),
			},
		},
	})

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
	
var aiClient OpenAIClient

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		panic("OPENAI_API_KEY environment variable not set")
	}

	err := aiClient.Configure(apiKey)
	if err != nil {
		panic(fmt.Sprintf("Failed to configure AI client: %v", err))
	}
}

func RequestGPTSuggestions(instanceInfo string) (string, error) {

	ctx := context.Background()
	return aiClient.GetCompletion(ctx, instanceInfo)
}
