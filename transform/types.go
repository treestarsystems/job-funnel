package transform

import (
	"encoding/xml"
	"time"

	"gorm.io/datatypes"
)

// SharedStructJobs contains common job information from each job listing.
type JobPost struct {
	JobTitle       string                      `bson:"job_title" json:"jobTitle" binding:"required"`
	Description    string                      `bson:"description" json:"description" binding:"required"`
	CodingLanguage datatypes.JSONSlice[string] `bson:"coding_language" json:"codingLanguage" binding:"required"`
	Database       datatypes.JSONSlice[string] `bson:"database" json:"database" binding:"required"`
	CompanyName    string                      `bson:"company_name" json:"companyName" binding:"required"`
	Pay            datatypes.JSONSlice[string] `bson:"pay" json:"pay" binding:"required"`
	Location       datatypes.JSONSlice[string] `bson:"location" json:"location" binding:"required"`
	Links          datatypes.JSONSlice[string] `bson:"link" json:"link" binding:"required"`
	CreatedAt      time.Time                   `bson:"created_at" json:"createdAt" binding:"required"`
	UpdatedAt      time.Time                   `bson:"updated_at" json:"updatedAt" binding:"required"`
	// CreatedAt      time.Time                   `bson:"createdAt" json:"createdAt" binding:"required"`
	// UpdatedAt      time.Time                   `bson:"updatedAt" json:"updatedAt" binding:"required"`
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
