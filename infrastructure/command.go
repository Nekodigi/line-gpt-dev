package infrastructure

import (
	"context"
	"fmt"

	"github.com/Nekodigi/line-gpt-dev/models"
	log "github.com/sirupsen/logrus"
)

func (fs *Firestore) GetCommand(ID string) (*models.Command, error) {
	ctx := context.Background()

	data, err := fs.c.Collection("chat_gpt").Doc("multi").Collection("commands").Doc(ID).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get message: %w", err)
	}

	command := &models.Command{}
	if err := data.DataTo(command); err != nil {
		return nil, fmt.Errorf("data.DataTo: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("firestore Get: %w", err)
	}
	return command, nil
}

func (fs *Firestore) SetCommand(command *models.Command) error {
	ctx := context.Background()

	_, err := fs.c.Collection("chat_gpt").Doc("multi").Collection("commands").Doc(command.Id).Set(ctx, command)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	return nil
}

func (fs *Firestore) CreateCommand(ID string) (*models.Command, error) {
	ctx := context.Background()

	command := models.NewCommand(ID)
	_, err := fs.c.Collection("chat_gpt").Doc("multi").Collection("commands").Doc(ID).Set(ctx, command)
	if err != nil {
		return nil, fmt.Errorf("firestore Set: %w", err)
	}
	log.Infof("new user created: %s", ID)
	return command, nil
}
