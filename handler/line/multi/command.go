package multi

import (
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func GetSimple(text string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		},
	}
}

func GetTranslate(text string, langTo string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: fmt.Sprintf("Translate this into %s. Add pinyin below if Chinese:", langTo),
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		},
	}
}
