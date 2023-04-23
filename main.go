package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Nekodigi/line-gpt-dev/handler"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type Request struct {
	Operation string `json:"operation"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func main() {
	//handler.Firestore()
	if len(os.Args) == 2 && os.Args[1] == "test" {
		// handler.Firestore()
		//msg := Request{}
		messages := []openai.ChatCompletionMessage{}
		strs := []string{}
		//t := []openai.ChatCompletionMessage{{Role: "system", Content: "test"}}
		values := []interface{}{"system", "test"}
		err := json.Unmarshal([]byte(fmt.Sprintf(`[{"Role": "%s", "Content": "%s"}]`, values...)), &messages)
		json.Unmarshal([]byte(`["apple", "banana"]`), &strs)
		fmt.Println(strs)
		if err != nil {
			fmt.Printf("Invalid command: %v", err)
		}
		fmt.Printf("Operation: %s\n", messages)
	} else {
		engine := gin.Default()
		handler.Router(engine)
		engine.Run(":8080")
	}
}
