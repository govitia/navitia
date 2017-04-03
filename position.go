package gonavitia

import (
	"fmt"
)

type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func (coords Coordinates) formatURL() string {
	return fmt.Sprintf("%3.3f;%3.3f", coords.Longitude, coords.Latitude)
}
