package openf1

import (
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
