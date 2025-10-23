package openf1

import (
	"net/http"
)

const (
	apiURL     = "https://api.openf1.org/"
	apiVersion = "v1"
)

type Client struct {
	apiURL     string
	apiVersion string
	path       string
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		apiURL:     apiURL,
		apiVersion: apiVersion,
		path:       "",
		httpClient: &http.Client{},
	}
}

// SetAPIVersion sets the version of the OpenF1 API to use (e.g. "v1")
func (c *Client) SetAPIVersion(version string) {
	c.apiVersion = version
}

// SetPath sets the specific API path (e.g. "/car_data") to be appended to the base URL
func (c *Client) SetPath(path string) {
	c.path = path
}

// SetHTTPClient allows user to provide their own http.Client
func (c *Client) SetHTTPClient(hc *http.Client) {
	c.httpClient = hc
}
