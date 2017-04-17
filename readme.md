# navitia is a Go client for the [navitia](navitia.io) API for public transit & transportation [![Build Status](https://travis-ci.org/aabizri/navitia.svg?branch=dev)](https://travis-ci.org/aabizri/navitia) [![GoDoc](https://godoc.org/github.com/aabizri/navitia?status.svg)](https://godoc.org/github.com/aabizri/navitia)

This is the development version of navitia.

It needs at least go 1.7 to work as tests use testing.T.Run for subtests.

## Install

`go get -u github.com/aabizri/navitia`

## Coverage

- Coverage [/coverage]: You can easily navigate through regions covered by navitia.io, with the coverage api. The shape of the region is provided in GeoJSON, though this is not yet implemented. [(navitia.io doc)](http://doc.navitia.io/#coverage)
- Journeys [/journeys]: This computes journeys or isochrone tables. [(navitia.io doc)](http://doc.navitia.io/#journeys)
- Places [/places]: Allows you to search in all geographical objects using their names, returning a list of places. [(navitia.io doc)](http://doc.navitia.io/#autocomplete-on-geographical-objects)

## Getting started

### Creating a new session

First, you should have an API key from navitia.io, if you don't already have one, it's [this way !](https://www.navitia.io/register/)
```golang
session, err := navitia.New(APIKEY)
```

### Finding places

```golang
import(
	"github.com/aabizri/navitia"
	"github.com/aabizri/navitia/types"
	"context"
)

// Create a request
req := JourneyRequest{
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
	"github.com/aabizri/navitia"
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

## Going further

Obviously, this is a very simple example of what navitia can do, [check out the documentation !](https://godoc.org/github.com/aabizri/navitia)

## Footnotes

I made this project as I wanted to explore and push my go skills, and I'm really up for you to contribute ! Send me a pull request and/or contact me if you have any questions! ( [@aabizri](https://twitter.com/aabizri) on twitter)
