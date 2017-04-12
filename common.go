/*
Package types implements support for the types used in the Navitia API (see doc.navitia.io), simplified and modified for idiomatic Go use.

This package was and is developped as a supporting library for the gonavitia API client (https://github.com/aabizri/gonavitia) but can be used to build other API clients.

This support includes or will include, for each type.
	- JSON Unmarshalling via UnmarshalJSON(b []byte), in the format of the navitia.io API
	- Validity Checking via Check()
	- Pretty-printing via String()

This package is still a work in progress. It is not API-Stable, and won't be until the v1 release.

Currently supported types
	- Journey ["journey"]
	- Section ["section"]
	- Region ["region"]
	- Place (This is an interface for your ease-of-use, which is implemented by the five following types)
	- Address ["address"]
	- StopPoint ["stop_point"]
	- StopArea ["stop_area"]
	- AdministrativeRegion ["administrative_region"]
	- POI ["poi"]
	- Line ["line"]
	- Route ["route"]
	- And others, such as DisplayInformations ["display_informations"], PTDateTime ["pt-date-time"], StopTime ["stop_time"], Coordinates ["coord"].
*/
package types

import "time"

// DataFreshness codes for a specific data freshness requirement: realtime or base_schedule
type DataFreshness string

const (
	// DataFreshnessRealTime means you'll get undisrupted journeys
	DataFreshnessRealTime DataFreshness = "realtime"
	// DataFreshnessBaseSchedule means you can get disrupted journeys in the response.
	DataFreshnessBaseSchedule = "base_schedule"
)

// A QueryEscaper implements QueryEscape, which returns an escaped representation of the type for use in URL queries.
// Implemented by both ID and Coordinates
type QueryEscaper interface {
	QueryEscape() string
}

// A DisplayInformations hold informations useful to display
// Used by Section ["section"], RouteSchedule ["route_schedule"], StopSchedule ["stop_schedule"], Departure ["departure"], Arrival ["arrival"].
type DisplayInformations struct {
	// The headsign associated with the object
	Headsign string `json:"headsign"`

	// The name of the belonging network
	Network string `json:"network"`

	// A direction to take
	Direction string `json:"direction"`

	// The commercial mode in ID Form
	CommercialMode ID `json:"commercial_mode"`

	// The physical mode in ID Form
	PhysicalMode ID `json:"physical_mode"`

	// The label of the object
	Label string `json:"label"`

	// Hexadecimal color of the line
	Color string `json:"color"`

	// The text color for this section
	TextColor Color `json:"text_color"`

	// The code of the line
	Code string `json:"code"`

	// Description
	Description string `json:"description"`

	// Equipments on this object
	Equipments []Equipment `json:"equipments"`
}

// A PTDateTime (pt stands for “public transport”) is a complex date time object to manage the difference between stop and leaving times at a stop.
// It is used by:
// 	- Row in Schedule
// 	- StopSchedule
// 	- StopDatetime
type PTDateTime struct {
	// Date/Time of departure
	Departure time.Time

	// Date/Time of arrival
	Arrival time.Time
}
