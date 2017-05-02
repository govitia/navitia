# navitia is a Go client for the [navitia](navitia.io) API for public transit & transportation
[![Build Status](https://travis-ci.org/aabizri/navitia.svg?branch=dev)](https://travis-ci.org/aabizri/navitia) [![Go Report Card](https://goreportcard.com/badge/github.com/aabizri/navitia)](https://goreportcard.com/report/github.com/aabizri/navitia) [![GoDoc](https://godoc.org/github.com/aabizri/navitia?status.svg)](https://godoc.org/github.com/aabizri/navitia)

This is navitia -dev.

## Dependencies

- It needs at least go 1.7 to work as we use context & tests use testing.T.Run for subtests.
- The dependencies are directly pulled in by `go get`, but for you

## Install

`go get -u github.com/aabizri/navitia`

## Coverage

- Coverage [/coverage]: You can easily navigate through regions covered by navitia.io, with the coverage api. The shape of the region is provided in GeoJSON, though this is not yet implemented. [(navitia.io doc)](http://doc.navitia.io/#coverage)
- Journeys [/journeys]: This computes journeys or isochrone tables. [(navitia.io doc)](http://doc.navitia.io/#journeys)
- Places [/places]: Allows you to search in all geographical objects using their names, returning a list of places. [(navitia.io doc)](http://doc.navitia.io/#autocomplete-on-geographical-objects)
- Connections (Departures & Arrivals) [/departures,/arrivals]: Retrieve departures & arrivals for a specific ressource or place. [(navitia.io doc)](http://doc.navitia.io/#departures)

## Getting started

### Creating a new session

First, you should have an API key from navitia.io, if you don't already have one, it's [this way !](https://www.navitia.io/register/)
```golang
session, err := navitia.New(APIKEY)
```

### Finding places

```golang
// Create a request
req := navitia.PlacesRequest{
	Query: "10 rue du caire, Paris",
	Types: []string{"address"},
	Count: 1,
}

// Execute it
ctx := context.Background()
res, _ := session.Places(ctx, req)

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

// Create the scope
scope := session.Scope(regionID)

// Requests places in this scope
res, _ := scope.Places(context.Background(),req)
```

### Going further

Obviously, this is a very simple example of what navitia can do, [check out the documentation !](https://godoc.org/github.com/aabizri/navitia)

## What's new in dev ?

- Moved testdata loading & unmarshal testing to `navitia/testutils`, rewriting it in the process to be better

## Footnotes

I made this project as I wanted to explore and push my go skills, and I'm really up for you to contribute ! Send me a pull request and/or contact me if you have any questions! ( [@aabizri](https://twitter.com/aabizri) on twitter)
