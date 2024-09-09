package extract

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	cf "github.com/iarsham/cf-forbidden"
)

type ItemRSS struct {
	XMLName string `xml:"item"`
	Title   string `xml:"title"`
	Link    string `xml:"link"`
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
		errorMessage := fmt.Errorf("error - %v", err)
		return "", errorMessage
	}
	response, err := client.Get(url, cf.M{})
	if err != nil {
		errorMessage := fmt.Errorf("error - %v", err)
		return "", errorMessage
	}
	body := response.Body
	fmt.Println(body)

	return body, nil
}
