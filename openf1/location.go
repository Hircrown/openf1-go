package openf1

import (
	"fmt"
	"time"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const locationPath = "/location"

// Location retrieves location data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
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
	locations, err := c.Location(types.LocationFilter{
		SessionKey:   sessionKey,
		DriverNumber: driverNumber,
		Date:         fmt.Sprintf(">=%s&date<=%s", lapStartTime, lapEndTime),
	})
	if err != nil {
		return nil, err
	}
	return locations, nil
}
