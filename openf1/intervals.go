package openf1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const intervalsPath = "/intervals"

// Intervals retrieves intervals data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) Intervals(filter types.IntervalFilter) ([]types.Interval, error) {
	fullURL := createFullURL(filter, c, intervalsPath)
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
		var driver []types.Interval
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

// LappedDrivers returns information about lapped drivers for the given session key.
func (c *Client) LappedDrivers(session_key string) ([]types.Driver, error) {
	intervals, err := c.Intervals(types.IntervalFilter{
		SessionKey:  session_key,
		GapToLeader: "+1 LAP",
	})
	if err != nil {
		return nil, err
	}
	var drivers []types.Driver
	for _, interval := range intervals {
		driver, err := c.Driver(types.DriverFilter{
			SessionKey:   session_key,
			DriverNumber: interval.DriverNumber,
		})
		if err != nil {
			return nil, err
		}
		fmt.Println("driver len:", len(driver))
		drivers = append(drivers, driver[0])
	}
	return drivers, nil
}
