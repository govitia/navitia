# navitia is a Go client for the [navitia](navitia.io) API for public transit & transportation
[![Build Status](https://travis-ci.org/aabizri/navitia.svg?branch=dev)](https://travis-ci.org/aabizri/navitia) [![Go Report Card](https://goreportcard.com/badge/github.com/aabizri/navitia)](https://goreportcard.com/report/github.com/aabizri/navitia) [![GoDoc](https://godoc.org/github.com/aabizri/navitia?status.svg)](https://godoc.org/github.com/aabizri/navitia)

This is navitia -dev.

## Dependencies

- It needs at least go 1.8 to work.
- The dependencies are directly pulled in by `go get`.

## Install

`go get -u github.com/aabizri/navitia`

## Coverage

- Coverage [/coverage]: You can easily navigate through regions covered by navitia.io, with the coverage api. The shape of the region is provided in GeoJSON, though this is not yet implemented. [(navitia.io doc)](http://doc.navitia.io/#coverage)
- Journeys [/journeys]: This computes journeys or isochrone tables. [(navitia.io doc)](http://doc.navitia.io/#journeys)
- Places [/places]: Allows you to search in all geographical objects using their names, returning a list of places. [(navitia.io doc)](http://doc.navitia.io/#autocomplete-on-geographical-objects)
- Connections (Departures & Arrivals) [/departures,/arrivals]: Retrieve departures & arrivals for a specific resource or place. [(navitia.io doc)](http://doc.navitia.io/#departures)
- Inverted Geocoding (Finding your address from your coordinates) [/coords]: Retrieve the address & associated region ID given some coordinates [(navitia.io doc)](http://doc.navitia.io/#inverted-geocoding)
- Public Transportation searching [/pt_objects]: Allows you to search in all public transportation objects using their names, returning a list of public transportation objects. [(navitia.io doc)](http://doc.navitia.io/#autocomplete-on-public-transport-objects)
- Public Transporation Objects exploration [/lines, /networks, etc.]: Allows you to explore public transportation objects in a given region, returning a list of these. This is highly untested and very probably broken for now, but that's why it's not on master. [(navitia.io doc)](http://doc.navitia.io/#public-transportation-objects-exploration)

## Getting started

### Creating a new session

First, you should have an API key from navitia.io, if you don't already have one, it's [this way !](https://www.navitia.io/register/)
```golang
session, err := navitia.New(APIKEY)
```

### Finding places

```golang
// Create a request for a single place, don't do that in the real world
opt := navitia.PlacesRequest{
	Types: []string{"address"},
	Count: 1,
}

// Execute it
ctx := context.Background()
res, _ := session.Places(ctx, "10 rue du caire", opt)

// Create a variable to store it
var myPlace types.Container

// Check if there are enough results, and then assign the first element as your place
if len(res.Places) != 0 {
	myPlace = res.Places[0]
}

// Print the name
fmt.Printf("Name: %s\n", myPlace.Name)
```
### Calculating a journey

```golang
var (
	myPlace types.Container
	thatOtherPlace types.Container
)

// Create a request, having already found two corresponding places
request := navitia.JourneyRequest{
	From: myPlace.ID,
	To:   thatOtherPlace.ID,
}

// Execute it
ctx := context.Background()
res, _ := session.Journeys(ctx, request)

// Print the duration of the first journey
fmt.Printf("Duration of journey #0: %s\n", res.Journeys[0].Duration.String())
```

### Paging

Unfortunately, paging isn't supported by Regions nor by Places requests. You'll have to play with the `PlacesRequest.Count` value in the latter case.
We'll use a Journey to showcase the paging:

```golang
// Obtain a journey like last time...

// Create a value to store the paginated result
var paginated *JourneyResults = res

// Iterate by checking if the Paging.Next function is not nil !
for paginated.Paging.Next != nil {
	// Create a new JourneyResults to hold the data
	p := JourneyResults{}
	
	// Call the function
	_ = paginated.Paging.Next(ctx, testSession, &p)

	// Assign a pointer to the previously created data structure to paginated
	paginated = &p
}
```
Obviously, you'll want to stop paginating at some point, and most importantly do something with the value.
An example is on the way !

### Scoping

When you wish to make some requests requiring a specific coverage, or have more meaningful results in global requests, you create a `Scope`

```golang
var (
	session *navitia.Session
	regionID types.ID
	req navitia.PlacesRequest
)

// Create the scope based on a regionID
scope := session.Scope(regionID)

// Requests places in this scope
res, _ := scope.Places(context.Background(),req)
```

Or you can also use coordinates for creating a scope that can then be used just like the previous way

```golang
var coords types.Coordinates

// Create the scope based on coordinates
scope := session.Scope(coords.ID())
```

### Going further

Obviously, this is a very simple example of what navitia can do, [check out the documentation !](https://godoc.org/github.com/aabizri/navitia)

## What's new in dev ?

- Moved testdata loading & unmarshal testing to `navitia/testutils`, rewriting it in the process to be better
- API-Break: Changed NewCustom and unexported APIURL and APIKey in Session.
- Request testing
- New /coords method (Coords) for finding out an address and associated region ID given coordinates.
- New exploration method (Explore) for listing all public transportation objects of a specific type in a region.
- New places nearby (PlacesNear) method for finding out the places nearby (duh).

## Roadmap

### Versions

Currently we are in the **v0** phase, which means that there is no garantees of _any_ kind of api stability. However this will change and hopefully soon:

**v1** will come, in the form of a new branch, once we have 100% _meaningful_ code coverage via good tests, plus complete fuzzed-enabled functionality. This will also mean API stability, obviously.

### These will come

_...but it may not come in this form_

- Add easy chaining of requests: most resource endpoints (e.g ../stop\_areas/{stop\_area}) have sub-endpoints (e.g ../routes) following HATEOS principle. Currently you have to take the resource's ID and call a different method. We could simplify this these ways:
	- For results type having multiple results in them:
		- Having a `func (r {Results}) Related() map[ID]map[string](func(context.Context,*Session) Results)` method on the results type, we'll have to create an exported interface `Results` for this to work intelligently.
		- Having a single function `func Related(r {Results}) map[ID]map[string](func(context.Context, *Session) Results)` much like the method.
	- For results type having a single result in them:
		- Having each `Results` type that returns only one result implement have an `Explore(ctx context.Context, s *Session, selector string)` method, allowing us easier chaining. (i.e `session.Coords(ctx, coords).Explore(ctx,session,"stop_areas")`).
	- And there may be other ways, I'm open to suggestion !
- Implement the following endpoints, which will allow us to achieve 100% coverage of the endpoints. (which doesn't mean 100% coverage of the functionality, see above)
	- Stop Schedules [/stop_schedules] [(navitia.io doc)](http://doc.navitia.io/#stop-schedules) (in the works)
	- Route Schedules [/route_schedules] [(navitia.io doc)](http://doc.navitia.io/#route-schedules) (in the works)
	- Traffic Reports [/traffic_reports] [(navitia.io doc)](http://doc.navitia.io/#traffic-reports) (`types` package already has support)
	- Contributors [/contributors] [(navitia.io doc)](http://doc.navitia.io/#contributors)
	- Datasets [/datasets] [(navitia.io doc)](http://doc.navitia.io/#datasets)
	- Isochrones [/isochrones] [(navitia.io doc)](http://doc.navitia.io/#isochrones-currently-in-beta)
- Move the providers info from URLs to a proper Struct, with information about which endpoints are allowed/disallowed, which will be integrated in `Session` via a `map[string]bool` (has O(1) complexity) allowing us not to waste bandwidth.
- Add proper handling of remote errors parsing failures by returning a `RemoteError` with the indication of what went wrong.
- In `types` subpackage:
	- Add validation to all types
	- Minor API-Break: Move `EmbeddedXXX` to `ResourceXXX`
	- Major API-Break: Rename every mention of `types` to `resources`, including the name of the subpackage.
- Tooling
	- Have a tool, either in go or in good old bash to download the testdata. (in the works)

### Open questions

- Have the SSL certificates fingerprints of the providers in the code, allowing us to double-check?
- For `navitia/types`, which geometry library should we use ?

## Notes for package users.

Even though this code is under UNLICENSE, this doesn't mean the binary you make by importing this code is under UNLICENSE, indeed, as we use other libraries (see Used packages) you **have to** follow their terms. However, good news for you, the terms are **permissive**, as all of the code we import is under BSD license, either BSD 2-Clause or 3-Clause.

This is not legal advice, but that's a rough and incomplete guide of what you should do if you redistribute either a binary or the source code with the dependencies vendored:
	- Read the licenses (go on, they are short)
	- Reproduce the licenses of the dependencies in the distribution somewhere (if you simply vendor them, the LICENSE should come included, but you should still check!).
_(for the stdlib, don't use the name of the copyright holders or of the contributors to endorse/promote the product)_

## Notes for potential contributors

- Note that the navitia.io API currently only has the Brittany region with disruptions enabled, so that's the reason we test it based on that dataset.
- I would love for you to contribute as well ! Notice that the code is under the UNLICENSE license, and by contributing you agree to put your code in that same license. However, if you would like to have the code in a less permissive license, like the MIT license, tell me and I'll probably move the code to that license.

## Used packages

Thanks to these fabulous makers:

- [pkg/errors](https://github.com/pkg/errors) by Dave Cheney et al. (BSD-2 Clause)
- [mb0/mkt](https://github.com/mb0/wkt) by Martin Schnabel (BSD-2 Clause)
- [twpayne/go-geom](https://github.com/twpayne/go-geom) by Tom Payne (BSD 2-Clause)
- And the amazing standard library and runtime (BSD 3-Clause)

## Footnotes

I made this project as I wanted to explore and push my go skills, and I'm really up for you to contribute ! Send me a pull request and/or contact me if you have any questions! ( [@aabizri](https://twitter.com/aabizri) on twitter)
