# navitia is a Go client for the [navitia](navitia.io) API for public transit & transportation
[![Build Status](https://travis-ci.org/aabizri/navitia.svg?branch=master)](https://travis-ci.org/aabizri/navitia) [![Go Report Card](https://goreportcard.com/badge/github.com/govitia/navitia)](https://goreportcard.com/report/github.com/govitia/navitia) [![GoDoc](https://godoc.org/github.com/govitia/navitia?status.svg)](https://godoc.org/github.com/govitia/navitia)

This is govitia/navitia golang API.

## Dependencies

- It needs at least go 1.7 to work as we use context & tests use testing.T.Run for subtests.
- The dependencies are directly pulled in by `go get`, but for you

## Install

`go get -u github.com/govitia/navitia`

## Coverage

- Coverage [/coverage]: You can easily navigate through regions covered by navitia.io, with the coverage api. The shape of the region is provided in GeoJSON, though this is not yet implemented. [(navitia.io doc)](http://doc.navitia.io/#coverage)
- Journeys [/journeys]: This computes journeys or isochrone tables. [(navitia.io doc)](http://doc.navitia.io/#journeys)
- Places [/places]: Allows you to search in all geographical objects using their names, returning a list of places. [(navitia.io doc)](http://doc.navitia.io/#autocomplete-on-geographical-objects)

## Changelog
 
[Changelog](CHANGELOG.md)
 
## Getting started

### Creating a new session

First, you should have an API key from navitia.io, if you don't already have one, it's [this way !](https://www.navitia.io/register/)
```golang
session, err := navitia.New(APIKEY)
```

### Finding places

```golang
import(
	"github.com/govitia/navitia"
	"github.com/govitia/navitia/types"
	"context"
)

// Create a request
req := navitia.PlacesRequest{
	Query: "10 rue du caire, Paris",
	Types: []string{"address"},
	Count: 1,
}

// Execute it
res, _ := session.Places(context.Background(),req)

// Create a variable to store it
var myPlace types.Place

// Check if there are enough results, and then assign the first element as your place
if places := res.Places; len(places) != 0 {
	myPlace = res.Places
}
```
### Calculating a journey

```golang
import(
	"github.com/govitia/navitia"
	"fmt"
	"context"
)

// Create a request, having already found two corresponding places
request := navitia.JourneyRequest{
	From: myPlace,
	To: thatOtherPlace,
}

// Execute it
res, _ := session.Journeys(context.Background(),request)

// Print it (JourneysResults implements Stringer)
fmt.Println(res)
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
import (
	"github.com/govitia/navitia"
	"github.com/govitia/navitia/types"
	"context"
)

var (
	session *navitia.Session
	regionID types.ID
	req navitia.PlacesRequest
)

scope := session.Scope(regionID)

// Requests places in this scope
res, _ := scope.Places(context.Background(),req)
```

### Going further

Obviously, this is a very simple example of what navitia can do, [check out the documentation !](https://godoc.org/github.com/govitia/navitia)
