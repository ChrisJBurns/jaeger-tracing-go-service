package main

import (
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employee struct {
    FirstName string
    LastName  string
    Occupation string
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin@localhost:27017/admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("test").Collection("employees")

	ash := Employee{"Ash", "Ketchum", "Pokemon Trainer"}
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	var result Employee
	err = collection.FindOne(context.TODO(), bson.D{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Documents %s: ", result)
}