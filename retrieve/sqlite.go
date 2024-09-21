package retrieve

import (
	"errors"
	"fmt"
	"job-funnel/utils"
	"log"
	"os"
	"strings"
)

func RetrieveDbFromSqliteAll() []utils.JobPost {
	var retrievedData []utils.JobPost
	utils.DB.Table(*utils.TableName).Find(&retrievedData)
	return retrievedData
}

func RetrieveDbFromSqliteSearch(searchTerm string) ([]utils.JobPost, error) {
	var jobPosts []utils.JobPost

	// Need a way to get the correct file path no matter the OS.
	// This will rerun the connection to the database if the file does not exist.
	fileName := fmt.Sprintf("./%v", os.Getenv("DB_SQLITE_FILENAME"))
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Println("info - SQLite: Database file does not exist, recreating...")
		// LoadDbConnectToSqlite()
	}

	// Perform the search query
	err := utils.DB.Table(*utils.TableName).Where(
		`LOWER(job_id) LIKE ? 
		OR LOWER(job_title) LIKE ? 
		OR LOWER(description) LIKE ? 
		OR LOWER(coding_language) LIKE ? 
		OR LOWER(coding_framework) LIKE ? 
		OR LOWER(database) LIKE ? 
		OR LOWER(company_name) LIKE ?
		OR LOWER(pay) LIKE ?
		OR LOWER(work_location) LIKE ?
		OR LOWER(links) LIKE ?
		OR LOWER(appled_to_job) LIKE ?`,
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
