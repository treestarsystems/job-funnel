package extract

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"golang.org/x/net/html"
)

func FetchHTML(url string) (string, error) {
	// Validate URL
	u := govalidator.IsURL(url)
	if !u {
		return "", fmt.Errorf("error - HTML Invalid URL: %s", url)
	}
	// Retrieve HTML data
	resp, err := http.Get(string(url))
	if err != nil {
		return "", fmt.Errorf("error - HTML retrieve data: %v - %s", err, url)
	}
	// Validate HTML response
	if resp.StatusCode > 299 {
		return "", fmt.Errorf("error - HTML non-200 status code: %v - %s", resp.StatusCode, url)
	}
	// Extract HTML body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error - HTML extracting html body: %v - %s", err, url)
	}
	return string(body), err
}

// This needs to be test. I am not sure if it works the way I will need it to.
// It may be better to use github.com/anaskhan96/soup like I didn in DishDashGo.
func ParseHTMLBody(htmlBody string, tagName string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("error - HTML parsing body: %v", err)
	}

	var elements []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tagName {
			var buf strings.Builder
			html.Render(&buf, n)
			elements = append(elements, buf.String())
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return elements, nil
}
