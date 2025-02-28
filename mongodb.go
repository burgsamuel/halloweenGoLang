package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbPassword string = os.Getenv("MONGOPASSWORD")

var uri string = "mongodb+srv://admin:"+ dbPassword + "@cluster0.r2jvy.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"





func InsertSpot(spot Spot){

	// Set client options and connect to MongoDB Atlas
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.TODO())

    collection := client.Database("halloween").Collection("locations")

    insertResult, err := collection.InsertOne(context.TODO(), spot)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted location with ID: %v\n", insertResult.InsertedID)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

func RetrieveMarkers() ([]bson.M, error){

		// Set client options and connect to MongoDB Atlas
		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}
		defer client.Disconnect(context.TODO())

		collection := client.Database("halloween").Collection("locations")

		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return []bson.M{}, nil 
			}
		}
		defer cursor.Close(context.TODO())

		var results []bson.M	

		for cursor.Next(context.TODO()) {
			var result bson.M
			err := cursor.Decode(&result)
			if err != nil {
				return nil, err
			}
			results = append(results, result)
		}
		if err != nil {
			return nil, err
		}
		return results, nil
	}