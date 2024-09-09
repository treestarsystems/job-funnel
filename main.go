package main

import (
	"fmt"
	"job-funnel/extract"
)

func main() {
	rssData, err := extract.FetchRSS("https://weworkremotely.com/categories/remote-back-end-programming-jobs.rss")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rssData)
}
