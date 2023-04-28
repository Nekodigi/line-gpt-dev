package main

import (
	"os"

	"github.com/Nekodigi/line-gpt-dev/config"
	"github.com/Nekodigi/line-gpt-dev/handler"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Operation string `json:"operation"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func main() {
	//handler.Firestore()
	if len(os.Args) == 2 && os.Args[1] == "setup" {
		config.Setup()
	} else {
		engine := gin.Default()
		handler.Router(engine)
		engine.Run(":8080")
	}
}
