package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBSet() *mongo.Client{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err :=  mongo.Connect(
		ctx,
        options.Client().ApplyURI("mongodb://localhost:27017"),
	)

	if err != nil {
		log.Fatal(err)
	}

	// defer func ()  {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		log.Fatal(err)
			
	// 	}
	// }()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("failed to connect to mongodb")
		return nil
	}
	fmt.Println("Successfully connected to mongodb")

	return client

	
}

func UserData(client *mongo.Client, CollectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return collection

}

func ProductData(client *mongo.Client, CollectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return collection
}