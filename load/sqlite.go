package load

import (
	"errors"
	"fmt"
	"job-funnel/transform"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var TableName *string
var SqliteDbName *string

func LoadDbConnectToSqlite() {
	tableName := os.Getenv("DB_TABLE_NAME")
	sqliteDbName := os.Getenv("DB_SQLITE_FILENAME")
	TableName = &tableName
	SqliteDbName = &sqliteDbName

	db, err := gorm.Open(sqlite.Open(*SqliteDbName), &gorm.Config{})
	if err != nil {
		log.Printf("error - SQLite: Unable to establish database connection: %s", err)
	}
	// Migrate the schema/Create the table
	err = db.Table(*TableName).AutoMigrate(&LoadDbInsertGorm{})
	if err != nil {
		log.Printf("error - SQLite: Unable to migrate the schema: %s", err)
	}
	DB = db
}

func loadDbDataToSqlite(data transform.JobPost) {
	// Need a way to get the correct file path no matter the OS.
	// This will rerun the connection to the database if the file does not exist.
	fileName := fmt.Sprintf("./%v", os.Getenv("DB_SQLITE_FILENAME"))
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Println("info - SQLite: Database file does not exist, recreating...")
		LoadDbConnectToSqlite()
	}

	// Save = Upsert
	DB.Table(*TableName).Where(transform.JobPost{JobTitle: data.JobTitle}).Assign(transform.JobPost{
		Description:     data.Description,
		CodingLanguage:  data.CodingLanguage,
		CodingFramework: data.CodingFramework,
		Database:        data.Database,
		CompanyName:     data.CompanyName,
		Pay:             data.Pay,
		WorkLocation:    data.WorkLocation,
		Links:           data.Links,
	}).FirstOrCreate(&transform.JobPost{
		Description:     data.Description,
		CodingLanguage:  data.CodingLanguage,
		CodingFramework: data.CodingFramework,
		Database:        data.Database,
		CompanyName:     data.CompanyName,
		Pay:             data.Pay,
		WorkLocation:    data.WorkLocation,
		Links:           data.Links,
	})
}
