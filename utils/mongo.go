package utils

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB variables
var ClientMongo *mongo.Client
var CollectionMongo *mongo.Collection
var CtxMongo = context.TODO()

func LoadDbConnectToMongoDb() {
	mongoDbUri := os.Getenv("DB_MONGODB_URI")
	mongoDbName := os.Getenv("DB_NAME")
	mongoDbCollectionName := os.Getenv("DB_TABLE_NAME")
	clientOptions := options.Client().ApplyURI(mongoDbUri)
	ClientMongo, err := mongo.Connect(CtxMongo, clientOptions)
	if err != nil {
		log.Fatalf("error - MongoDB: Unable to establish database connection: %s", err)
	}
	// This reduces the codes resilience to failure. So we may want to remove this.
	err = ClientMongo.Ping(CtxMongo, nil)
	if err != nil {
		log.Fatalf("error - MongoDB: Unable to ping database instance: %s", err)
	}

	CollectionMongo = ClientMongo.Database(mongoDbName).Collection(mongoDbCollectionName)
}
