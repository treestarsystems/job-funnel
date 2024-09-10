package transform

import (
	"encoding/xml"
	"job-funnel/extract"
	"strings"
)

type BuilderWeworkremotely_com interface {
	FetchRSSWeworkremotely_com(url string) (JobRssWeworkremotely_com, error)
	ParseJobRssWeworkremotely_com(body string) (JobRssWeworkremotely_com, error)
}

func FetchRSSWeworkremotely_com(url string) (JobRssWeworkremotely_com, error) {
	body, err := extract.FetchRSS(url)
	if err != nil {
		return JobRssWeworkremotely_com{}, err
	}
	parsed, err := ParseJobRssWeworkremotely_com(body)
	if err != nil {
		return JobRssWeworkremotely_com{}, err
	}
	return parsed, nil
}

func ParseJobRssWeworkremotely_com(body string) (JobRssWeworkremotely_com, error) {
	reader := strings.NewReader(body)
	decoder := xml.NewDecoder(reader)

	var job JobRssWeworkremotely_com
	err := decoder.Decode(&job)
	if err != nil {
		return job, err
	}
	return job, nil
}

// func (j *JobRssWeworkremotely_com) Transform() []Job {
// 	var jobs []Job
// 	for _, item := range j.Channel.Item {
// 		job := Job{
// 			Title:       item.Title,
// 			Company:     item.Category,
// 			Location:    item.Region,
// 			Link:        item.Link,
// 			Description: item.Description,
// 			PostedAt:    item.PubDate,
// 			ExpiresAt:   item.ExpiresAt,
// 		}
// 		jobs = append(jobs, job)
// 	}
// 	return jobs
// }
