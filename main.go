package main

import (
	"job-funnel/api"
	"job-funnel/cron"
	"job-funnel/load"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error - Error loading .env file: %s", err)
	}

	// Connect to the database
	if os.Getenv("DB_SQLITE_ENABLE") == "true" {
		load.LoadDbConnectToSqlite()
	}

	if os.Getenv("DB_MONGODB_ENABLE") == "true" {
		load.LoadDbConnectToMongoDb()
	}

	// Initialize cron jobs
	cron.InitCron()

	// Start webserver
	api.StartServer()

	// Wait for waitgroup to finish
	// waitgroup.Wait()
}
