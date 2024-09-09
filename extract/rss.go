package extract

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	cf "github.com/iarsham/cf-forbidden"
)

// type ItemRSS struct {
// 	XMLName     xml.Name `xml:"item"`
// 	Title       string   `xml:"title"`
// 	Link        string   `xml:"link"`
// 	Description string   `xml:"description"`
// }

// type ChannelRSS struct {
// 	XMLName xml.Name  `xml:"channel"`
// 	Items   []ItemRSS `xml:"item"`
// }

// type DocumentRSS struct {
// 	XMLName xml.Name   `xml:"rss"`
// 	Channel ChannelRSS `xml:"channel"`
// }

// fetchRSS fetches the RSS feed from the given URL and returns the response body.
func FetchRSS(url string) (string, error) {
	u := govalidator.IsURL(url)
	if !u {
		errorMessage := fmt.Errorf("error - Invalid URL: %s", url)
		return "", errorMessage
	}
	client, err := cf.New()
	if err != nil {
		fmt.Print(err)
	}
	response, err := client.Get(url, cf.M{"Authorization": ""})
	if err != nil {
		errorMessage := fmt.Errorf("error - %v", err)
		return "", errorMessage
	}
	body := response.Body
	headers := response.Headers
	fmt.Println(body, headers)

	return body, nil
}
