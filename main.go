package main

import (
	"fmt"
	"job-funnel/transform"
	"log"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error - Error loading .env file: %s", err)
	}

	jobs, err := transform.Weworkremotely_comCreateJobPostsRss("https://weworkremotely.com/categories/remote-back-end-programming-jobs.rss")
	if err != nil {
		fmt.Println(err)
	}
	for _, job := range jobs {
		fmt.Println(job.JobTitle)
	}

	// Connect to the database
	// load.ConnectDatabaseSQLite()

	// Establish a waitgroup
	var waitgroup sync.WaitGroup

	// Start webserver
	// api.StartServer()

	// Wait for waitgroup to finish
	waitgroup.Wait()
}
