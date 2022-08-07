package main

import (
	"context"
	"log"

	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbClient *mongo.Client
	dbOnce   sync.Once
)

func getDatabase(dbName string) *mongo.Database {

	dbOnce.Do(func() {
		// Creating mongodb client instance
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb:27017/"))
		if err != nil {
			log.Fatal(err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		dbClient = client
	})

	return dbClient.Database(dbName)
}

func getTransactionCollection(database string) (*mongo.Collection, context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	transactionsDatabase := getDatabase(database)
	transactionsCollection := transactionsDatabase.Collection("transactions")

	return transactionsCollection, ctx
}

func clearTransactionCollection(database string) {
	transactionsCollection, ctx := getTransactionCollection(database)

	err := transactionsCollection.Drop(ctx)
	if err != nil {
		panic(err)
	}
}
