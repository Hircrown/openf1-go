package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const positionPath = "/position"

// Position retrieves position data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) Position(filter types.PositionFilter) ([]types.Position, error) {
	fullURL := createFullURL(filter, c, positionPath)
	return doGet[types.Position](c.httpClient, fullURL)
}
