package main

import (
	"job-funnel/api"
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
	if os.Getenv("DB_ENABLE_SQLITE") == "true" {
		load.LoadDbConnectToSqlite()
	}

	// jobs, err := transform.Weworkremotely_comCreateJobPostsRss("https://weworkremotely.com/categories/remote-back-end-programming-jobs.rss")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, job := range jobs {
	// 	load.LoadDbData(job)
	// }

	// Establish a waitgroup
	// var waitgroup sync.WaitGroup

	// Start webserver
	api.StartServer()

	// Wait for waitgroup to finish
	// waitgroup.Wait()
}
