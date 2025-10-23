package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const meetingsPath = "/meetings"

// Meetings retrieves meetings data based on the given user filter.
// A meeting refers to a Grand Prix or testing weekend and usually includes multiple
// sessions (practice, qualifying, race, ...).
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) Meetings(filter types.MeetingFilter) ([]types.Meeting, error) {
	fullURL := createFullURL(filter, c, meetingsPath)
	return doGet[types.Meeting](c.httpClient, fullURL)
}
