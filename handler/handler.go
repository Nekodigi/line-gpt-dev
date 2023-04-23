package handler

import (
	"log"
	"net/http"

	"github.com/Nekodigi/line-gpt-dev/config"
	"github.com/Nekodigi/line-gpt-dev/handler/line/multi"
	"github.com/Nekodigi/line-gpt-dev/handler/line/simple"
	"github.com/Nekodigi/line-gpt-dev/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/sashabaranov/go-openai"
)

var (
	simpleLineClient *linebot.Client
	openaiClient     *openai.Client
	fs               *infrastructure.Firestore
)

func init() {
	var err error
	conf := config.Load()

	simpleLineClient, err = linebot.New(conf.SimpleLineSecret, conf.SimpleLineToken)
	if err != nil {
		log.Fatalf("simple_linebot.New: %+v", err)
	}
	openaiClient = openai.NewClient(conf.ChatGPTToken)

	fs, err = infrastructure.NewFirestore(conf.ProjectId)
	if err != nil {
		log.Fatalf("firestore.New: %+v", err)
	}
}

func Router(e *gin.Engine) {
	(&simple.Simple{Client: simpleLineClient, OpenAI: openaiClient}).Handle(e)
	(&multi.Multi{Client: simpleLineClient, OpenAI: openaiClient, Fs: fs}).Handle(e)
	e.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
}
