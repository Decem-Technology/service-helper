package bootstrap

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var fApp *firebase.App
var fDBClient *db.Client

// Firebase firebase management
type Firebase struct {
}

// CreateFirebaseConnection init firebase
func CreateFirebaseConnection() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("storage/firebase-service-account.json")
	conf := &firebase.Config{
		DatabaseURL: os.Getenv("FIREBASE_DATABASE_URL"),
	}
	var err error
	fApp, err = firebase.NewApp(ctx, conf, opt)
	if err != nil {
		panic(fmt.Sprintf("[Firebase] Error initializing app: %v", err))
	}

	fDBClient, err = fApp.Database(ctx)
	if err != nil {
		panic(fmt.Sprintf("[Firebase] Error initializing database client: %v", err))
	}

	fmt.Println("[Firebase] connected")

}

// Firebase get firebase
func (ctl *Firebase) Firebase() *firebase.App {
	return fApp
}

// RealTimeDB get firebase realtime db
func (ctl *Firebase) RealTimeDB() *db.Client {
	return fDBClient
}
