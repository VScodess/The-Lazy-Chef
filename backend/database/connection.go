package database

import (
	"context"
	"log"
	"time"

	"The-Lazy-Chef/backend/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global client variable
var Client *mongo.Client
var dbName string
var collectionName string

// Connect initializes the MongoDB connection and sets up the text index
func Connect() {
	cfg := config.LoadConfig()
	dbName = cfg.DatabaseName
	collectionName = cfg.CollectionName
	Client = initializeClient(cfg.MongoURI)
	connectToMongoDB(Client)
	verifyConnection(Client)
	createTextIndex()
}

// initializeClient creates a new MongoDB client with the provided URI
func initializeClient(mongoURI string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Failed to create MongoDB client: ", err)
	}
	return client
}

// connectToMongoDB establishes the connection to MongoDB
func connectToMongoDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}
}

// verifyConnection pings MongoDB to ensure the connection is successful
func verifyConnection(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB: ", err)
	}

	log.Println("Connected to MongoDB!")
}

// createTextIndex creates a text index on specific fields in the recipes collection
func createTextIndex() {
	collection := GetCollection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "name", Value: "text"},
			{Key: "ingredients", Value: "text"},
			{Key: "tags", Value: "text"},
			{Key: "summary", Value: "text"},
		},
		Options: options.Index().SetDefaultLanguage("english"),
	}

	if _, err := collection.Indexes().CreateOne(ctx, indexModel); err != nil {
		log.Fatal("Failed to create text index: ", err)
	} else {
		log.Println("Text index created successfully on Name, Ingredients, Tags, and Summary")
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database(dbName).Collection(collectionName)
}
