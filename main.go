package main

import (
	"job-funnel/extract"
)

func main() {
	extract.FetchRSS("https://weworkremotely.com/categories/remote-back-end-programming-jobs.rss")
	// rssData, err := extract.FetchRSS("https://weworkremotely.com/categories/remote-back-end-programming-jobs.rss")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(rssData)
}
