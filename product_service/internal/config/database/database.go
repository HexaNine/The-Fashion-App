package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func GetClient() *mongo.Client {
	
	if dbClient == nil {
		Connect()
	}
	
	return dbClient
}

func Connect()  {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URL")))
    
	if err != nil {
		log.Fatal("Error creating database client: ", err)
	}
	
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	log.Println("Successfully connected to MongoDB")

	dbClient = client
}