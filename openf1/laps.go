package openf1

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const lapsPath = "/laps"

// Laps retrieves laps data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) Laps(filter types.LapFilter) ([]types.Lap, error) {
	fullURL := createFullURL(filter, c, lapsPath)
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
		var laps []types.Lap
		err = json.NewDecoder(res.Body).Decode(&laps)
		if err != nil {
			return nil, err
		}
		return laps, nil
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

// DriverFastestLap retrieves the fastest lap of the given driver in the specified session.
func (c *Client) DriverFastestLap(session_key string, driver_number int) (types.Lap, error) {
	laps, err := c.Laps(types.LapFilter{
		SessionKey:   session_key,
		DriverNumber: driver_number,
	})
	fmt.Println(len(laps))
	if err != nil {
		return types.Lap{}, err
	}
	fastestLap := findFastestLap(laps)
	return fastestLap, nil
}

// SessionFastestLap retrieves the fastest lap in the specified session.
func (c *Client) SessionFastestLap(session_key string) (types.Lap, error) {
	driverNumbers, err := c.DriverNumbersBySession(session_key)
	if err != nil {
		return types.Lap{}, err
	}
	fastestLaps := make([]types.Lap, 0, 20)
	for _, num := range driverNumbers {
		lap, err := c.DriverFastestLap(session_key, num)
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
