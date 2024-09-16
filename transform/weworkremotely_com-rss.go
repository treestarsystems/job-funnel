package transform

import (
	"encoding/xml"
	"job-funnel/extract"
	"job-funnel/utils"
	"strings"
)

func Weworkremotely_comParseRss(rssXMLBody string) (utils.Weworkremotely_comRss, error) {
	reader := strings.NewReader(rssXMLBody)
	decoder := xml.NewDecoder(reader)
	var job utils.Weworkremotely_comRss
	err := decoder.Decode(&job)
	if err != nil {
		return job, err
	}
	return job, nil
}

// ProcessRss processes the given Rss feed URL and returns a utils.Weworkremotely_comRss struct.
func Weworkremotely_comProcessRss(url string) (utils.Weworkremotely_comRss, error) {
	body, err := extract.FetchRss(url)
	if err != nil {
		return utils.Weworkremotely_comRss{}, err
	}
	parsedBody, err := Weworkremotely_comParseRss(body)
	if err != nil {
		return utils.Weworkremotely_comRss{}, err
	}
	return parsedBody, nil
}

// ProcessJobPosts processes job posts from the given Rss feed URL and returns a slice of JobPost.
func Weworkremotely_comCreateJobPostsRss(feedURL string) ([]utils.JobPost, error) {
	data, err := Weworkremotely_comProcessRss(feedURL)
	if err != nil {
		return nil, err
	}

	var jobs []utils.JobPost
	for _, item := range data.Channel.Item {
		description := utils.RemoveHTMLTags(item.Description)
		codingLanguage := utils.ParseCommonProgrammingLanguages(description)
		codingFramework := utils.ParseCommonFrameworks(description)
		database := utils.ParseDatabaseTypes(description)
		pay := utils.ParseSalaries(description)
		links := utils.ParseNonImageLinks(description)
		workLocation := utils.ParseJobWorkLocation(description)
		job := utils.JobPost{
			JobTitle:        item.Title,
			Description:     description,
			CodingLanguage:  codingLanguage,
			CodingFramework: codingFramework,
			Database:        database,
			CompanyName:     "",
			Pay:             pay,
			WorkLocation:    workLocation,
			Links:           links,
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}
