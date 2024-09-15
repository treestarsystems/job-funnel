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
		return "", fmt.Errorf("error - Rss Invalid URL: %s", url)
	}
	client, err := cf.New()
	if err != nil {
		return "", fmt.Errorf("error - Rss Client creation: %v - %s", err, url)
	}
	response, err := client.Get(url, cf.M{})
	if err != nil {
		return "", fmt.Errorf("error - Rss Client GET Request: %v - %s", err, url)
	}
	body := response.Body
	return body, nil
}
