package extract

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"job-funnel/utils"
	"net/http"
	"time"
)

// MakeAPICall makes an API call to the specified URL with the given method and payload.
func FetchAPIResponse(url, method string, payload interface{}) (*utils.APIResponse, error) {
	// Convert payload to JSON
	var jsonPayload []byte
	var err error
	if payload != nil {
		jsonPayload, err = json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("error - API marshaling payload: %v - %s", err, url)
		}
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error - API creating request: %v - %s", err, url)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 10 * time.Second}

	// Make the API call
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error - API making request: %v - %s", err, url)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error - API reading response body: %v - %s", err, url)
	}

	// Check for non-200 status codes
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("error - API non-200 status code: %d, body: %s", resp.StatusCode, body)
	}

	// Parse the response body into APIResponse
	var apiResponse utils.APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("error - API unmarshaling response: %v - %s", err, url)
	}

	return &apiResponse, nil
}
