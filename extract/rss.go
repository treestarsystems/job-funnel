package extract

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
	cf "github.com/iarsham/cf-forbidden"
	// "xml"
)

// type ItemRSS struct {
// 	XMLName string `xml:"item"`
// 	Title   string `xml:"title"`
// 	Link    string `xml:"link"`
// }

type JobRssWeworkremotely_com struct {
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

// FetchRSS fetches the RSS feed from the given URL and returns the response body.
func FetchRSS(url string) (string, error) {
	u := govalidator.IsURL(url)
	if !u {
		errorMessage := fmt.Errorf("error - Invalid URL: %s", url)
		return "", errorMessage
	}
	client, err := cf.New()
	if err != nil {
		errorMessage := fmt.Errorf("error - Client creation %v", err)
		return "", errorMessage
	}
	response, err := client.Get(url, cf.M{})
	if err != nil {
		errorMessage := fmt.Errorf("error - Client GET Request %v", err)
		return "", errorMessage
	}
	body := response.Body
	// body := `
	// 	<person>
	// 		<name>John Doe</name>
	// 		<age>30</age>
	// 	</person>
	// `

	reader := strings.NewReader(body)
	decoder := xml.NewDecoder(reader)

	var job JobRssWeworkremotely_com
	err2 := decoder.Decode(&job)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
		return "", err2
	}

	fmt.Printf("%s", job.Channel.Item[0].Title)
	// fmt.Printf("%s", job.Name)

	return body, nil
	// // Convert the body string to an io.Reader
	// reader := strings.NewReader(body)

	// // Decode the XML
	// decoder := xml.NewDecoder(reader)
	// decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
	// 	return input, nil
	// }

	// var result interface{}
	// if err := decoder.Decode(&result); err != nil {
	// 	return "", fmt.Errorf("error - XML Decoding %v", err)
	// }

	// fmt.Print(body)

	// // Marshal the decoded XML back to a string (optional)
	// decodedXML, err := xml.Marshal(result)
	// if err != nil {
	// 	return "", fmt.Errorf("error - XML Marshalling %v", err)
	// }

	// return string(decodedXML), nil
}
