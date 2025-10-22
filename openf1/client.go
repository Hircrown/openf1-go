package openf1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Hircrown/openf1-go/openf1/types"
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

func doGet[T any](hc *http.Client, fullURL string) ([]T, error) {
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := hc.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200:
		defer res.Body.Close()
		var results []T
		err = json.NewDecoder(res.Body).Decode(&results)
		if err != nil {
			return nil, err
		}
		return results, nil
	case 422:
		fallthrough
	default:
		defer res.Body.Close()
		var errMsg types.ErrorMessage
		err := json.NewDecoder(res.Body).Decode(&errMsg)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(errMsg.Detail)
	}
}
