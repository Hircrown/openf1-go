package types

import "time"

// Use Filter to define query constraints.

type CarDataFilter struct {
	Brake        int       `url:"brake,omitempty"`
	Date         time.Time `url:"date,omitempty"`
	DriverNumber int       `url:"driver_number,omitempty"`
	Drs          int       `url:"drs,omitempty"`
	MeetingKey   string    `url:"meeting_key,omitempty"`
	NGear        int       `url:"n_gear,omitempty"`
	Rpm          int       `url:"rpm,omitempty"`
	SessionKey   string    `url:"session_key,omitempty"`
	Speed        string    `url:"speed,omitempty"`
	Throttle     int       `url:"throttle,omitempty"`
}

type DriverFilter struct {
	BroadcastName string `url:"broadcast_name,omitempty"`
	CountryCode   string `url:"country_code,omitempty"`
	DriverNumber  int    `url:"driver_number,omitempty"`
	FirstName     string `url:"first_name,omitempty"`
	FullName      string `url:"full_name,omitempty"`
	HeadshotURL   string `url:"headshot_url,omitempty"`
	LastName      string `url:"last_name,omitempty"`
	MeetingKey    int    `url:"meeting_key,omitempty"`
	NameAcronym   string `url:"name_acronym,omitempty"`
	SessionKey    int    `url:"session_key,omitempty"`
	TeamColour    string `url:"team_colour,omitempty"`
	TeamName      string `url:"team_name,omitempty"`
}

type IntervalFilter struct {
	Date         time.Time `url:"date,omitempty"`
	DriverNumber int       `url:"driver_number,omitempty"`
	GapToLeader  float64   `url:"gap_to_leader,omitempty"`
	Interval     float64   `url:"interval,omitempty"`
	MeetingKey   int       `url:"meeting_key,omitempty"`
	SessionKey   int       `url:"session_key,omitempty"`
}

type LapFilter struct {
	DateStart       time.Time `url:"date_start,omitempty"`
	DriverNumber    int       `url:"driver_number,omitempty"`
	DurationSector1 float64   `url:"duration_sector_1,omitempty"`
	DurationSector2 float64   `url:"duration_sector_2,omitempty"`
	DurationSector3 float64   `url:"duration_sector_3,omitempty"`
	I1Speed         int       `url:"i1_speed,omitempty"`
	I2Speed         int       `url:"i2_speed,omitempty"`
	IsPitOutLap     bool      `url:"is_pit_out_lap,omitempty"`
	LapDuration     float64   `url:"lap_duration,omitempty"`
	LapNumber       int       `url:"lap_number,omitempty"`
	MeetingKey      int       `url:"meeting_key,omitempty"`
	SessionKey      int       `url:"session_key,omitempty"`
	StSpeed         int       `url:"st_speed,omitempty"`
}

type LocationFilter struct {
	Date         time.Time `url:"date,omitempty"`
	DriverNumber int       `url:"driver_number,omitempty"`
	MeetingKey   int       `url:"meeting_key,omitempty"`
	SessionKey   int       `url:"session_key,omitempty"`
	X            int       `url:"x,omitempty"`
	Y            int       `url:"y,omitempty"`
	Z            int       `url:"z,omitempty"`
}

type MeetingFilter struct {
	CircuitKey          int       `url:"circuit_key,omitempty"`
	CircuitShortName    string    `url:"circuit_short_name,omitempty"`
	CountryCode         string    `url:"country_code,omitempty"`
	CountryKey          int       `url:"country_key,omitempty"`
	CountryName         string    `url:"country_name,omitempty"`
	DateStart           time.Time `url:"date_start,omitempty"`
	GmtOffset           string    `url:"gmt_offset,omitempty"`
	Location            string    `url:"location,omitempty"`
	MeetingKey          int       `url:"meeting_key,omitempty"`
	MeetingName         string    `url:"meeting_name,omitempty"`
	MeetingOfficialName string    `url:"meeting_official_name,omitempty"`
	Year                int       `url:"year,omitempty"`
}

type OvertakeFilter struct {
	Date                   time.Time `url:"date,omitempty"`
	MeetingKey             int       `url:"meeting_key,omitempty"`
	OvertakenDriverNumber  int       `url:"overtaken_driver_number,omitempty"`
	OvertakingDriverNumber int       `url:"overtaking_driver_number,omitempty"`
	Position               int       `url:"position,omitempty"`
	SessionKey             int       `url:"session_key,omitempty"`
}

