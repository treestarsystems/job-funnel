package retrieve

import (
	"errors"
	"fmt"
	"job-funnel/utils"
	"log"
	"os"
	"strings"
)

func retrieveDbFromSqliteAll() []utils.JobPost {
	var retrievedData []utils.JobPost
	utils.DB.Table(*utils.TableName).Find(&retrievedData)
	return retrievedData
}

func retrieveDbFromSqliteSearch(searchTerm string) ([]utils.JobPost, error) {
	var jobPosts []utils.JobPost

	// Need a way to get the correct file path no matter the OS.
	// This will rerun the connection to the database if the file does not exist.
	fileName := fmt.Sprintf("./%v", os.Getenv("DB_SQLITE_FILENAME"))
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Printf("info - SQLite: Database file does not exist, recreating...\n")
		// LoadDbConnectToSqlite()
	}

	// Perform the search query
	err := utils.DB.Table(*utils.TableName).Where(
		// `LOWER(job_id) LIKE ?
		`job_id LIKE ? 
		OR LOWER(job_title) LIKE ? 
		OR LOWER(description) LIKE ? 
		OR LOWER(coding_language) LIKE ? 
		OR LOWER(coding_framework) LIKE ? 
		OR LOWER(database) LIKE ? 
		OR LOWER(company_name) LIKE ?
		OR LOWER(pay) LIKE ?
		OR LOWER(work_location) LIKE ?
		OR LOWER(links) LIKE ?
		OR LOWER(applied_to_job) LIKE ?`,
		"%"+strings.ToLower(searchTerm)+"%",
		"%"+strings.ToLower(searchTerm)+"%",
		"%"+strings.ToLower(searchTerm)+"%",
		"%"+strings.ToLower(searchTerm)+"%",
		"%"+strings.ToLower(searchTerm)+"%",
		"%"+strings.ToLower(searchTerm)+"%",
		"%"+strings.ToLower(searchTerm)+"%",
		"%"+strings.ToLower(searchTerm)+"%",
		"%"+strings.ToLower(searchTerm)+"%",
		"%"+strings.ToLower(searchTerm)+"%",
		"%"+strings.ToLower(searchTerm)+"%",
	).Find(&jobPosts).Error

	if err != nil {
		return nil, err
	}

	return jobPosts, nil
}

// retrieveDbFromSqliteRandom retrieves one random job post from the SQLite database.
func retrieveDbFromSqliteRandom() ([]utils.JobPost, error) {
	var jobPosts []utils.JobPost

	// Perform the query to retrieve one random row
	err := utils.DB.Table(*utils.TableName).Order("RANDOM()").Limit(1).Find(&jobPosts).Error
	if err != nil {
		return nil, err
	}

	return jobPosts, nil
}
