package infrastructure

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

var (
	fs *Firestore
)

type (
	Firestore struct {
		app *firebase.App
		c   *firestore.Client
	}
)

func NewFirestore(projectID string) (*Firestore, error) {
	if fs == nil {
		conf := &firebase.Config{ProjectID: projectID}

		app, err := firebase.NewApp(context.Background(), conf)

		if err != nil {
			return nil, fmt.Errorf("firebase.NewApp: %w", err)
		}

		fc, err := app.Firestore(context.Background())
		if err != nil {
			return nil, fmt.Errorf("fs.app.Firestore: %w", err)
		}

		fs = &Firestore{
			app: app,
			c:   fc,
		}
	}
	return fs, nil
}
