package load

import (
	"errors"
	"fmt"
	"job-funnel/utils"
	"log"
	"os"
)

func loadDbDataToSqlite(data utils.JobPost, jobId string) {
	// TODO: Need a way to get the correct file path no matter the OS.
	// This will rerun the connection to the database if the file does not exist.
	fileName := fmt.Sprintf("./%v", os.Getenv("DB_SQLITE_FILENAME"))
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Printf("info - SQLite: Database file does not exist, recreating...\n")
		utils.LoadDbConnectToSqlite()
	}

	// Save = Upsert
	utils.DB.Table(*utils.TableName).Where(utils.JobPost{JobTitle: data.JobTitle}).Assign(utils.JobPost{
		JobSource:       data.JobSource,
		Description:     data.Description,
		CodingLanguage:  data.CodingLanguage,
		CodingFramework: data.CodingFramework,
		Database:        data.Database,
		CompanyName:     data.CompanyName,
		Pay:             data.Pay,
		WorkLocation:    data.WorkLocation,
		Links:           data.Links,
	}).FirstOrCreate(&utils.JobPost{
		JobId:           jobId,
		JobSource:       data.JobSource,
		Description:     data.Description,
		CodingLanguage:  data.CodingLanguage,
		CodingFramework: data.CodingFramework,
		Database:        data.Database,
		CompanyName:     data.CompanyName,
		Pay:             data.Pay,
		WorkLocation:    data.WorkLocation,
		Links:           data.Links,
		AppliedToJob:    []string{},
	})
}
