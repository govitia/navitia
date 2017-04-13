# gonavitia/types is a library for dealing with types returned by the [Navitia](navitia.io) API. [![Build Status](https://travis-ci.org/aabizri/navitia.svg?branch=dev)](https://travis-ci.org/aabizri/navitia) [![GoDoc](https://godoc.org/github.com/aabizri/gonavitia/types?status.svg)](https://godoc.org/github.com/aabizri/gonavitia/types)

Package types implements support for the types used in the Navitia API (see doc.navitia.io), simplified and modified for idiomatic Go use.

This package is still a work in progress. It is not API-Stable, and won't be until the v1 release.

This package was and is developped as a supporting library for the [gonavitia API client](https://github.com/aabizri/gonavitia) but can be used to build other API clients.

## Install
Simply run `go get -u github.com/aabizri/gonavitia/types`.

## Coverage
This support includes or will include, for each type.
	- JSON Unmarshalling via UnmarshalJSON(b []byte), in the format of the navitia.io API
	- Validity Checking via Check()
	- Pretty-printing via String()

Currently supported types (with corresponding navitia type names in brackets, [see the navitia doc](doc.navitia.io))
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
