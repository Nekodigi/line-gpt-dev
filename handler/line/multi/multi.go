package multi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Nekodigi/line-gpt-dev/consts"
	"github.com/Nekodigi/line-gpt-dev/infrastructure"
	"github.com/Nekodigi/line-gpt-dev/models"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

type (
	Multi struct {
		Client *linebot.Client
		OpenAI *openai.Client
		Fs     *infrastructure.Firestore
	}
)

func (l *Multi) Handle(e *gin.Engine) {
	e.POST("/multi", func(c *gin.Context) {
		log.Infof("/callback called")

		if e := l.callback(c.Request); e != nil {
			log.Errorf("line callback: %+v", e)
		}
		c.Status(http.StatusOK)
	})

}

func (l *Multi) callback(r *http.Request) error {
	events, err := l.Client.ParseRequest(r)
	if err != nil {
		return fmt.Errorf("line ParseRequest: %w", err)
	}

	for _, event := range events {
		var err error

		user, err := l.Fs.GetUser(event.Source.UserID)
		if err != nil {
			user, err = l.Fs.CreateUser(event.Source.UserID)
		} else {

		}
		//action := l.react(event, pd, member)
		res := "メッセージがありません"

		if event.Type == linebot.EventTypeFollow {
			res = "友達登録ありがとうございます！ただいま会話モードになっています。翻訳、要約、励ましと話しかけるとモードが切り替わる他、カスタムと話しかけるとカスタムコマンドを作成できます。"
		}
		if event.Type == linebot.EventTypeMessage {
			switch msg := event.Message.(type) {
			case *linebot.TextMessage:
				res = l.reactText(msg.Text, user)
			}

		}

		if _, err2 := l.Client.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(res),
		).Do(); err2 != nil {
			log.Errorf("error reply: %+v", err2)
		}
	}
	return nil
}

func (l *Multi) reactText(text string, user *models.User) string {
	res := "メッセージがありません"
	switch user.Status {
	case consts.StatusIdle:
		switch text {
		case "カスタム":
			user.Status = consts.StatusAskCommandName
			res = "カスタムコマンド名を入力してください。"
		default:
			command, err := l.Fs.GetCommand(text)
			if err != nil {
				res = l.askChatGPT(text, user).Choices[0].Message.Content
			} else {
				user.ChatType = command.Id
				user.Args = []string{}
				res = fmt.Sprintf("「%s」モードになりました。", command.Id)
				if len(command.Args) != 0 {
					user.Status = consts.StatusAskArgValues

					res = fmt.Sprintf("「%s」モードになりました。「%s」を入力してください。", command.Id, command.Args[0])
				}
			}
			//TODO Get command if not hit ask.

		}
	case consts.StatusAskCommandName:
		if text == "カスタム" {
			user.Status = consts.StatusIdle
			res = "その名前は使用できません。"
		} else {
			l.Fs.CreateCommand(text)
			user.WorkingCommand = text
			user.Status = consts.StatusAskCommand
			res = `命令文を入力してください。ex)[{"Role": "user", "Content": "%s"}]`
		}
	case consts.StatusAskCommand:
		fmt.Println("ask command")

		command, _ := l.Fs.GetCommand(user.WorkingCommand)
		messages := []openai.ChatCompletionMessage{}
		err := json.Unmarshal([]byte(text), &messages)
		if err != nil {
			user.Status = consts.StatusAskCommand
			res = "命令文が不正です。もう一度命令文を入力してください。"
		} else {
			command.Messages = text
			fmt.Printf("Set command: %v", messages)
			l.Fs.SetCommand(command)
			user.Status = consts.StatusAskArgs
			res = `引数を入力してください。ex)["言語"]`
		}
	case consts.StatusAskArgs:
		fmt.Println("ask args")
		command, _ := l.Fs.GetCommand(user.WorkingCommand)
		args := []string{}
		err := json.Unmarshal([]byte(text), &args)
		if err != nil {
			user.Status = consts.StatusAskArgs
			res = "引数が不正です。もう一度引数を入力してください。"
		} else {
			command.Args = args
			fmt.Printf("Set args: %v", args)
			l.Fs.SetCommand(command)
			user.Status = consts.StatusIdle
			res = fmt.Sprintf("命令「%s」の作成が完了しました。", command.Id)
		}
	case consts.StatusAskArgValues:
		command, _ := l.Fs.GetCommand(user.ChatType)
		user.Args = append(user.Args, text)
		if len(user.Args) == len(command.Args) {
			user.Status = consts.StatusIdle
			res = "引数の入力が完了しました。"
		} else {
			res = fmt.Sprintf("「%s」を入力してください。", command.Args[len(user.Args)])
		}
	case consts.StatusAskLangTo:
		user.LangTo = text
		user.Status = consts.StatusIdle
		res = fmt.Sprintf("翻訳モード(%s)になりました", text)
	}
	l.Fs.SetUser(user)
	return res
}

func (l *Multi) askChatGPT(text string, user *models.User) openai.ChatCompletionResponse {
	command, err := l.Fs.GetCommand(user.ChatType)

	if err != nil {
		fmt.Printf("GetCommand error: %v\n", err)
	}

	// message := GetSimple(text)
	// switch user.ChatType {
	// case consts.TypeTranslate:
	// 	message = GetTranslate(text, user.LangTo)
	// }

	args_ := append(user.Args, text)
	args := make([]interface{}, len(args_))
	for i, v := range args_ {
		args[i] = v
	}

	embededMessage := fmt.Sprintf(command.Messages, args...)
	fmt.Println(embededMessage)

	embededMessage = strings.Replace(embededMessage, "\n", " ", -1)
	messages := []openai.ChatCompletionMessage{}
	err = json.Unmarshal([]byte(embededMessage), &messages)
	if err != nil {
		fmt.Printf("Unmarshal error: %v\n", err)
	}

	fmt.Println(messages)

	resp, err := l.OpenAI.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}
	fmt.Println(resp)
	return resp
}
