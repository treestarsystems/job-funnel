package main

import (
	"fmt"
	"job-funnel/transform"
)

func main() {
	data, err := transform.FetchRSSWeworkremotely_com("https://weworkremotely.com/remote-jobs.rss")
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range data.Channel.Item {
		fmt.Printf("%s\n", item.Title)
	}
}
