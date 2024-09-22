package tasks

import (
	"fmt"
	"job-funnel/load"
	"job-funnel/transform"
	"log"
)

func InitTasks() {
	Weworkremotely_comRss()
}

func Weworkremotely_comRss() {
	log.Printf("Task: Weworkremotely.com - Executing RSS Feed Job.")
	jobs, err := transform.Weworkremotely_comCreateJobPostsRss("https://weworkremotely.com/categories/remote-back-end-programming-jobs.rss")
	if err != nil {
		errorMessage := fmt.Sprintf("error - Error processing Rss feed: %s\n", err)
		log.Print(errorMessage)
	}
	for _, job := range jobs {
		load.LoadDbData(job)
	}
	log.Println("Task: Weworkremotely.com - Execution Complete!")
}
