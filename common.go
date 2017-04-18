/*
Package types implements support for the types used in the Navitia API (see doc.navitia.io), simplified and modified for idiomatic Go use.

This package was and is developed as a supporting library for the gonavitia API client (https://github.com/aabizri/gonavitia) but can be used to build other API clients.

This support includes or will include, for each type.
	- JSON Unmarshalling via UnmarshalJSON(b []byte), in the format of the navitia.io API
	- Validity Checking via Check()
	- Pretty-printing via String()

This package is still a work in progress. It is not API-Stable, and won't be until the v1 release.

Currently supported types
	- Journey ["journey"]
	- Section ["section"]
	- Region ["region"]
	- Isochrone ["isochrone"]
	- Place (This is an interface for your ease-of-use, which is implemented by the five following types)
	- Address ["address"]
	- StopPoint ["stop_point"]
	- StopArea ["stop_area"]
	- Admin ["administrative_region"]
	- POI ["poi"]
	- PlaceContainer ["place"] (this is the official type returned by the navitia api)
	- Line ["line"]
	- Route ["route"]
	- And others, such as Display ["display_informations"], PTDateTime ["pt-date-time"], StopTime ["stop_time"], Coordinates ["coord"].
*/
package types

// Version is the version of this package
const Version = "dev"

// A QueryEscaper implements QueryEscape, which returns an escaped representation of the type for use in URL queries.
// Implemented by both ID and Coordinates
type QueryEscaper interface {
	QueryEscape() string
}
