package app

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewDB() *firestore.Client {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("tugas-akhir-ea222-firebase-adminsdk-fbsvc-57a243d81f.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
