package firebase

import (
	"context"
	"fmt"
	firebase "firebase.google.com/go/v4"
)

type FirebaseInterface interface {
	GetFirebase() *firebase.App
}

type FirebaseApp struct {
	app *firebase.App
}

func Init() (FirebaseInterface, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase admin: %v", err)
	}

	return &FirebaseApp{
		app: app,
	}, nil
}

func (fc *FirebaseApp) GetFirebase() *firebase.App {
	return fc.app
}