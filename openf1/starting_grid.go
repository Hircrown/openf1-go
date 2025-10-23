package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const startingGridPath = "/session_result"

// As reported in the openf1 api, this is a BETA endpoint
// StartingGrid retrieves starting grid data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) StartingGrid(filter types.StartingGridFilter) ([]types.StartingGrid, error) {
	fullURL := createFullURL(filter, c, startingGridPath)
	return doGet[types.StartingGrid](c.httpClient, fullURL)
}
