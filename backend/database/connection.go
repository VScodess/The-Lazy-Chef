package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global client vairable
var Client *mongo.Client

func Connect() {

	Client = initializeClient()
	connectToMongoDB(Client)
	verifyConnection(Client);
	createTextIndex();

}

func initializeClient() *mongo.Client {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" // Fallback to localhost for development
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal("Failed to create MongoDB client: ", err)
	}
	return client
}

func connectToMongoDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}
}

func verifyConnection(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB: ", err)
	}

	log.Println("Connected to MongoDB!")
}

func createTextIndex() {
	collection := GetCollection("recipes")
	
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

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal("Failed to create text index: ", err)
	} else {
		log.Println("Text index created successfully on Name, Ingredients, Tags and Summary")
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("lazy_chef").Collection(collectionName)
}
