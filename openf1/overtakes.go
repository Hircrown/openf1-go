package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const overtakesPath = "/overtakes"

// As reported in the openf1 api, this is a BETA endpoint.
// This data is only available during races and may be incomplete.
// Overtakes retrieves overtakes data based on the given user filter.
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) Overtakes(filter types.OvertakeFilter) ([]types.Overtake, error) {
	fullURL := createFullURL(filter, c, overtakesPath)
	return doGet[types.Overtake](c.httpClient, fullURL)
}
