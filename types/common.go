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
