package load

import (
	"job-funnel/utils"
	"os"
)

func LoadDbData(data utils.JobPost) {
	// Upsert job posts to the database
	if os.Getenv("DB_SQLITE_ENABLE") == "true" {
		loadDbDataToSqlite(data)
	}

	if os.Getenv("DB_MONGODB_ENABLE") == "true" {
		loadDbDataToMongoDb(data)
	}
}
