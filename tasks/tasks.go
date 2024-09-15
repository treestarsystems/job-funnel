package tasks

import (
	"fmt"
	"job-funnel/load"
	"job-funnel/transform"
)

func InitTasks() {
	Weworkremotely_comRss()
}

func Weworkremotely_comRss() {
	fmt.Println("Cron: Weworkremotely.com - Executing RSS Feed Job.")
	jobs, err := transform.Weworkremotely_comCreateJobPostsRss("https://weworkremotely.com/categories/remote-back-end-programming-jobs.rss")
	if err != nil {
		fmt.Println(err)
	}
	for _, job := range jobs {
		load.LoadDbData(job)
	}
	fmt.Println("Cron: Weworkremotely.com - Execution Complete!")
}
