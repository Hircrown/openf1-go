package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const teamRadioPath = "/team_radio"

// TeamRadio retrieves team radio data based on the given user filter.
// Only a limited selection of communications are included, not the complete record of radio interactions.
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) TeamRadio(filter types.TeamRadioFilter) ([]types.TeamRadio, error) {
	fullURL := createFullURL(filter, c, teamRadioPath)
	return doGet[types.TeamRadio](c.httpClient, fullURL)
}
