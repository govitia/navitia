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

// ID formats & escapes the coordinates for use in queries as an ID
func (c Coordinates) ID() ID {
	return ID(fmt.Sprintf("%3.3f;%3.3f", c.Longitude, c.Latitude))
}

// String pretty-prints a Coordinates and satisfies the Stringer interface
func (c Coordinates) String() string {
	return fmt.Sprintf("Longitude %3.3f ; Latitude %3.3f", c.Longitude, c.Latitude)
}
