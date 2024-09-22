package retrieve

import (
	"context"
	"fmt"
	"job-funnel/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// queryMongoDb performs a query on the MongoDB collection and returns the results.
func retrieveDbFromMongoDbQuery(filter bson.M) ([]utils.JobPost, error) {
	var jobPosts []utils.JobPost

	// Set a context with a timeout for the query
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Perform the search query
	cursor, err := utils.CollectionMongo.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error - MongoDB: Unable to perform search query: %w", err)
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor and decode each document into a JobPost
	for cursor.Next(ctx) {
		var jobPost utils.JobPost
		if err := cursor.Decode(&jobPost); err != nil {
			return nil, fmt.Errorf("error - MongoDB: Unable to decode job post: %w", err)
		}
		jobPosts = append(jobPosts, jobPost)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error - MongoDB: Cursor error: %w", err)
	}

	return jobPosts, nil
}

// retrieveDbFromMongoDbAll retrieves all job posts from the MongoDB collection.
func retrieveDbFromMongoDbAll() ([]utils.JobPost, error) {
	resultJobPosts, err := retrieveDbFromMongoDbQuery(bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error - MongoDB: Unable to perform search query: %w", err)
	}
	return resultJobPosts, nil
}

// TODO: Need to review this function to ensure it is working as expected. I think it can be improved.
// SearchJobPostsInMongoDb searches for job posts in MongoDB based on a search term.
func retrieveDbFromMongoDbSearch(searchTerm string) ([]utils.JobPost, error) {
	// Define the filter for the search query
	filter := bson.M{
		"$or": []bson.M{
			{"job_id": bson.M{"$regex": searchTerm}},
			{"job_title": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"description": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"coding_language": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"coding_framework": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"database": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"company_name": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"pay": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"work_location": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"links": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"applied_to_job": bson.M{"$regex": searchTerm, "$options": "i"}},
		},
	}

	resultJobPosts, err := retrieveDbFromMongoDbQuery(filter)
	if err != nil {
		return nil, fmt.Errorf("error - MongoDB: Unable to perform search query: %w", err)
	}
	return resultJobPosts, nil
}
