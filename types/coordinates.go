package types

import (
	"fmt"
)

// Coordinates code for coordinates
type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

// FormatURL formats the coordinates for use in queries
func (coords Coordinates) FormatURL() (string, error) {
	return fmt.Sprintf("%3.3f;%3.3f", coords.Longitude, coords.Latitude), nil
}

// String pretty-prints a Coordinates and satisifes the Stringer interface
func (coords Coordinates) String() string {
	return fmt.Sprintf("Longitude %3.3f ; Latitude %3.3f", coords.Longitude, coords.Latitude)
}
