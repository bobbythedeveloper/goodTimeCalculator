package hello

import "time"

type sunRiseSunSet struct {
	Results struct {
		Sunrise                   string `json:"sunrise"`
		Sunset                    string `json:"sunset"`
		SolarNoon                 string `json:"solar_noon"`
		DayLength                 string `json:"day_length"`
		CivilTwilightBegin        string `json:"civil_twilight_begin"`
		CivilTwilightEnd          string `json:"civil_twilight_end"`
		NauticalTwilightBegin     string `json:"nautical_twilight_begin"`
		NauticalTwilightEnd       string `json:"nautical_twilight_end"`
		AstronomicalTwilightBegin string `json:"astronomical_twilight_begin"`
		AstronomicalTwilightEnd   string `json:"astronomical_twilight_end"`
	} `json:"results"`
	Status string `json:"status"`
}

type kaalamType struct {
	startTime time.Time
	endTime   time.Time
}

// ResponseJSON object to be sent over network. Newer API to serialize in protobuf
type ResponseJSON struct {
	RahuKaalStartTime time.Time `json:"rahuKaalStartTime"`
	RahuKaalEndTime   time.Time `json:"rahuKaalEndTime"`
}
