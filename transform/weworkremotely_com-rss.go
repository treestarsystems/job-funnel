package transform

import (
	"encoding/xml"
	"job-funnel/extract"
	"strings"
)

type Weworkremotely_comRSSBuilder interface {
	FetchRSSWeworkremotely_com(url string) (Weworkremotely_comRss, error)
	ParseWeworkremotely_comRss(body string) (Weworkremotely_comRss, error)
}

func ProcessRSSWeworkremotely_com(url string) (Weworkremotely_comRss, error) {
	body, err := extract.FetchRSS(url)
	if err != nil {
		return Weworkremotely_comRss{}, err
	}
	parsedBody, err := ParseWeworkremotely_comRss(body)
	if err != nil {
		return Weworkremotely_comRss{}, err
	}
	return parsedBody, nil
}

func ParseWeworkremotely_comRss(rssXMLBody string) (Weworkremotely_comRss, error) {
	reader := strings.NewReader(rssXMLBody)
	decoder := xml.NewDecoder(reader)
	var job Weworkremotely_comRss
	err := decoder.Decode(&job)
	if err != nil {
		return job, err
	}
	return job, nil
}

// func (j *Weworkremotely_comRss) Transform() []Job {
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
