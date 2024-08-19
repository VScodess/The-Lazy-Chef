package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global client vairable
var Client *mongo.Client

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal("Failed to create MongoDB client: ", err)
	}

	//Establish connection to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	// Verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB: ", err)
	}

	log.Println("Connected to MongoDB!")

	Client = client
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("lazy_chef").Collection(collectionName)
}