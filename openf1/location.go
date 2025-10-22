package openf1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const locationPath = "/location"

// Location retrieves location data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) Location(filter types.LocationFilter) ([]types.Location, error) {
	fullURL := createFullURL(filter, c, locationPath)
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
		var locations []types.Location
		err = json.NewDecoder(res.Body).Decode(&locations)
		if err != nil {
			return nil, err
		}
		return locations, nil
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

func (c *Client) LocationByLap(session_key, lap string, driver_number int) ([]types.Location, error) {
	lapData, err := c.Laps(types.LapFilter{
		SessionKey:   session_key,
		LapNumber:    lap,
		DriverNumber: driver_number,
	})
	if err != nil {
		return nil, err
	}
	duration := time.Duration(lapData[0].LapDuration * float64(time.Second))
	lapStartTime := lapData[0].DateStart
	lapEndTime := lapStartTime.Add(duration)
	locations, err := c.Location(types.LocationFilter{
		SessionKey:   session_key,
		DriverNumber: driver_number,
		Date:         fmt.Sprintf(">=%s&date<=%s", lapStartTime, lapEndTime),
	})
	if err != nil {
		return nil, err
	}
	return locations, nil
}
