package tasks

import (
	"job-funnel/load"
	"job-funnel/transform"
	"log"
)

func InitTasks() {
	Weworkremotely_comRss()
}

func Weworkremotely_comRss() {
	log.Println("Task: Weworkremotely.com - Executing RSS Feed Job.")
	jobs, err := transform.Weworkremotely_comCreateJobPostsRss("https://weworkremotely.com/categories/remote-back-end-programming-jobs.rss")
	if err != nil {
		log.Println(err)
	}
	for _, job := range jobs {
		load.LoadDbData(job)
	}
	log.Println("Task: Weworkremotely.com - Execution Complete!")
}
