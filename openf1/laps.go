package openf1

import (
	"math"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const lapsPath = "/laps"

// Laps retrieves laps data based on the given user filter.
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) Laps(filter types.LapFilter) ([]types.Lap, error) {
	fullURL := createFullURL(filter, c, lapsPath)
	return doGet[types.Lap](c.httpClient, fullURL)

}

// DriverFastestLap retrieves the fastest lap of the given driver in the specified session.
func (c *Client) DriverFastestLap(sessionKey string, driverNumber int) (types.Lap, error) {
	laps, err := c.Laps(types.LapFilter{
		SessionKey:   sessionKey,
		DriverNumber: driverNumber,
	})
	if err != nil {
		return types.Lap{}, err
	}
	fastestLap := findFastestLap(laps)
	return fastestLap, nil
}

// SessionFastestLap retrieves the fastest lap in the specified session.
func (c *Client) SessionFastestLap(sessionKey string) (types.Lap, error) {
	driverNumbers, err := c.DriverNumbersBySession(sessionKey)
	if err != nil {
		return types.Lap{}, err
	}
	fastestLaps := make([]types.Lap, 0, 20)
	for _, num := range driverNumbers {
		lap, err := c.DriverFastestLap(sessionKey, num)
		if err != nil {
			return types.Lap{}, err
		}
		fastestLaps = append(fastestLaps, lap)
	}

	fastestSessionLap := findFastestLap(fastestLaps)
	return fastestSessionLap, nil
}

// findFastestLap is a helper function that returns the fastest lap from a slice of laps.
func findFastestLap(laps []types.Lap) types.Lap {
	var idx int
	fastest := math.Inf(1)

	for i, lap := range laps {
		if lap.LapDuration < fastest && lap.LapDuration != 0.0 {
			fastest = lap.LapDuration
			idx = i
		}
	}
	return laps[idx]
}
