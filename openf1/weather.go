package openf1

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const weatherPath = "/weather"

// Weather retrieves weather data based on the given user filter.
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) Weather(filter types.WeatherFilter) ([]types.Weather, error) {
	fullURL := createFullURL(filter, c, weatherPath)
	return doGet[types.Weather](c.httpClient, fullURL)
}

// RainySessionsForYear retrieves all the sessions that experienced rain during the specified year.
func (c *Client) RainySessionsForYear(year int) ([]types.Session, error) {
	rainySessions, err := c.Weather(types.WeatherFilter{
		Rainfall: 1,
		Date:     fmt.Sprintf(">%d-01-01&date<%d-12-31", year, year),
	})
	if err != nil {
		return nil, err
	}

	var uniqueSessionKey = map[int]struct{}{}
	for _, session := range rainySessions {
		if _, ok := uniqueSessionKey[session.SessionKey]; ok {
			continue
		}
		uniqueSessionKey[session.SessionKey] = struct{}{}
	}
	var sessions []types.Session
	for k := range uniqueSessionKey {
		// The api has a rate limit of max 3 request per second
		time.Sleep(300 * time.Millisecond)
		session, err := c.Sessions(types.SessionFilter{
			SessionKey: strconv.Itoa(k),
		})
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session[0])
	}
	return sessions, nil
}
