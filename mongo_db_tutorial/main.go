package main

import (
	// "context"
	// "fmt"

	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/v2/x/mongo/driver/mongocrypt/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	godotenv.Load()
	mongo_uri := os.Getenv("MONGO_URI")


	// set client options
	clientOptions := options.Client().ApplyURI(mongo_uri)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	// collection := client.Database("test").Collection("trainers")

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to Mongodb closed.")
}
