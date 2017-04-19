package types

import (
	"fmt"
)

// Coordinates code for coordinates used throughout the API
//
// This is the Go representation of "Coord". It implements Place.
//
// See http://doc.navitia.io/#standard-objects
type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// PlaceID formats & escapes the coordinates for use in queries as an ID
//
// Helps satisfy Place.
func (c Coordinates) PlaceID() string {
	return fmt.Sprintf("%3.3f;%3.3f", c.Longitude, c.Latitude)
}

// PlaceName returns a name for this coordinate.
//
// Helps satisfy Place.
func (c Coordinates) PlaceName() string {
	return c.PlaceID()
}

// PlaceType returns "coord".
//
// Helps satisfy Place.
func (c Coordinates) PlaceType() string {
	return "coord"
}

// String pretty-prints a Coordinates and satisfies the Stringer interface
func (c Coordinates) String() string {
	return fmt.Sprintf("Longitude %3.3f ; Latitude %3.3f", c.Longitude, c.Latitude)
}
