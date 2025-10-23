package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const positionPath = "/position"

// Position retrieves driver positions throughout a session, including initial
// placement and subsequent changes based on the given user filter.
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) Position(filter types.PositionFilter) ([]types.Position, error) {
	fullURL := createFullURL(filter, c, positionPath)
	return doGet[types.Position](c.httpClient, fullURL)
}
