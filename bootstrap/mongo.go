package bootstrap

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	// MongoDB mongodb database management
	MongoDB struct {
	}
)

// dbMongo variable for define connection
var dbMongo *mongo.Client

// CreateMongoConnection make connection
func CreateMongoConnection() *mongo.Client {
	connectionString := os.Getenv("MONGODB_CONNECTION")
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(fmt.Sprintf("[MongoDB] Initial connection fail, error: %s", err))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(fmt.Sprintf("[MongoDB] connect database fail, error: %s", err))
	}
	// defer client.Disconnect(ctx)
	fmt.Println("[MongoDB] connected")

	dbMongo = client
	return client
}

// DB get mongo connection
func (c *MongoDB) DB(db string, opt ...*options.DatabaseOptions) *mongo.Database {
	return dbMongo.Database(db, opt...)
}
