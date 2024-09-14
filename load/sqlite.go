package load

import (
	"job-funnel/transform"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// var DB *gorm.DB
// var TableName *string
// var SqliteDbName *string

// func LoadDbConnectToSqlite() {
// 	// Get table name
// 	tableName := os.Getenv("DB_TABLE_NAME")
// 	sqliteDbName := os.Getenv("DB_SQLITE_FILENAME")
// 	TableName = &tableName
// 	SqliteDbName = &sqliteDbName
// 	// Create DB Connection
// 	db, err := gorm.Open(sqlite.Open(*SqliteDbName), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("error - SQLite: Unable to establish database connection: %s", err)
// 	}
// 	// Migrate the schema/Create the table
// 	err = db.Table(*TableName).AutoMigrate(&LoadDbInsertGorm{})
// 	if err != nil {
// 		log.Fatalf("error - SQLite: Unable to migrate the schema: %s", err)
// 	}
// 	DB = db
// }

// func loadDbDataToSqlite(data transform.JobPost) {
// 	// Precheck if DB is connected
// 	if DB == nil {
// 		LoadDbConnectToSqlite()
// 	}

// 	// Save = Upsert
// 	result := DB.Table(*TableName).Where(transform.JobPost{JobTitle: data.JobTitle}).Assign(transform.JobPost{
// 		Description:     data.Description,
// 		CodingLanguage:  data.CodingLanguage,
// 		CodingFramework: data.CodingFramework,
// 		Database:        data.Database,
// 		CompanyName:     data.CompanyName,
// 		Pay:             data.Pay,
// 		Location:        data.Location,
// 		Links:           data.Links,
// 	}).FirstOrCreate(&transform.JobPost{
// 		Description:     data.Description,
// 		CodingLanguage:  data.CodingLanguage,
// 		CodingFramework: data.CodingFramework,
// 		Database:        data.Database,
// 		CompanyName:     data.CompanyName,
// 		Pay:             data.Pay,
// 		Location:        data.Location,
// 		Links:           data.Links,
// 	})

// 	if result.Error != nil {
// 		fmt.Println("error - SQLite: Unable to upsert data: %s", result.Error)
// 		fmt.Print(DB)
// 	}

// }

var DB *gorm.DB
var TableName *string
var SqliteDbName *string

// LoadDbConnectToSqlite initializes the SQLite database connection and sets the global DB variable.
func LoadDbConnectToSqlite() {
	// Get table name
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

// CheckDbConnection checks if the database connection is alive and reconnects if necessary.
func CheckDbConnection() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("error - SQLite: Unable to establish database connection: %s", err)
	}

	// Ping the database to check if the connection is alive
	err = sqlDB.Ping()
	if err != nil {
		log.Println("info - SQLite: Database connection lost, reconnecting...")
		LoadDbConnectToSqlite()
	}
}

func loadDbDataToSqlite(data transform.JobPost) {
	// Precheck if DB is connected
	if DB == nil {
		LoadDbConnectToSqlite()
	} else {
		CheckDbConnection()
	}

	// Save = Upsert
	result := DB.Table(*TableName).Where(transform.JobPost{JobTitle: data.JobTitle}).Assign(transform.JobPost{
		Description:     data.Description,
		CodingLanguage:  data.CodingLanguage,
		CodingFramework: data.CodingFramework,
		Database:        data.Database,
		CompanyName:     data.CompanyName,
		Pay:             data.Pay,
		Location:        data.Location,
		Links:           data.Links,
	}).FirstOrCreate(&transform.JobPost{})

	if result.Error != nil {
		log.Printf("error - SQLite: Unable to upsert job post: %s", result.Error)
	}
}
