package extract

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	cf "github.com/iarsham/cf-forbidden"
)

// FetchRss fetches the Rss feed from the given URL and returns the response body.
func FetchRss(url string) (string, error) {
	u := govalidator.IsURL(url)
	if !u {
		errorMessage := fmt.Errorf("error - Rss Invalid URL: %s", url)
		return "", errorMessage
	}
	client, err := cf.New()
	if err != nil {
		errorMessage := fmt.Errorf("error - Rss Client creation: %v - %s", err, url)
		return "", errorMessage
	}
	response, err := client.Get(url, cf.M{})
	if err != nil {
		errorMessage := fmt.Errorf("error - Rss Client GET Request: %v - %s", err, url)
		return "", errorMessage
	}
	body := response.Body
	return body, nil
}
