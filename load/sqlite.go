package load

import (
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
	// Get table name
	tableName := os.Getenv("DB_TABLE_NAME")
	sqliteDbName := os.Getenv("DB_SQLITE_FILENAME")
	TableName = &tableName
	SqliteDbName = &sqliteDbName
	// Create DB Connection
	db, err := gorm.Open(sqlite.Open(*SqliteDbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("error - Unable to establish database connection: %s", err)
	}
	// Migrate the schema/Create the table
	err = db.Table(*TableName).AutoMigrate(&LoadDbGormInsertJobPost{})
	if err != nil {
		log.Fatalf("error - Unable to migrate the schema: %s", err)
	}
	DB = db
}

func loadDbDataToSqlite(data transform.JobPost) {
	// Save = Upsert
	DB.Table(*TableName).Where(LoadDbGormInsertJobPost{JobTitle: data.JobTitle}).Assign(LoadDbGormInsertJobPost{
		Description:    data.Description,
		CodingLanguage: data.CodingLanguage,
		Database:       data.Database,
		CompanyName:    data.CompanyName,
		Pay:            data.Pay,
		Location:       data.Location,
		Links:          data.Links,
	}).FirstOrCreate(&LoadDbGormInsertJobPost{
		Description:    data.Description,
		CodingLanguage: data.CodingLanguage,
		Database:       data.Database,
		CompanyName:    data.CompanyName,
		Pay:            data.Pay,
		Location:       data.Location,
		Links:          data.Links,
	})
}
