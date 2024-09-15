package load

import (
	"context"
	"job-funnel/types"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClientMongo *mongo.Client
var collectionMongo *mongo.Collection
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
		log.Fatal("error - MongoDB: Unable to ping database instance: %s", err)
	}

	collectionMongo = ClientMongo.Database(mongoDbName).Collection(mongoDbCollectionName)
}

func loadDbDataToMongoDb(data types.JobPost) error {
	filter := bson.M{"job_title": data.JobTitle}
	update := bson.M{
		"$set": bson.M{
			"job_title":        data.JobTitle,
			"description":      data.Description,
			"coding_language":  data.CodingLanguage,
			"coding_framework": data.CodingFramework,
			"database":         data.Database,
			"company_name":     data.CompanyName,
			"pay":              data.Pay,
			"work_location":    data.WorkLocation,
			"links":            data.Links,
			"created_at":       data.CreatedAt,
			"updated_at":       data.UpdatedAt,
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := collectionMongo.UpdateOne(CtxMongo, filter, update, opts)
	if err != nil {
		log.Printf("error - MongoDB: Database write failure: %s", err)
		return err
	}
	return nil
}
