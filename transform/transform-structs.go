package transform

import "encoding/xml"

// SharedStructJobs contains common job information from each job listing.
type JobPost struct {
	JobTitle       string   `json:"title,omitempty"`
	Description    string   `json:"description,omitempty"`
	CodingLanguage []string `json:"codinglanguage,omitempty"`
	Database       []string `json:"database,omitempty"`
	CompanyName    string   `json:"companyname,omitempty"`
	Pay            []string `json:"pay,omitempty"`
	Location       []string `json:"location,omitempty"`
	Links          []string `json:"link,omitempty"`
}

type Weworkremotely_comRss struct {
	XMLName xml.Name `xml:"rss" json:"rss,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	Version string   `xml:"version,attr" json:"version,omitempty"`
	Dc      string   `xml:"dc,attr" json:"dc,omitempty"`
	Media   string   `xml:"media,attr" json:"media,omitempty"`
	Channel struct {
		Text        string `xml:",chardata" json:"text,omitempty"`
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		Ttl         string `xml:"ttl"`
		Item        []struct {
			Text    string `xml:",chardata" json:"text,omitempty"`
			Content struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				URL  string `xml:"url,attr" json:"url,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"content" json:"content,omitempty"`
			Title       string `xml:"title"`
			Region      string `xml:"region"`
			Category    string `xml:"category"`
			Type        string `xml:"type"`
			Description string `xml:"description"`
			PubDate     string `xml:"pubDate"`
			ExpiresAt   string `xml:"expires_at"`
			Guid        string `xml:"guid"`
			Link        string `xml:"link"`
		} `xml:"item" json:"item,omitempty"`
	} `xml:"channel" json:"channel,omitempty"`
}
