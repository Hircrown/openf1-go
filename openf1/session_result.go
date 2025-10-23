package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const sessionResultPath = "/session_result"

// As reported in the openf1 api, this is a BETA endpoint
// SessionResult retrieves standings after a session based on the given user filter.
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) SessionResult(filter types.SessionResultFilter) ([]types.SessionResult, error) {
	fullURL := createFullURL(filter, c, sessionResultPath)
	return doGet[types.SessionResult](c.httpClient, fullURL)
}
