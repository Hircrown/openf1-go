package openf1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const driverPath = "/drivers"

// Driver retrieves driver data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) Driver(filter types.DriverFilter) ([]types.Driver, error) {
	fullURL := createFullURL(filter, c, driverPath)
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
		var driver []types.Driver
		err = json.NewDecoder(res.Body).Decode(&driver)
		if err != nil {
			return nil, err
		}
		return driver, nil
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
