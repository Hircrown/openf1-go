package openf1

import (
	"fmt"
	"strings"

	"github.com/Hircrown/openf1-go/openf1/types"
)

const raceControlPath = "/position"

var validFlags = map[string]struct{}{
	"GREEN":           {},
	"CLEAR":           {},
	"YELLOW":          {},
	"DOUBLE YELLOW":   {},
	"RED":             {},
	"CHEQUERED":       {},
	"BLUE":            {},
	"BLACK AND WHITE": {},
}

func validateFlagColor(input string) (string, error) {
	flag := strings.ToUpper(strings.TrimSpace(input))

	if _, ok := validFlags[flag]; !ok {
		return "", fmt.Errorf("invalid flag color: %q", input)
	}
	return flag, nil
}

// RaceControl retrieves race control data based on the given user filter.
// Wrong filter parameter won't result in error but in an empty slice.
// Excessive data requests may result in an error.
func (c *Client) RaceControl(filter types.RaceControlFilter) ([]types.RaceControl, error) {
	fullURL := createFullURL(filter, c, raceControlPath)
	return doGet[types.RaceControl](c.httpClient, fullURL)
}

// Flag retrieves race control flag events for a specified flag color and session key.
func (c *Client) Flag(flagColor, sessionKey string) ([]types.RaceControl, error) {
	validatedFlag, err := validateFlagColor(flagColor)
	if err != nil {
		return nil, err
	}
	flags, err := c.RaceControl(types.RaceControlFilter{
		SessionKey: sessionKey,
		Flag:       validatedFlag,
	})
	if err != nil {
		return nil, err
	}
	return flags, nil
}

// SafetyCar retrieves all race control events related to the safety car for the specified session key.
// The message field specifies whether it is a virtual safety car or not.
func (c *Client) SafetyCar(sessionKey string) ([]types.RaceControl, error) {
	safetyCar, err := c.RaceControl(types.RaceControlFilter{
		SessionKey: sessionKey,
		Category:   "SafetyCar",
	})
	if err != nil {
		return nil, err
	}
	return safetyCar, nil
}
