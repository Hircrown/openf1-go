package types

import "time"

// Use Filter to define query constraints.
// Most fields are defined as strings to allow flexible filtering.
// This enables passing special keywords like "latest" for MeetingKey and SessionKey,
// and using comparison operators (<, <=, >, >=) to filter data more dynamically.
// For more information about data filtering, visit:
// https://openf1.org/#data-filtering

type CarDataFilter struct {
	Brake        string `url:"brake,omitempty"`
	Date         string `url:"date,omitempty"`
	DriverNumber int    `url:"driver_number,omitempty"`
	Drs          string `url:"drs,omitempty"`
	MeetingKey   string `url:"meeting_key,omitempty"`
	NGear        string `url:"n_gear,omitempty"`
	Rpm          string `url:"rpm,omitempty"`
	SessionKey   string `url:"session_key,omitempty"`
	Speed        string `url:"speed,omitempty"`
	Throttle     string `url:"throttle,omitempty"`
}

type DriverFilter struct {
	BroadcastName string `url:"broadcast_name,omitempty"`
	CountryCode   string `url:"country_code,omitempty"`
	DriverNumber  int    `url:"driver_number,omitempty"`
	FirstName     string `url:"first_name,omitempty"`
	FullName      string `url:"full_name,omitempty"`
	HeadshotURL   string `url:"headshot_url,omitempty"`
	LastName      string `url:"last_name,omitempty"`
	MeetingKey    string `url:"meeting_key,omitempty"`
	NameAcronym   string `url:"name_acronym,omitempty"`
	SessionKey    string `url:"session_key,omitempty"`
	TeamColour    string `url:"team_colour,omitempty"`
	TeamName      string `url:"team_name,omitempty"`
}

type IntervalFilter struct {
	Date         time.Time `url:"date,omitempty"`
	DriverNumber int       `url:"driver_number,omitempty"`
	GapToLeader  string    `url:"gap_to_leader,omitempty"`
	Interval     string    `url:"interval,omitempty"`
	MeetingKey   string    `url:"meeting_key,omitempty"`
	SessionKey   string    `url:"session_key,omitempty"`
}

type LapFilter struct {
	DateStart       string `url:"date_start,omitempty"`
	DriverNumber    int    `url:"driver_number,omitempty"`
	DurationSector1 string `url:"duration_sector_1,omitempty"`
	DurationSector2 string `url:"duration_sector_2,omitempty"`
	DurationSector3 string `url:"duration_sector_3,omitempty"`
	I1Speed         string `url:"i1_speed,omitempty"`
	I2Speed         string `url:"i2_speed,omitempty"`
	IsPitOutLap     bool   `url:"is_pit_out_lap,omitempty"`
	LapDuration     string `url:"lap_duration,omitempty"`
	LapNumber       string `url:"lap_number,omitempty"`
	MeetingKey      string `url:"meeting_key,omitempty"`
	SessionKey      string `url:"session_key,omitempty"`
	StSpeed         string `url:"st_speed,omitempty"`
}

type LocationFilter struct {
	Date         string `url:"date,omitempty"`
	DriverNumber int    `url:"driver_number,omitempty"`
	MeetingKey   string `url:"meeting_key,omitempty"`
	SessionKey   string `url:"session_key,omitempty"`
	X            int    `url:"x,omitempty"`
	Y            int    `url:"y,omitempty"`
	Z            int    `url:"z,omitempty"`
}

type MeetingFilter struct {
	CircuitKey          int    `url:"circuit_key,omitempty"`
	CircuitShortName    string `url:"circuit_short_name,omitempty"`
	CountryCode         string `url:"country_code,omitempty"`
	CountryKey          int    `url:"country_key,omitempty"`
	CountryName         string `url:"country_name,omitempty"`
	DateStart           string `url:"date_start,omitempty"`
	GmtOffset           string `url:"gmt_offset,omitempty"`
	Location            string `url:"location,omitempty"`
	MeetingKey          string `url:"meeting_key,omitempty"`
	MeetingName         string `url:"meeting_name,omitempty"`
	MeetingOfficialName string `url:"meeting_official_name,omitempty"`
	Year                string `url:"year,omitempty"`
}

type OvertakeFilter struct {
	Date                   string `url:"date,omitempty"`
	MeetingKey             string `url:"meeting_key,omitempty"`
	OvertakenDriverNumber  int    `url:"overtaken_driver_number,omitempty"`
	OvertakingDriverNumber int    `url:"overtaking_driver_number,omitempty"`
	Position               int    `url:"position,omitempty"`
	SessionKey             string `url:"session_key,omitempty"`
}

