package load

import (
	"job-funnel/utils"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadDbConnectToMongoDb() {
	mongoDbUri := os.Getenv("DB_MONGODB_URI")
	mongoDbName := os.Getenv("DB_NAME")
	mongoDbCollectionName := os.Getenv("DB_TABLE_NAME")
	clientOptions := options.Client().ApplyURI(mongoDbUri)
	ClientMongo, err := mongo.Connect(utils.CtxMongo, clientOptions)
	if err != nil {
		log.Fatalf("error - MongoDB: Unable to establish database connection: %s", err)
	}
	// This reduces the codes resilience to failure. So we may want to remove this.
	err = ClientMongo.Ping(utils.CtxMongo, nil)
	if err != nil {
		log.Fatalf("error - MongoDB: Unable to ping database instance: %s", err)
	}

	utils.CollectionMongo = ClientMongo.Database(mongoDbName).Collection(mongoDbCollectionName)
}

func loadDbDataToMongoDb(data utils.JobPost, jobId string) error {
	filter := bson.M{"job_title": data.JobTitle}
	update := bson.M{
		"$set": bson.M{
			"job_source":       data.JobSource,
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
		"$setOnInsert": bson.M{
			"job_id":         jobId,
			"applied_to_job": []string{},
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := utils.CollectionMongo.UpdateOne(utils.CtxMongo, filter, update, opts)
	if err != nil {
		log.Println("error - MongoDB: Database write failure: %s", err)
		return err
	}
	return nil
}
