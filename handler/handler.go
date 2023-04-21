package handler

import (
	"log"

	"github.com/Nekodigi/line-gpt-dev/config"
	"github.com/Nekodigi/line-gpt-dev/handler/line/simple"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/sashabaranov/go-openai"
)

var (
	simpleLineClient *linebot.Client
	openaiClient     *openai.Client
)

func init() {
	var err error
	conf := config.Load()

	simpleLineClient, err = linebot.New(conf.SimpleLineSecret, conf.SimpleLineToken)
	if err != nil {
		log.Fatalf("simple_linebot.New: %+v", err)
	}
	openaiClient = openai.NewClient(conf.ChatGPTToken)

}

func Router(e *gin.Engine) {
	(&simple.Simple{Client: simpleLineClient, OpenAI: openaiClient}).Handle(e)
}
