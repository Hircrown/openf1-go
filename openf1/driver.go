package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const driverPath = "/drivers"

// Driver retrieves driver informations based on the given user filter.
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) Driver(filter types.DriverFilter) ([]types.Driver, error) {
	fullURL := createFullURL(filter, c, driverPath)
	return doGet[types.Driver](c.httpClient, fullURL)
}

// DriverNumbersBySession returns a slice of driver numbers who participated in the specified session.
// Only the numeric driver identifiers are returned, not full driver details.
func (c *Client) DriverNumbersBySession(sessionKey string) ([]int, error) {
	laps, err := c.Laps(types.LapFilter{
		SessionKey: sessionKey,
		LapNumber:  "1",
	})
	if err != nil {
		return nil, err
	}
	driverNumbers := make([]int, 0, 20)
	for _, lap := range laps {
		driverNumbers = append(driverNumbers, lap.DriverNumber)
	}
	return driverNumbers, nil
}
