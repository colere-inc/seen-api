package infrastructure

import (
	"context"

	"github.com/colere-inc/seen-api/app/common/config"

	firebase "firebase.google.com/go"
)

func NewFirebase() *firebase.App {
	f, err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: config.ProjectID})
	if err != nil {
		panic(err)
	}
	return f
}
