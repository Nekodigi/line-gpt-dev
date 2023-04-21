package simple

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

type (
	Simple struct {
		Client *linebot.Client
		OpenAI *openai.Client
	}
)

func (s *Simple) Handle(e *gin.Engine) {
	e.POST("/simple", func(c *gin.Context) {
		log.Infof("/callback called")

		if e := s.callback(c.Request); e != nil {
			log.Errorf("line callback: %+v", e)
		}
		c.Status(http.StatusOK)
	})

	e.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
}

func (s *Simple) callback(r *http.Request) error {
	events, err := s.Client.ParseRequest(r)
	if err != nil {
		return fmt.Errorf("line ParseRequest: %w", err)
	}

	for _, event := range events {
		var err error

		if err != nil {
			//first follow
		} else {

		}
		//action := l.react(event, pd, member)

		if event.Type == linebot.EventTypeMessage {
			switch msg := event.Message.(type) {
			case *linebot.TextMessage:
				resp, err := s.OpenAI.CreateChatCompletion(
					context.Background(),
					openai.ChatCompletionRequest{
						Model: openai.GPT3Dot5Turbo,
						Messages: []openai.ChatCompletionMessage{
							{
								Role:    openai.ChatMessageRoleUser,
								Content: msg.Text,
							},
						},
					},
				)

				if err != nil {
					fmt.Printf("ChatCompletion error: %v\n", err)
				}

				if _, err2 := s.Client.ReplyMessage(
					event.ReplyToken,
					linebot.NewTextMessage(resp.Choices[0].Message.Content),
				).Do(); err2 != nil {
					log.Errorf("error reply: %+v", err2)
				}
			}

		}
	}
	return nil
}
