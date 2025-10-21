package openf1

import (
	"strings"

	"github.com/google/go-querystring/query"
)

// cleanQuery fixes encoded query parameters where comparison operators like >= or <=
// are incorrectly encoded with an extra '=' (e.g., =>= becomes =%3E%3D).
// It replaces sequences like "=%3C" and "=%3E" with just "%3C" and "%3E"
// respectively to ensure correct query encoding.
func cleanQuery(query string) string {
	query = strings.ReplaceAll(query, "=%3C", "%3C")
	query = strings.ReplaceAll(query, "=%3E", "%3E")
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