type PitFilter struct {
	Date         string `url:"date,omitempty"`
	DriverNumber int    `url:"driver_number,omitempty"`
	LapNumber    string `url:"lap_number,omitempty"`
	MeetingKey   string `url:"meeting_key,omitempty"`
	PitDuration  string `url:"pit_duration,omitempty"`
	SessionKey   string `url:"session_key,omitempty"`
}

type PositionFilter struct {
	Date         string `url:"date,omitempty"`
	DriverNumber int    `url:"driver_number,omitempty"`
	MeetingKey   string `url:"meeting_key,omitempty"`
	Position     string `url:"position,omitempty"`
	SessionKey   string `url:"session_key,omitempty"`
}

type RaceControlFilter struct {
	Category     string `url:"category,omitempty"`
	Date         string `url:"date,omitempty"`
	DriverNumber int    `url:"driver_number,omitempty"`
	Flag         string `url:"flag,omitempty"`
	LapNumber    string `url:"lap_number,omitempty"`
	MeetingKey   string `url:"meeting_key,omitempty"`
	Message      string `url:"message,omitempty"`
	Scope        string `url:"scope,omitempty"`
	Sector       *int   `url:"sector,omitempty"`
	SessionKey   string `url:"session_key,omitempty"`
}

type SessionFilter struct {
	CircuitKey       int    `url:"circuit_key,omitempty"`
	CircuitShortName string `url:"circuit_short_name,omitempty"`
	CountryCode      string `url:"country_code,omitempty"`
	CountryKey       int    `url:"country_key,omitempty"`
	CountryName      string `url:"country_name,omitempty"`
	DateEnd          string `url:"date_end,omitempty"`
	DateStart        string `url:"date_start,omitempty"`
	GmtOffset        string `url:"gmt_offset,omitempty"`
	Location         string `url:"location,omitempty"`
	MeetingKey       string `url:"meeting_key,omitempty"`
	SessionKey       string `url:"session_key,omitempty"`
	SessionName      string `url:"session_name,omitempty"`
	SessionType      string `url:"session_type,omitempty"`
	Year             string `url:"year,omitempty"`
}

type SessionResultFilter struct {
	Dnf          bool    `url:"dnf,omitempty"`
	DNS          bool    `url:"dns,omitempty"`
	Dsq          bool    `url:"dsq,omitempty"`
	DriverNumber int     `url:"driver_number,omitempty"`
	Duration     float64 `url:"duration,omitempty"`
	GapToLeader  string  `url:"gap_to_leader,omitempty"`
	NumberOfLaps int     `url:"number_of_laps,omitempty"`
	MeetingKey   string  `url:"meeting_key,omitempty"`
	Position     string  `url:"position,omitempty"`
	SessionKey   string  `url:"session_key,omitempty"`
}

type StartingGridFilter struct {
	Position     string  `url:"position,omitempty"`
	DriverNumber int     `url:"driver_number,omitempty"`
	LapDuration  float64 `url:"lap_duration,omitempty"`
	MeetingKey   string  `url:"meeting_key,omitempty"`
	SessionKey   string  `url:"session_key,omitempty"`
}

type StintFilter struct {
	Compound       string `url:"compound,omitempty"`
	DriverNumber   int    `url:"driver_number,omitempty"`
	LapEnd         string `url:"lap_end,omitempty"`
	LapStart       string `url:"lap_start,omitempty"`
	MeetingKey     string `url:"meeting_key,omitempty"`
	SessionKey     string `url:"session_key,omitempty"`
	StintNumber    string `url:"stint_number,omitempty"`
	TyreAgeAtStart string `url:"tyre_age_at_start,omitempty"`
}

type TeamRadioFilter struct {
	Date         string `url:"date,omitempty"`
	DriverNumber int    `url:"driver_number,omitempty"`
	MeetingKey   string `url:"meeting_key,omitempty"`
	RecordingURL string `url:"recording_url,omitempty"`
	SessionKey   string `url:"session_key,omitempty"`
}

type WeatherFilter struct {
	AirTemperature   string `url:"air_temperature,omitempty"`
	Date             string `url:"date,omitempty"`
	Humidity         string `url:"humidity,omitempty"`
	MeetingKey       string `url:"meeting_key,omitempty"`
	Pressure         string `url:"pressure,omitempty"`
	Rainfall         string `url:"rainfall,omitempty"`
	SessionKey       string `url:"session_key,omitempty"`
	TrackTemperature string `url:"track_temperature,omitempty"`
	WindDirection    string `url:"wind_direction,omitempty"`
	WindSpeed        string `url:"wind_speed,omitempty"`
}
