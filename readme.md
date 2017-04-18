# navitia/types is a library for dealing with types returned by the [Navitia](navitia.io) API. [![Build Status](https://travis-ci.org/aabizri/navitia-types.svg?branch=dev)](https://travis-ci.org/aabizri/navitia-types) [![GoDoc](https://godoc.org/github.com/aabizri/navitia-types?status.svg)](https://godoc.org/github.com/aabizri/navitia-types)

Package types implements support for the types used in the Navitia API (see doc.navitia.io), simplified and modified for idiomatic Go use.

This is version dev of the package. It is not completely API-Stable, and won't be until the v1 release.

This package was and is developped as a supporting library for the [navitia API client](https://github.com/aabizri/navitia) but can be used to build other navitia API clients.

## Install
Simply run `go get -u github.com/aabizri/navitia/types`.

## Coverage
This support includes or will include, for each type.
- JSON Unmarshalling via UnmarshalJSON(b []byte), in the format of the navitia.io API
- Validity Checking via Check()
- Pretty-printing via String()

Currently supported types (with corresponding navitia type names in brackets, [see the navitia doc](doc.navitia.io))
- Journey ["journey"]
- Section ["section"]
- Region ["region"]
- Isochrone ["isochrone"]
- Place (This is an interface for your ease-of-use, which is implemented by the five following types)
- Address ["address"]
- StopPoint ["stop_point"]
- StopArea ["stop_area"]
- AdministrativeRegion ["administrative_region"]
- POI ["poi"]
- PlaceContainer ["place"] (this is the official type returned by the navitia api)
- Line ["line"]
- Route ["route"]
- And others, such as DisplayInformations ["display_informations"], PTDateTime ["pt-date-time"], StopTime ["stop_time"], Coordinates ["coord"].
	
## What's new in dev
Let's see
