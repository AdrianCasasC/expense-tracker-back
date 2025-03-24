package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDB initializes MongoDB connection
func ConnectDB() *mongo.Database {
	// Retrieve MongoDB URI and Database Name
	mongoURI := os.Getenv("MONGO_URI_CONNECTION_STRING")
	dbName := os.Getenv("MONGO_DATABASE_NAME")

	// Create a new client and connect to MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping error: ", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Select Database
	DB = client.Database(dbName)
	return DB
}

// GetCollection returns a reference to a specific collection
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
