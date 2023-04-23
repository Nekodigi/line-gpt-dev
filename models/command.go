package models

type Command struct {
	Id       string
	Args     []string
	Messages string
}

func NewCommand(ID string) *Command {
	return &Command{
		Id:       ID,
		Args:     []string{},
		Messages: "",
	}
}
