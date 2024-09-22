package retrieve

import (
	"job-funnel/utils"
	"log"
	"os"
)

// RetrieveDbDataAll is wrapper for MongoDB and SQLite find methods.
func RetrieveDbDataAll() []utils.JobPost {

	// Upsert job posts to the database
	// if os.Getenv("DB_MONGODB_ENABLE") == "true" {
	// 	// loadDbDataToMongoDb(data, jobId)
	// 	return []utils.JobPost{}
	// }

	if os.Getenv("DB_SQLITE_ENABLE") == "true" {
		return retrieveDbFromSqliteAll()
	}
	return []utils.JobPost{}
}

// RetrieveDbDataSearch is wrapper for MongoDB and SQLite search methods that retrieves job posts based on the search term.
func RetrieveDbDataSearch(searchTerm string) []utils.JobPost {

	// Upsert job posts to the database
	if os.Getenv("DB_MONGODB_ENABLE") == "true" {
		resultJobPosts, err := retrieveDbFromMongoDbSearch(searchTerm)
		if err != nil {
			log.Print(err)
			return []utils.JobPost{}
		}
		return resultJobPosts
	}

	if os.Getenv("DB_SQLITE_ENABLE") == "true" {
		resultJobPosts, err := retrieveDbFromSqliteSearch(searchTerm)
		if err != nil {
			log.Print(err)
			return []utils.JobPost{}
		}
		return resultJobPosts
	}
	return []utils.JobPost{}
}
