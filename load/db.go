package load

import (
	"job-funnel/transform"
	"os"
)

func LoadDbData(data transform.JobPost) {

	// Save to SQLite
	if os.Getenv("DB_ENABLE_SQLITE") == "true" {
		loadDbDataToSqlite(data)
	}
	if os.Getenv("DB_ENABLE_MONGO") == "true" {
		// loadDbDataToMongoDB(data)
	}
}
