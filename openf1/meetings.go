package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const meetingsPath = "/meetings"

// Meetings retrieves meetings data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) Meetings(filter types.MeetingFilter) ([]types.Meeting, error) {
	fullURL := createFullURL(filter, c, meetingsPath)
	return doGet[types.Meeting](c.httpClient, fullURL)
}
