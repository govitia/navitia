package types

import (
	"fmt"
)

type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func (coords Coordinates) FormatURL() (string, error) {
	return fmt.Sprintf("%3.3f;%3.3f", coords.Longitude, coords.Latitude), nil
}

func (coords Coordinates) String() string {
	return fmt.Sprintf("Longitude %3.3f ; Latitude %3.3f", coords.Longitude, coords.Latitude)
}
