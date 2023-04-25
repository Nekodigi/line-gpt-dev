package models

type User struct {
	Id             string
	Args           []string
	ChatType       string
	Status         string
	LangFrom       string
	LangTo         string
	WorkingCommand string
}

func NewUser(ID string) *User {
	return &User{
		Id:             ID,
		Args:           []string{},
		Status:         "idle",
		ChatType:       "会話",
		LangFrom:       "jp",
		LangTo:         "en",
		WorkingCommand: "",
	}
}
