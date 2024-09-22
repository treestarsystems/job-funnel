package utils

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLite variables
var DB *gorm.DB
var TableName *string
var SQLiteDbName *string

func LoadDbConnectToSqlite() {
	tableName := os.Getenv("DB_TABLE_NAME")
	sqliteDbName := os.Getenv("DB_SQLITE_FILENAME")
	TableName = &tableName
	SQLiteDbName = &sqliteDbName

	db, err := gorm.Open(sqlite.Open(*SQLiteDbName), &gorm.Config{})
	if err != nil {
		log.Printf("error - SQLite: Unable to establish database connection: %s\n", err)
	}
	// Migrate the schema/Create the table
	err = db.Table(*TableName).AutoMigrate(&LoadDbInsertGorm{})
	if err != nil {
		log.Printf("error - SQLite: Unable to migrate the schema: %s\n", err)
	}
	DB = db
}
