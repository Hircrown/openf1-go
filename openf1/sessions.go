package openf1

import (
	"github.com/Hircrown/openf1-go/openf1/types"
)

const sessionsPath = "/sessions"

// Sessions retrieves sessions data based on the given user filter.
// A session refers to a distinct period of track activity during a Grand Prix
// or testing weekend (practice, qualifying, sprint, race, ...)
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) Sessions(filter types.SessionFilter) ([]types.Session, error) {
	fullURL := createFullURL(filter, c, sessionsPath)
	return doGet[types.Session](c.httpClient, fullURL)
}

// The sessionType-sessionName combinations vary depending on the Grand Prix format.
// Note: The query is case-sensitive.
//
// Classic GP Format       | Sprint GP Format
// ------------------------|--------------------------
// Practice - Practice 1   | Practice - Practice 1
// Practice - Practice 2   | Qualifying - Sprint
// Practice - Practice 3   | Race - Sprint
// Qualifying - Qualifying | Qualifying - Qualifying
// Race - Race             | Race - Race
func (c *Client) GetSessionKey(cityName, sessionType, sessionName, year string) (int, error) {
	sessions, err := c.Sessions(types.SessionFilter{
		Location:    capitalize(cityName),
		SessionType: capitalize(sessionType),
		SessionName: capitalize(sessionName),
		Year:        year,
	})
	if err != nil {
		return 0, err
	}
	return sessions[0].SessionKey, nil
}
