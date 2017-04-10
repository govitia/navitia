package types

import (
	"fmt"
)

// Coordinates code for coordinates
type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

// QueryEscape formats & escapes the coordinates for use in queries
func (coords Coordinates) QueryEscape() string {
	return fmt.Sprintf("%3.3f;%3.3f", coords.Longitude, coords.Latitude)
}

// String pretty-prints a Coordinates and satisifes the Stringer interface
func (coords Coordinates) String() string {
	return fmt.Sprintf("Longitude %3.3f ; Latitude %3.3f", coords.Longitude, coords.Latitude)
}
