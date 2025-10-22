package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const driverPath = "/drivers"

// Driver retrieves driver data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) Driver(filter types.DriverFilter) ([]types.Driver, error) {
	fullURL := createFullURL(filter, c, driverPath)
	return doGet[types.Driver](c.httpClient, fullURL)
}

// DriverNumbersBySession returns a slice of driver numbers who participated in the specified session.
// Only the numeric driver identifiers are returned, not full driver details.
func (c *Client) DriverNumbersBySession(session_key string) ([]int, error) {
	laps, err := c.Laps(types.LapFilter{
		SessionKey: session_key,
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
