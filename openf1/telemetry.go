package openf1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const telemetryPath = "/car_data"

// Telemetry retrieves telemetry data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) Telemetry(filter types.CarDataFilter) ([]types.CarData, error) {
	fullURL := createFullURL(filter, c)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200:
		defer res.Body.Close()
		var telemetry []types.CarData
		err = json.NewDecoder(res.Body).Decode(&telemetry)
		if err != nil {
			return nil, err
		}
		return telemetry, nil
	case 422:
		fallthrough
	default:
		defer res.Body.Close()
		var errMsg types.ErrorMessage
		err := json.NewDecoder(res.Body).Decode(&errMsg)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(errMsg.Detail)
	}

}
