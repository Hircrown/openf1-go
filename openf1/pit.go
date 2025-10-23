package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const pitPath = "/pit"

// Pits retrieves pits data based on the given user filter.
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) Pit(filter types.PitFilter) ([]types.Pit, error) {
	fullURL := createFullURL(filter, c, pitPath)
	return doGet[types.Pit](c.httpClient, fullURL)
}
