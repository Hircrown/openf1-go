package openf1

import (
	"strconv"
	"time"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const weatherPath = "/weather"

// Weather retrieves weather data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) Weather(filter types.WeatherFilter) ([]types.Weather, error) {
	fullURL := createFullURL(filter, c, weatherPath)
	return doGet[types.Weather](c.httpClient, fullURL)
}

func (c *Client) RainySessionsForYear(year int) ([]types.Session, error) {
	rainySessions, err := c.Weather(types.WeatherFilter{
		Rainfall: 1,
		Date:     ">2025-01-01",
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
		time.Sleep(350 * time.Millisecond)
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
