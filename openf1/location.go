package openf1

import (
	"time"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const locationPath = "/location"

// Location retrieves The approximate location of the cars on the circuit data based on the given user filter,
// at a sample rate of about 3.7 Hz. Useful for gauging their progress along the track, but lacks details
// about lateral placement — i.e. whether the car is on the left or right side of the track.
// The origin point (0, 0, 0) appears to be arbitrary and not tied to any specific location on the track.
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) Location(filter types.LocationFilter) ([]types.Location, error) {
	fullURL := createFullURL(filter, c, locationPath)
	return doGet[types.Location](c.httpClient, fullURL)
}

func (c *Client) LocationByLap(sessionKey, lap string, driverNumber int) ([]types.Location, error) {
	lapData, err := c.Laps(types.LapFilter{
		SessionKey:   sessionKey,
		LapNumber:    lap,
		DriverNumber: driverNumber,
	})
	if err != nil {
		return nil, err
	}
	duration := time.Duration(lapData[0].LapDuration * float64(time.Second))
	lapStartTime := lapData[0].DateStart
	lapEndTime := lapStartTime.Add(duration)
	dateQuery, err := ValuesBetween(
		types.LocationFilter{}, "Date", lapStartTime.String(), lapEndTime.String(), true,
	)
	if err != nil {
		return nil, err
	}
	locations, err := c.Location(types.LocationFilter{
		SessionKey:   sessionKey,
		DriverNumber: driverNumber,
		Date:         dateQuery,
	})
	if err != nil {
		return nil, err
	}
	return locations, nil
}
