package openf1

import (
	"fmt"
	"time"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const telemetryPath = "/car_data"

// Telemetry retrieves telemetry data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) Telemetry(filter types.CarDataFilter) ([]types.CarData, error) {
	fullURL := createFullURL(filter, c, telemetryPath)
	return doGet[types.CarData](c.httpClient, fullURL)
}

func (c *Client) DriverFastestLapTelemetry(sessionKey string, driverNumber int) ([]types.CarData, error) {
	fastestLap, err := c.DriverFastestLap(sessionKey, driverNumber)
	if err != nil {
		return nil, err
	}
	duration := time.Duration(fastestLap.LapDuration * float64(time.Second))
	lapStartTime := fastestLap.DateStart
	lapEndTime := lapStartTime.Add(duration)
	telemetry, err := c.Telemetry(types.CarDataFilter{
		SessionKey:   sessionKey,
		DriverNumber: driverNumber,
		Date:         fmt.Sprintf(">=%s&date<=%s", lapStartTime, lapEndTime),
	})
	if err != nil {
		return nil, err
	}
	return telemetry, nil
}

func (c *Client) TelemetryByLap(sessionKey, lap string, driverNumber int) ([]types.CarData, error) {
	lapData, err := c.Laps(types.LapFilter{
		SessionKey:   sessionKey,
		LapNumber:    lap,
		DriverNumber: driverNumber,
	})
	if err != nil {
		return nil, err
	}
	if len(lapData) == 0 {
		err := fmt.Errorf(
			"request successful but returned no data for the following query parameters:\n"+
				"  session key: %s\n"+
				"  driver number: %d\n"+
				"  lap number: %s",
			sessionKey,
			driverNumber,
			lap,
		)
		return nil, err
	}
	duration := time.Duration(lapData[0].LapDuration * float64(time.Second))
	lapStartTime := lapData[0].DateStart
	lapEndTime := lapStartTime.Add(duration)
	telemetry, err := c.Telemetry(types.CarDataFilter{
		SessionKey:   sessionKey,
		DriverNumber: driverNumber,
		Date:         fmt.Sprintf(">=%s&date<=%s", lapStartTime, lapEndTime),
	})
	if err != nil {
		return nil, err
	}
	return telemetry, nil
}
