package retrieve

import (
	"job-funnel/utils"
	"os"
)

func RetrieveDbDataAll() []utils.JobPost {

	// Upsert job posts to the database
	// if os.Getenv("DB_MONGODB_ENABLE") == "true" {
	// 	// loadDbDataToMongoDb(data, jobId)
	// 	return []utils.JobPost{}
	// }

	if os.Getenv("DB_SQLITE_ENABLE") == "true" {
		return RetrieveDbFromSqliteAll()
	}
	return []utils.JobPost{}
}
