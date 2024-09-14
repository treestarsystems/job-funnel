package main

import (
	"fmt"
	"job-funnel/api"
	"job-funnel/cron"
	"job-funnel/load"
	"job-funnel/transform"
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

	jobs, err := transform.Weworkremotely_comCreateJobPostsRss("https://weworkremotely.com/categories/remote-back-end-programming-jobs.rss")
	if err != nil {
		fmt.Println(err)
	}
	for _, job := range jobs {
		load.LoadDbData(job)
	}

	// Start webserver
	api.StartServer()

	// Initialize cron jobs
	cron.InitCron()

	// Wait for waitgroup to finish
	// waitgroup.Wait()
}
