package types

import "time"

type CarData struct {
	Brake        int       `json:"brake"`
	Date         time.Time `json:"date"`
	DriverNumber int       `json:"driver_number"`
	Drs          int       `json:"drs"`
	MeetingKey   int       `json:"meeting_key"`
	NGear        int       `json:"n_gear"`
	Rpm          int       `json:"rpm"`
	SessionKey   int       `json:"session_key"`
	Speed        int       `json:"speed"`
	Throttle     int       `json:"throttle"`
}

type Driver struct {
	BroadcastName string `json:"broadcast_name"`
	CountryCode   string `json:"country_code"`
	DriverNumber  int    `json:"driver_number"`
	FirstName     string `json:"first_name"`
	FullName      string `json:"full_name"`
	HeadshotURL   string `json:"headshot_url"`
	LastName      string `json:"last_name"`
	MeetingKey    int    `json:"meeting_key"`
	NameAcronym   string `json:"name_acronym"`
	SessionKey    int    `json:"session_key"`
	TeamColour    string `json:"team_colour"`
	TeamName      string `json:"team_name"`
}

type Interval struct {
	Date         time.Time `json:"date"`
	DriverNumber int       `json:"driver_number"`
	GapToLeader  any       `json:"gap_to_leader"`
	Interval     float64   `json:"interval"`
	MeetingKey   int       `json:"meeting_key"`
	SessionKey   int       `json:"session_key"`
}

type Lap struct {
	DateStart       time.Time `json:"date_start"`
	DriverNumber    int       `json:"driver_number"`
	DurationSector1 float64   `json:"duration_sector_1"`
	DurationSector2 float64   `json:"duration_sector_2"`
	DurationSector3 float64   `json:"duration_sector_3"`
	I1Speed         int       `json:"i1_speed"`
	I2Speed         int       `json:"i2_speed"`
	IsPitOutLap     bool      `json:"is_pit_out_lap"`
	LapDuration     float64   `json:"lap_duration"`
	LapNumber       int       `json:"lap_number"`
	MeetingKey      int       `json:"meeting_key"`
	SegmentsSector1 []int     `json:"segments_sector_1"`
	SegmentsSector2 []int     `json:"segments_sector_2"`
	SegmentsSector3 []int     `json:"segments_sector_3"`
	SessionKey      int       `json:"session_key"`
	StSpeed         int       `json:"st_speed"`
}

type Location struct {
	Date         time.Time `json:"date"`
	DriverNumber int       `json:"driver_number"`
	MeetingKey   int       `json:"meeting_key"`
	SessionKey   int       `json:"session_key"`
	X            int       `json:"x"`
	Y            int       `json:"y"`
	Z            int       `json:"z"`
}

type Meeting struct {
	CircuitKey          int       `json:"circuit_key"`
	CircuitShortName    string    `json:"circuit_short_name"`
	CountryCode         string    `json:"country_code"`
	CountryKey          int       `json:"country_key"`
	CountryName         string    `json:"country_name"`
	DateStart           time.Time `json:"date_start"`
	GmtOffset           string    `json:"gmt_offset"`
	Location            string    `json:"location"`
	MeetingKey          int       `json:"meeting_key"`
	MeetingName         string    `json:"meeting_name"`
	MeetingOfficialName string    `json:"meeting_official_name"`
	Year                int       `json:"year"`
}

type Overtake struct {
	Date                   time.Time `json:"date"`
	MeetingKey             int       `json:"meeting_key"`
	OvertakenDriverNumber  int       `json:"overtaken_driver_number"`
	OvertakingDriverNumber int       `json:"overtaking_driver_number"`
	Position               int       `json:"position"`
	SessionKey             int       `json:"session_key"`
}

type Pit struct {
	Date         time.Time `json:"date"`
	DriverNumber int       `json:"driver_number"`
	LapNumber    int       `json:"lap_number"`
	MeetingKey   int       `json:"meeting_key"`
	PitDuration  float64   `json:"pit_duration"`
	SessionKey   int       `json:"session_key"`
}

type Position struct {
	Date         time.Time `json:"date"`
	DriverNumber int       `json:"driver_number"`
	MeetingKey   int       `json:"meeting_key"`
	Position     int       `json:"position"`
	SessionKey   int       `json:"session_key"`
}

type RaceControl struct {
	Category     string    `json:"category"`
	Date         time.Time `json:"date"`
	DriverNumber int       `json:"driver_number"`
	Flag         string    `json:"flag"`
	LapNumber    int       `json:"lap_number"`
	MeetingKey   int       `json:"meeting_key"`
	Message      string    `json:"message"`
	Scope        string    `json:"scope"`
	Sector       *int      `json:"sector"`
	SessionKey   int       `json:"session_key"`
}

type Session struct {
	CircuitKey       int       `json:"circuit_key"`
	CircuitShortName string    `json:"circuit_short_name"`
	CountryCode      string    `json:"country_code"`
	CountryKey       int       `json:"country_key"`
	CountryName      string    `json:"country_name"`
	DateEnd          time.Time `json:"date_end"`
	DateStart        time.Time `json:"date_start"`
	GmtOffset        string    `json:"gmt_offset"`
	Location         string    `json:"location"`
	MeetingKey       int       `json:"meeting_key"`
	SessionKey       int       `json:"session_key"`
	SessionName      string    `json:"session_name"`
	SessionType      string    `json:"session_type"`
	Year             int       `json:"year"`
}

type SessionResult struct {
	Dnf          bool    `json:"dnf"`
	DNS          bool    `json:"dns"`
	Dsq          bool    `json:"dsq"`
	DriverNumber int     `json:"driver_number"`
	Duration     float64 `json:"duration"`
	GapToLeader  int     `json:"gap_to_leader"`
	NumberOfLaps int     `json:"number_of_laps"`
	MeetingKey   int     `json:"meeting_key"`
	Position     int     `json:"position"`
	SessionKey   int     `json:"session_key"`
}

type StartingGrid struct {
	Position     int     `json:"position"`
	DriverNumber int     `json:"driver_number"`
	LapDuration  float64 `json:"lap_duration"`
	MeetingKey   int     `json:"meeting_key"`
	SessionKey   int     `json:"session_key"`
}

type Stint struct {
	Compound       string `json:"compound"`
	DriverNumber   int    `json:"driver_number"`
	LapEnd         int    `json:"lap_end"`
	LapStart       int    `json:"lap_start"`
	MeetingKey     int    `json:"meeting_key"`
	SessionKey     int    `json:"session_key"`
	StintNumber    int    `json:"stint_number"`
	TyreAgeAtStart int    `json:"tyre_age_at_start"`
}

type TeamRadio struct {
	Date         time.Time `json:"date"`
	DriverNumber int       `json:"driver_number"`
	MeetingKey   int       `json:"meeting_key"`
	RecordingURL string    `json:"recording_url"`
	SessionKey   int       `json:"session_key"`
}

type Weather struct {
	AirTemperature   float64   `json:"air_temperature"`
	Date             time.Time `json:"date"`
	Humidity         float64   `json:"humidity"`
	MeetingKey       int       `json:"meeting_key"`
	Pressure         float64   `json:"pressure"`
	Rainfall         int       `json:"rainfall"`
	SessionKey       int       `json:"session_key"`
	TrackTemperature float64   `json:"track_temperature"`
	WindDirection    int       `json:"wind_direction"`
	WindSpeed        float64   `json:"wind_speed"`
}

type ErrorMessage struct {
	Detail string `json:"detail"`
}
