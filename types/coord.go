package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Coord code for coordinates used throughout the API.
// This is the Go representation of "Coord". It implements Place.
// See http://doc.navitia.io/#standard-objects.
type Coord struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// jsonCoord define the JSON implementation of Coord struct
type jsonCoord struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
}

// ID formats coordinates for use in queries as an ID.
func (c Coord) ID() ID {
	return ID(fmt.Sprintf("%3.3f;%3.3f", c.Longitude, c.Latitude))
}

// UnmarshalJSON implements json.Unmarshaller for a Coord
func (c *Coord) UnmarshalJSON(b []byte) error {
	var data jsonCoord

	err := json.Unmarshal(b, &data)
	if err != nil {
		return fmt.Errorf("error while unmarshalling Coord struct : %w", err)
	}

	// Create the error generator
	gen := unmarshalErrorMaker{"Coord", b}

	// Now parse the values
	c.Longitude, err = strconv.ParseFloat(data.Longitude, 64)
	if err != nil {
		return gen.err(err, "Longitude", "lon", data.Longitude, "error in strconv.ParseFloat")
	}
	c.Latitude, err = strconv.ParseFloat(data.Latitude, 64)
	if err != nil {
		return gen.err(err, "Latitude", "lat", data.Latitude, "error in strconv.ParseFloat")
	}

	return nil
}