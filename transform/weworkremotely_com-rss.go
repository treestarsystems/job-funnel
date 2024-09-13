dddpackage transform

import (
	"encoding/xml"
	"job-funnel/extract"
	"job-funnel/utils"
	"strings"
)

func Weworkremotely_comParseRss(rssXMLBody string) (Weworkremotely_comRss, error) {
	reader := strings.NewReader(rssXMLBody)
	decoder := xml.NewDecoder(reader)
	var job Weworkremotely_comRss
	err := decoder.Decode(&job)
	if err != nil {
		return job, err
	}
	return job, nil
}

// ProcessRss processes the given Rss feed URL and returns a Weworkremotely_comRss struct.
func Weworkremotely_comProcessRss(url string) (Weworkremotely_comRss, error) {
	body, err := extract.FetchRss(url)
	if err != nil {
		return Weworkremotely_comRss{}, err
	}
	parsedBody, err := Weworkremotely_comParseRss(body)
	if err != nil {
		return Weworkremotely_comRss{}, err
	}
	return parsedBody, nil
}

// ProcessJobPosts processes job posts from the given Rss feed URL and returns a slice of JobPost.
func Weworkremotely_comCreateJobPostsRss(feedURL string) ([]JobPost, error) {
	data, err := Weworkremotely_comProcessRss(feedURL)
	if err != nil {
		return nil, err
	}

	var jobs []JobPost
	for _, item := range data.Channel.Item {
		description := utils.RemoveHTMLTags(item.Description)
		codingLanguages := utils.ParseProgrammingLanguages(description)
		database := utils.ParseDatabaseTypes(description)
		pay := utils.ParseSalaries(description)
		links := utils.ParseNonImageLinks(description)
		location := utils.ParseCityOrState(description)
		job := JobPost{
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
	}
	return jobs, nil
}
