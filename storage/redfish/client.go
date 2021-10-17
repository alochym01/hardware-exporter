package redfish

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var Client *APIClient

// var Client APIClient

// APIClient ...
type APIClient struct {
	User       string
	Pass       string
	HTTPClient *http.Client
	Server     string
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
	// fmt.Println("Storage URL -- ", url)
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
func NewAPIClient(user, pass string) *APIClient {
	// Create a custom Transport
	// The default value of Transport's MaxIdleConnsPerHost.
	// const DefaultMaxIdleConnsPerHost = 2
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.IdleConnTimeout = 60 * time.Second
	transport.MaxIdleConns = 100
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	// Disable SSL check
	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	// Can set User/Pass from CLI
	return &APIClient{
		User: user,
		Pass: pass,
		HTTPClient: &http.Client{
			Transport: transport,
			Timeout:   time.Duration(10) * time.Second,
		},
		Server: "",
	}
}
