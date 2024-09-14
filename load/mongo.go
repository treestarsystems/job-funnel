package load

import (
	"context"
	"job-funnel/transform"
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
	// mongoDbName := os.Getenv("DB_NAME")
	// mongoDbCollectionName := os.Getenv("DB_TABLE_NAME")
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

	// collection = client.Database(mongoDbName).Collection(mongoDbCollectionName)
	collectionMongo = ClientMongo.Database("jobs").Collection("job_posts")
}

func loadDbDataToMongoDb(data transform.JobPost) error {

	// filter := bson.M{"jobTitle": data.JobTitle}
	update := bson.M{
		"$set": bson.M{
			"job_title":       data.JobTitle,
			"description":     data.Description,
			"coding_language": data.CodingLanguage,
			"database":        data.Database,
			"company_name":    data.CompanyName,
			"pay":             data.Pay,
			"location":        data.Location,
			"links":           data.Links,
			"created_at":      data.CreatedAt,
			"updated_at":      data.UpdatedAt,
		},
	}
	// opts := options.Update().SetUpsert(true)

	// _, err := collectionMongo.UpdateOne(CtxMongo, filter, update, opts)
	result, err := collectionMongo.InsertOne(CtxMongo, update)
	if err != nil {
		return err
	}
	log.Printf("MongoDB: Upserted job post with ID: %s", result.InsertedID)
	return nil
}
