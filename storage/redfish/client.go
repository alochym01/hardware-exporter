package redfish

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var Client *APIClient

// var Client APIClient

// APIClient ...
type APIClient struct {
	User       string
	Pass       string
	HTTPClient *http.Client
	ChasURL    string
	SysURL     string
	// URL        string
	Host string
}

// Get ....
func (c APIClient) Get(url string) ([]byte, error) {
	// Make a http request
	res, err := c.fetch(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Close http.Request connection
	defer res.Body.Close()

	// read the whole body into a []bytes
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c APIClient) fetch(url string) (*http.Response, error) {
	// Create a new request
	fmt.Println("Storage URL -- ", url)
	// req, err := http.NewRequest("GET", c.URL, nil)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// Add header accept application/json
	req.Header.Add("Accept", `application/json`)

	// Set username/password in a http request
	req.SetBasicAuth(c.User, c.Pass)

	// Make a http request with custom Header
	return c.HTTPClient.Do(req)
}

// NewAPIClient return a APIClient struct
func NewAPIClient(c *http.Client) *APIClient {
	return &APIClient{
		User:       "root",
		Pass:       "calvin",
		HTTPClient: c,
		// URL:        "",
		ChasURL: "",
		SysURL:  "",
		Host:    "",
	}
}
