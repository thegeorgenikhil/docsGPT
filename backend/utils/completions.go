package utils

import (
	"context"
	"log"

	gogpt "github.com/sashabaranov/go-openai"
	"github.com/thegeorgenikhil/docsGPT/models"
)

const (
	MAX_TOKEN      = 200
	PREAMBLE_BEGIN = "You are a chat support bot. Your goal is to be chipper, cheerful and helpful. You have been. You have been given all the necessary documents which are stored in a database to help you answer user queries. It has been determined that the user is asking you a query. After searching the database, we have found out three relevant sections of the document that might help you to answer the user's question.\nThe three relevant sections are:\n"
	PREAMBLE_END   = "\nNow the user is asking a unique question we haven't seen before.Using the above reference material, craft them the best answer you can. If you don't think the above references give a good answer, simply tell the user you don't know how to help them.\n\n The query asked by the user is: "
)

func GetGPTResponse(client *gogpt.Client, ctx context.Context, prompt string) (*string, error) {
	req := gogpt.CompletionRequest{
		Prompt:      prompt,
		MaxTokens:   MAX_TOKEN,
		Temperature: 0,
		Model:       gogpt.GPT3TextDavinci003,
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &resp.Choices[0].Text, nil
}

func BuildPrompt(query string, documents []models.Content) string {
	prompt := PREAMBLE_BEGIN
	for _, document := range documents {
		prompt += document.Content + "\n"
	}
	prompt += PREAMBLE_END + query
	return prompt
}
