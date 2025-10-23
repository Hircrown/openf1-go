package openf1

import (
	"strconv"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const sessionResultPath = "/session_result"

// As reported in the openf1 api, this is a BETA endpoint
// SessionResult retrieves standings after a session based on the given user filter.
// An incorrect filter parameters resulting in no results will raise an error.
// Excessive data requests may result in an error.
func (c *Client) SessionResult(filter types.SessionResultFilter) ([]types.SessionResult, error) {
	fullURL := createFullURL(filter, c, sessionResultPath)
	return doGet[types.SessionResult](c.httpClient, fullURL)
}

func (c *Client) QualifyingResult(countryName, year string, isSprintRace bool) ([]types.SessionResult, error) {
	qualifyingType := "Qualifying"
	if isSprintRace {
		qualifyingType = "Sprint"
	}
	sessionKey, err := c.GetSessionKey(countryName, "Qualifying", qualifyingType, year)
	if err != nil {
		return nil, err
	}
	result, err := c.SessionResult(types.SessionResultFilter{
		SessionKey: strconv.Itoa(sessionKey),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) RaceResult(countryName, year string, isSprintRace bool) ([]types.SessionResult, error) {
	raceType := "Race"
	if isSprintRace {
		raceType = "Sprint"
	}
	sessionKey, err := c.GetSessionKey(countryName, "Race", raceType, year)
	if err != nil {
		return nil, err
	}
	result, err := c.SessionResult(types.SessionResultFilter{
		SessionKey: strconv.Itoa(sessionKey),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
