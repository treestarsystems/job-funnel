package main

import (
	"fmt"
	"job-funnel/transform"
	"job-funnel/utils"
)

type Job struct {
	Title       string
	Region      string
	Category    string
	Type        string
	Description string
	PubDate     string
	ExpiresAt   string
	Guid        string
	Link        string
}

func main() {
	data, err := transform.ProcessRSSWeworkremotely_com("https://weworkremotely.com/categories/remote-back-end-programming-jobs.rss")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Print(data.Channel.Item[0].Category)
	// for _, item := range data.Channel.Item {
	// 	fmt.Printf("%s\n\n\n", item.Category)
	// }
	var jobs []transform.SharedStructJobs
	// for _, item := range data.Channel.Item {
	// 	job := Job{
	// 		Title:       item.Title,
	// 		Region:      item.Region,
	// 		Category:    item.Category,
	// 		Type:        item.Type,
	// 		Description: item.Description,
	// 		PubDate:     item.PubDate,
	// 		ExpiresAt:   item.ExpiresAt,
	// 		Guid:        item.Guid,
	// 		Link:        item.Link,
	// 	}
	// 	jobs = append(jobs, job)
	// }
	for _, item := range data.Channel.Item {

		description := utils.RemoveHTMLTags(item.Description)
		codingLanguages := utils.ExtractProgrammingLanguages(description)
		database := utils.ExtractDatabaseTypes(description)
		pay := utils.ExtractSalaries(description)
		links := utils.ExtractNonImageLinks(description)
		location := utils.ExtractCityOrState(description)
		job := transform.SharedStructJobs{
			JobTitle:       item.Title,
			Description:    description,
			CodingLanguage: codingLanguages,
			Database:       database,
			CompanyName:    "",
			Pay:            pay,
			Location:       location,
			Links:          links,
		}
		jobs = append(jobs, job)
		// fmt.Println(job.Location)
	}
	fmt.Println(jobs[0])
}
