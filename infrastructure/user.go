package infrastructure

import (
	"context"
	"fmt"

	"github.com/Nekodigi/line-gpt-dev/models"
	log "github.com/sirupsen/logrus"
)

func (fs *Firestore) GetUser(ID string) (*models.User, error) {
	ctx := context.Background()

	data, err := fs.c.Collection("chat_gpt").Doc("multi").Collection("users").Doc(ID).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user := &models.User{}
	if err := data.DataTo(user); err != nil {
		return nil, fmt.Errorf("data.DataTo: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("firestore Get: %w", err)
	}
	return user, nil
}

func (fs *Firestore) SetUser(user *models.User) error {
	ctx := context.Background()

	_, err := fs.c.Collection("chat_gpt").Doc("multi").Collection("users").Doc(user.Id).Set(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	return nil
}

func (fs *Firestore) CreateUser(ID string) (*models.User, error) {
	ctx := context.Background()

	user := models.NewUser(ID)
	_, err := fs.c.Collection("chat_gpt").Doc("multi").Collection("users").Doc(ID).Set(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("firestore Set: %w", err)
	}
	log.Infof("new user created: %s", ID)
	return user, nil
}
