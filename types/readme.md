# navitia-types is a library for working with types returned by the [Navitia](navitia.io) API. [![Build Status](https://travis-ci.org/aabizri/navitia-types.svg?branch=dev)](https://travis-ci.org/aabizri/navitia-types) [![GoDoc](https://godoc.org/github.com/aabizri/navitia-types?status.svg)](https://godoc.org/github.com/aabizri/navitia-types)

Package types implements support for the types used in the Navitia API (see doc.navitia.io), simplified and modified for idiomatic Go use.

This is navitia-types -dev. It is not API-Stable, and won't be until the v1 release, but it's getting closer !
This package was and is developped as a supporting library for the [navitia API client](https://github.com/aabizri/navitia) but can be used to build other navitia API clients.

## Install

Simply run `go get -u github.com/aabizri/navitia/types`.

## Coverage

Preview of the supported types, see [the doc](https://godoc.org/github.com/aabizri/navitia-types) for more information, and the [navitia.io doc](http://doc.navitia.io) for information about the remote API.

|Type Name|Description|Navitia Name|
|---|---|---|
|[`Journey`](https://godoc.org/github.com/aabizri/navitia-types#Journey)|A journey (X-->Y)|"journey"|
|[`Section`](https://godoc.org/github.com/aabizri/navitia-types#Section)|A section of a `Journey`|"section"|
|[`Region`](https://godoc.org/github.com/aabizri/navitia-types#Region)|A region covered by the API|"region"|
|[`Isochrone`](https://godoc.org/github.com/aabizri/navitia-types#Region)|A region covered by the API|"isochrone"|
|[`Container`](https://godoc.org/github.com/aabizri/navitia-types#Container)|This contains a Place or a PTObject|"place"/"pt_object"|
|[`Place`](https://godoc.org/github.com/aabizri/navitia-types#Place)|Place is an empty interface, by convention used to identify an `Address`, [`StopPoint`](https://godoc.org/github.com/aabizri/navitia-types#StopPoint), [`StopArea`](https://godoc.org/github.com/aabizri/navitia-types#StopArea), [`POI`](https://godoc.org/github.com/aabizri/navitia-types#POI), [`Admin`](https://godoc.org/github.com/aabizri/navitia-types#Admin) & [`Coordinates`](https://godoc.org/github.com/aabizri/navitia-types#Coordinates).|
|[`PTObject`](https://godoc.org/github.com/aabizri/navitia-types#Place)|PTObject is an empty interface by convention used to identify a Public Transportation object|
|[`Line`](https://godoc.org/github.com/aabizri/navitia-types#Line)|A public transit line.|"line"|
|[`Route`](https://godoc.org/github.com/aabizri/navitia-types#Route)|A specific route within a `Line`.|"route"|

And others, such as [`Display`](https://godoc.org/github.com/aabizri/navitia-types#Display) ["display_informations"], [`PTDateTime`](https://godoc.org/github.com/aabizri/navitia-types#PTDateTime) ["pt-date-time"], [`StopTime`](https://godoc.org/github.com/aabizri/navitia-types#StopTime) ["stop_time"]

## Getting started

```golang
import (
	"fmt"

	"github.com/aabizri/navitia-types"
)

func main() {
	data := []byte{"some journey's json"}
	var j types.Journey
	_ = j.UnmarshalJSON(data)
}
```

### Going further

Obviously, this is a very simple example of what navitia-types can do, [check out the documentation !](https://godoc.org/github.com/aabizri/navitia-types)

## What's new in -dev

- Let's see

## TODO

### Documentation

- Update `readme.md` to reflect new changes
- Add links to the doc.navitia.io documentation to every type

### Testing

- `(*PTDateTime).UnmarshalJSON`
- `ErrInvalidPlaceContainer.Error`
- `Equipment.Known`
- Every Type should have at least one file to be tested against

## Footnotes

I made this project as I wanted to explore and push my go skills, and I'm really up for you to contribute ! Send me a pull request and/or contact me if you have any questions! ( [@aabizri](https://twitter.com/aabizri) on twitter)
