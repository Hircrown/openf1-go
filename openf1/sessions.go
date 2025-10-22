package openf1

import (
	"fmt"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const sessionsPath = "/sessions"

// Sessions retrieves sessions data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
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
func (c *Client) GetSessionKey(countryName, sessionType, sessionName, year string) (int, error) {
	sessions, err := c.Sessions(types.SessionFilter{
		CountryName: countryName,
		SessionType: sessionType,
		SessionName: sessionName,
		Year:        year,
	})
	if err != nil {
		return 0, err
	}
	if len(sessions) == 0 {
		err := fmt.Errorf(
			"request successful but returned no data for:\n"+
				"  country: %s\n"+
				"  session type: %s\n"+
				"  session name: %s\n"+
				"  year: %s",
			countryName,
			sessionType,
			sessionName,
			year,
		)
		return 0, err
	}
	return sessions[0].SessionKey, nil
}
