package controllers

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func RecommendedLaptop(input string) (string, error) {

	ctx := context.Background()
	client := openai.NewClient(os.Getenv("api_key"))
	model := openai.GPT3Dot5Turbo

	messages := []openai.ChatCompletionMessage{{
		Role:    openai.ChatMessageRoleSystem,
		Content: "Anda adalah AI Laptop Recommendation. Berikan Respon menggunakan kalimat 'With a budget of Rp <budget>, you can get a <purpose> laptop with <spec> that can handle the ... . Here's a recommendation for a <brand> laptop within ur budget ...'",
	},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		},
	}

	resp, err := GetCompletionFromMessage(ctx, client, messages, model)
	if err != nil {
		return "", err
	}

	answer := resp.Choices[0].Message.Content
	return answer, nil
}

func GetCompletionFromMessage(ctx context.Context, client *openai.Client, messages []openai.ChatCompletionMessage, model string) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx, openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)

	return resp, err
}
