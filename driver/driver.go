package driver

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"log"
)

type DB struct {
	Client *db.Client
}

func GetDBClient() (DB, error) {
	cxt := context.Background()

	conf := &firebase.Config {
		DatabaseURL: "https://dev-palm.firebaseio.com",
	}

	opt := option.WithCredentialsFile("driver/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), conf, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Database(cxt)

	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	return DB{ Client: client }, err
}
