package utils

import (
	"context"

	gogpt "github.com/sashabaranov/go-openai"
)

func GetEmbeddings(client *gogpt.Client, ctx context.Context, textRow []string) (*gogpt.EmbeddingResponse, error) {
	req := gogpt.EmbeddingRequest{
		Model: gogpt.AdaEmbeddingV2,
		Input: textRow,
	}

	res, err := client.CreateEmbeddings(ctx, req)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
