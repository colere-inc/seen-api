package infrastructure

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"

	"github.com/colere-inc/seen-api/app/common/config"
)

type DB = firestore.Client

func NewDB() *DB {
	app := newFirebaseApp()
	db, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func newFirebaseApp() *firebase.App {
	app, err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: config.ProjectID})
	if err != nil {
		log.Fatalln(err)
	}
	return app
}
