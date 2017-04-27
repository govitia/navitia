package types

import "github.com/paulmach/go.geojson"

// An Isochrone is sent back by the /isochrones service, it gives you a multi-polygon geojson response which represent a same time travel zone.
//
// See https://en.wikipedia.org/wiki/Isochrone_map for what is an isochrone.
//
// See http://doc.navitia.io/#isochrones-currently-in-beta
type Isochrone geojson.Geometry
