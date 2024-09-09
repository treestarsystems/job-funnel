package extract

import (
	"encoding/xml"
	"io"
	"net/http"
)

type ItemRSS struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
}

type ChannelRSS struct {
	XMLName xml.Name  `xml:"channel"`
	Items   []ItemRSS `xml:"item"`
}

type DocumentRSS struct {
	XMLName xml.Name   `xml:"rss"`
	Channel ChannelRSS `xml:"channel"`
}

func Fetch(url string) (DocumentRSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return DocumentRSS{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DocumentRSS{}, err
	}
	var doc DocumentRSS
	err = xml.Unmarshal(body, &doc)
	if err != nil {
		return DocumentRSS{}, err
	}
	return doc, nil
}
