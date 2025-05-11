package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://<username>:<password>@cluster0.mongodb.net/golang?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// NOTE: IMPORTANT: This is the client instance that will be used to connect to the MongoDB database
var collection *mongo.Collection

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	checkError(err)

	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(colName) // getting a handle for the collection
	fmt.Println("Collection instance is ready")
}
