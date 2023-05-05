package infrastructure

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type DB = firestore.Client

func NewDB(f *firebase.App) *DB {
	db, err := f.Firestore(context.Background())
	if err != nil {
		panic(err)
	}
	return db
}
