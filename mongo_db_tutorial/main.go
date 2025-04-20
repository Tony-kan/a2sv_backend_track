package main

import (
	// "context"
	// "fmt"

	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	collection := client.Database("test").Collection("trainers")

	// collection := client.Database("test").Collection("trainers")
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 12, "Cerun Town"}
	brock := Trainer{"Brock", 15, "Pewter Town"}

	// Data insertion to db
	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), ash)

	// insertResult,err :=
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted ash document : ", insertResult.InsertedID)

	// Insert many documents
	trainers := []interface{}{misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted trainers : ", insertManyResult.InsertedIDs)

	// update document
	filter := bson.D{{"name", "Ash"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Matched %v documents and updated %v documents. \n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// find documents

	// create a value into which the result can be decoded
	var result Trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found a single document : %+v\n", result)

	// finding multiple documents
	findOptions := options.Find()
	findOptions.SetLimit(2)

	// Here's an array in which you can store the decoded documents
	var results []*Trainer

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	// finding multiple documents returns a cursor
	// iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// close the curso once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %v\n", results)

	// delete documents
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// disconnect
	//err = client.Disconnect(context.TODO())
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("Connection to Mongodb closed.")
}
