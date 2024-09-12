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

func ConnectDatabase() {
	// Get table name
	tableName := os.Getenv("TABLE_NAME")
	sqliteDbName := os.Getenv("SQLITE_DB_NAME")
	TableName = &tableName
	SqliteDbName = &sqliteDbName
	// Create DB Connection
	db, err := gorm.Open(sqlite.Open(*SqliteDbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("error - Unable to establish database connection: %s", err)
	}
	// Migrate the schema/Create the table
	err = db.Table(*TableName).AutoMigrate(&transform.JobPost{})
	if err != nil {
		log.Fatalf("error - Unable to migrate the schema: %s", err)
	}
	DB = db
}

func LoadDataToSqlite(data transform.JobPost) {
	// Save = Upsert
	DB.Table(*TableName).Where(DbLoadJobPost{JobTitle: data.JobTitle}).Assign(DbLoadJobPost{
		Description:    data.Description,
		CodingLanguage: data.CodingLanguage,
		Database:       data.Database,
		CompanyName:    data.CompanyName,
		Pay:            data.Pay,
		Location:       data.Location,
		Links:          data.Links,
	}).FirstOrCreate(&DbLoadJobPost{
		Description:    data.Description,
		CodingLanguage: data.CodingLanguage,
		Database:       data.Database,
		CompanyName:    data.CompanyName,
		Pay:            data.Pay,
		Location:       data.Location,
		Links:          data.Links,
	})
}
