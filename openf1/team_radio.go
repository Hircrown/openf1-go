package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const teamRadioPath = "/team_radio"

// TeamRadio retrieves team radio data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) TeamRadio(filter types.TeamRadioFilter) ([]types.TeamRadio, error) {
	fullURL := createFullURL(filter, c, teamRadioPath)
	return doGet[types.TeamRadio](c.httpClient, fullURL)
}
