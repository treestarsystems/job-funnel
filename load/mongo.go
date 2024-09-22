package load

import (
	"job-funnel/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
		log.Printf("error - MongoDB: Database write failure: %s\n", err)
		return err
	}
	return nil
}
