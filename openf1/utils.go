package openf1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Hircrown/openf1-go/openf1/types"
	"github.com/google/go-querystring/query"
)

// cleanQuery fixes encoded query parameters where comparison operators like >= or <=
// are incorrectly encoded with an extra '=' (e.g., =>= becomes =%3E%3D).
// It replaces sequences like "=%3C" and "=%3E" with just "%3C" and "%3E"
// respectively to ensure correct query encoding.
func cleanQuery(query string) string {
	query = strings.ReplaceAll(query, "=%3C", "%3C")
	query = strings.ReplaceAll(query, "=%3E", "%3E")
	query = strings.ReplaceAll(query, "%26", "&")
	return query
}

// createFullURL generates a full request URL from the given filter struct.
func createFullURL[T any](filter T, c *Client, path string) string {
	v, _ := query.Values(filter)
	query := v.Encode()
	query = cleanQuery(query)
	c.SetPath(path + "?" + query)
	fullURL := c.apiURL + c.apiVersion + c.path
	return fullURL
}

// doGet is a helper function that manages an HTTP GET request using the provided http.Client and URL.
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
		if len(results) == 0 {
			return nil, errors.New("could not find any data for the query: " + fullURL)
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
