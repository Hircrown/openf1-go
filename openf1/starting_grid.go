package openf1

import (
	"strconv"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const startingGridPath = "/starting_grid"

// As reported in the openf1 api, this is a BETA endpoint
// StartingGrid retrieves the starting grid for the upcoming race based on user filter.
// Need to pass the qualifying session_key to get the race starting grid!
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) StartingGrid(filter types.StartingGridFilter) ([]types.StartingGrid, error) {
	fullURL := createFullURL(filter, c, startingGridPath)
	return doGet[types.StartingGrid](c.httpClient, fullURL)
}

// StartingGridGP returns the starting grid for a specific Grand Prix, including both sprint and classic races.
func (c *Client) StartingGridGP(cityName, year string, isSprintRace bool) ([]types.StartingGrid, error) {
	qualifyingType := "Qualifying"
	if isSprintRace {
		qualifyingType = "Sprint Qualifying"
	}
	sessionKey, err := c.GetSessionKey(cityName, "Qualifying", qualifyingType, year)
	if err != nil {
		return nil, err
	}
	grid, err := c.StartingGrid(types.StartingGridFilter{
		SessionKey: strconv.Itoa(sessionKey),
	})
	if err != nil {
		return nil, err
	}
	return grid, nil
}
