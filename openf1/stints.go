package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const stintsPath = "/stints"

// Stints retrieves stints data based on the given user filter.
// A stint refers to a period of continuous driving by a driver during a session.
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) Stints(filter types.StintFilter) ([]types.Stint, error) {
	fullURL := createFullURL(filter, c, stintsPath)
	return doGet[types.Stint](c.httpClient, fullURL)
}

// DriverStints returns a list of stints for the specified driver in the given session.
func (c *Client) DriverStints(sessionKey string, driverNumber int) ([]types.Stint, error) {
	stints, err := c.Stints(types.StintFilter{
		SessionKey:   sessionKey,
		DriverNumber: driverNumber,
	})
	if err != nil {
		return nil, err
	}
	return stints, nil
}

// StartingCompound returns the starting tyre compound for each driver in the grid
// for the specified session. By checking TyreAgeAtStart, you can identify drivers
// who started with fresh tyres.
func (c *Client) StartingCompound(sessionKey string) ([]types.Stint, error) {
	stints, err := c.Stints(types.StintFilter{
		SessionKey: sessionKey,
		LapStart:   "1",
	})
	if err != nil {
		return nil, err
	}
	return stints, nil
}
