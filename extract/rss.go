package extract

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	cf "github.com/iarsham/cf-forbidden"
)

// FetchRSS fetches the RSS feed from the given URL and returns the response body.
func FetchRSS(url string) (string, error) {
	u := govalidator.IsURL(url)
	if !u {
		errorMessage := fmt.Errorf("error - RSS Invalid URL: %s", url)
		return "", errorMessage
	}
	client, err := cf.New()
	if err != nil {
		errorMessage := fmt.Errorf("error - RSS Client creation: %v", err)
		return "", errorMessage
	}
	response, err := client.Get(url, cf.M{})
	if err != nil {
		errorMessage := fmt.Errorf("error - RSS Client GET Request: %v", err)
		return "", errorMessage
	}
	body := response.Body
	return body, nil
}
