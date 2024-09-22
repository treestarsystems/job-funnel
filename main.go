package main

import (
	"job-funnel/api"
	"job-funnel/communication"
	"job-funnel/cron"
	"job-funnel/tasks"
	"job-funnel/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Initialize and check for command line flags
	utils.InitCommandLineFlags()

	// Load environment variables
	err := godotenv.Load(utils.EnvFilePath)
	if err != nil {
		log.Fatalf("error - Error loading .env file: %s", err)
	}

	// Connect to the databases
	if os.Getenv("DB_SQLITE_ENABLE") == "true" {
		utils.LoadDbConnectToSqlite()
	}

	if os.Getenv("DB_MONGODB_ENABLE") == "true" {
		utils.LoadDbConnectToMongoDb()
	}

	if os.Getenv("COMMUNICATION_DISCORD_ENABLE") == "true" {
		// Start as a non-blocking goroutine
		go communication.InitDiscordBot()
	}

	// Initial run of tasks on startup as a non-blocking goroutine
	go tasks.InitTasks()

	// Initialize cron jobs
	cron.InitCron()

	// Start webserver
	api.StartServer()
}
