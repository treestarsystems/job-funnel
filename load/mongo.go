package load

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func LoadDbConnectToMongoDb() {
	mongoDbUri := os.Getenv("DB_MONGO_URI")
	mongoDbName := os.Getenv("DB_NAME")
	mongoDbCollectionName := os.Getenv("DB_TABLE_NAME")
	clientOptions := options.Client().ApplyURI(mongoDbUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(mongoDbName).Collection(mongoDbCollectionName)
}