type PitFilter struct {
	Date         time.Time `url:"date,omitempty"`
	DriverNumber int       `url:"driver_number,omitempty"`
	LapNumber    int       `url:"lap_number,omitempty"`
	MeetingKey   int       `url:"meeting_key,omitempty"`
	PitDuration  float64   `url:"pit_duration,omitempty"`
	SessionKey   int       `url:"session_key,omitempty"`
}

type PositionFilter struct {
	Date         time.Time `url:"date,omitempty"`
	DriverNumber int       `url:"driver_number,omitempty"`
	MeetingKey   int       `url:"meeting_key,omitempty"`
	Position     int       `url:"position,omitempty"`
	SessionKey   int       `url:"session_key,omitempty"`
}

type RaceControlFilter struct {
	Category     string      `url:"category,omitempty"`
	Date         time.Time   `url:"date,omitempty"`
	DriverNumber int         `url:"driver_number,omitempty"`
	Flag         string      `url:"flag,omitempty"`
	LapNumber    int         `url:"lap_number,omitempty"`
	MeetingKey   int         `url:"meeting_key,omitempty"`
	Message      string      `url:"message,omitempty"`
	Scope        string      `url:"scope,omitempty"`
	Sector       interface{} `url:"sector,omitempty"`
	SessionKey   int         `url:"session_key,omitempty"`
}

type SessionFilter struct {
	CircuitKey       int       `url:"circuit_key,omitempty"`
	CircuitShortName string    `url:"circuit_short_name,omitempty"`
	CountryCode      string    `url:"country_code,omitempty"`
	CountryKey       int       `url:"country_key,omitempty"`
	CountryName      string    `url:"country_name,omitempty"`
	DateEnd          time.Time `url:"date_end,omitempty"`
	DateStart        time.Time `url:"date_start,omitempty"`
	GmtOffset        string    `url:"gmt_offset,omitempty"`
	Location         string    `url:"location,omitempty"`
	MeetingKey       int       `url:"meeting_key,omitempty"`
	SessionKey       int       `url:"session_key,omitempty"`
	SessionName      string    `url:"session_name,omitempty"`
	SessionType      string    `url:"session_type,omitempty"`
	Year             int       `url:"year,omitempty"`
}

type SessionResultFilter struct {
	Dnf          bool    `url:"dnf,omitempty"`
	DNS          bool    `url:"dns,omitempty"`
	Dsq          bool    `url:"dsq,omitempty"`
	DriverNumber int     `url:"driver_number,omitempty"`
	Duration     float64 `url:"duration,omitempty"`
	GapToLeader  int     `url:"gap_to_leader,omitempty"`
	NumberOfLaps int     `url:"number_of_laps,omitempty"`
	MeetingKey   int     `url:"meeting_key,omitempty"`
	Position     int     `url:"position,omitempty"`
	SessionKey   int     `url:"session_key,omitempty"`
}

type StartingGridFilter struct {
	Position     int     `url:"position,omitempty"`
	DriverNumber int     `url:"driver_number,omitempty"`
	LapDuration  float64 `url:"lap_duration,omitempty"`
	MeetingKey   int     `url:"meeting_key,omitempty"`
	SessionKey   int     `url:"session_key,omitempty"`
}

type StintFilter struct {
	Compound       string `url:"compound,omitempty"`
	DriverNumber   int    `url:"driver_number,omitempty"`
	LapEnd         int    `url:"lap_end,omitempty"`
	LapStart       int    `url:"lap_start,omitempty"`
	MeetingKey     int    `url:"meeting_key,omitempty"`
	SessionKey     int    `url:"session_key,omitempty"`
	StintNumber    int    `url:"stint_number,omitempty"`
	TyreAgeAtStart int    `url:"tyre_age_at_start,omitempty"`
}

type TeamRadioFilter struct {
	Date         time.Time `url:"date,omitempty"`
	DriverNumber int       `url:"driver_number,omitempty"`
	MeetingKey   int       `url:"meeting_key,omitempty"`
	RecordingURL string    `url:"recording_url,omitempty"`
	SessionKey   int       `url:"session_key,omitempty"`
}

type WeatherFilter struct {
	AirTemperature   float64   `url:"air_temperature,omitempty"`
	Date             time.Time `url:"date,omitempty"`
	Humidity         int       `url:"humidity,omitempty"`
	MeetingKey       int       `url:"meeting_key,omitempty"`
	Pressure         float64   `url:"pressure,omitempty"`
	Rainfall         int       `url:"rainfall,omitempty"`
	SessionKey       int       `url:"session_key,omitempty"`
	TrackTemperature float64   `url:"track_temperature,omitempty"`
	WindDirection    int       `url:"wind_direction,omitempty"`
	WindSpeed        float64   `url:"wind_speed,omitempty"`
}
