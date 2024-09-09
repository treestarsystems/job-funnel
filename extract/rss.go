package extract

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	// "xml"

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
		errorMessage := fmt.Errorf("error - Client creation %v", err)
		return "", errorMessage
	}
	response, err := client.Get(url, cf.M{})
	if err != nil {
		errorMessage := fmt.Errorf("error - Client GET Request %v", err)
		return "", errorMessage
	}
	body := response.Body
	// Convert the body string to an io.Reader
	reader := strings.NewReader(body)

	// Decode the XML
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return input, nil
	}

	var result interface{}
	if err := decoder.Decode(&result); err != nil {
		return "", fmt.Errorf("error - XML Decoding %v", err)
	}

	// Marshal the decoded XML back to a string (optional)
	decodedXML, err := xml.Marshal(result)
	if err != nil {
		return "", fmt.Errorf("error - XML Marshalling %v", err)
	}

	return string(decodedXML), nil
}
