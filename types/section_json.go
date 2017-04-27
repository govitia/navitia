package types

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
)

/*
UnmarshalJSON implements json.Unmarshaller for a Section

Behaviour:
	- If "from" is empty, then don't populate the From field.
	- Same for "to"
*/
func (s *Section) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		// Pointers to the corresponding real values
		Type       *SectionType   `json:"type"`
		ID         *ID            `json:"id"`
		From       *Container     `json:"from"`
		To         *Container     `json:"to"`
		Mode       *string        `json:"mode"`
		StopTimes  *[]StopTime    `json:"stop_date_times"`
		Display    *Display       `json:"display_informations"`
		Additional *[]PTMethod    `json:"additional_informations"`
		Path       *[]PathSegment `json:"path"`

		// Values to process
		Departure string            `json:"departure_date_time"`
		Arrival   string            `json:"arrival_date_time"`
		Duration  int64             `json:"duration"`
		Geo       *geojson.Geometry `json:"geojson"`
	}{
		Type:       &s.Type,
		ID:         &s.ID,
		From:       &s.From,
		To:         &s.To,
		Mode:       &s.Mode,
		Display:    &s.Display,
		Additional: &s.Additional,
		StopTimes:  &s.StopTimes,
		Path:       &s.Path,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Create the error generator
	gen := unmarshalErrorMaker{"Section", b}

	// For departure and arrival, we use parseDateTime
	s.Departure, err = parseDateTime(data.Departure)
	if err != nil {
		return gen.err(err, "Departure", "departure_date_time", data.Departure, "parseDateTime failed")
	}
	s.Arrival, err = parseDateTime(data.Arrival)
	if err != nil {
		return gen.err(err, "Arrival", "arrival_date_time", data.Arrival, "parseDateTime failed")
	}

	// As the given duration is in second, let's multiply it by one second to have the correct value
	s.Duration = time.Duration(data.Duration) * time.Second

	// Now let's deal with the geom
	if data.Geo != nil {
		// Catch an error !
		if data.Geo.Coordinates == nil {
			return gen.err(nil, "Geo", "geojson", data.Geo, "Geo.Coordinates is nil, can't continue as that will cause a panic")
		}

		// Let's decode it
		geot, err := data.Geo.Decode()
		if err != nil {
			return gen.err(err, "Geo", "geojson", data.Geo, "Geo.Decode() failed")
		}
		// And let's assert the type
		geo, ok := geot.(*geom.LineString)
		if !ok {
			return gen.err(err, "Geo", "geojson", data.Geo, "Geo type assertion failed!")
		}
		// Now let's assign it
		s.Geo = geo
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaller for a PTDateTime
func (ptdt *PTDateTime) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	data := &struct {
		Departure string `json:"departure_date_time"`
		Arrival   string `json:"arrival_date_time"`
	}{}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Create the error generator
	gen := unmarshalErrorMaker{"PTDateTime", b}

	// Now we use parseDateTime
	ptdt.Departure, err = parseDateTime(data.Departure)
	if err != nil {
		return gen.err(err, "Departure", "departure_date_time", data.Departure, "parseDateTime failed")
	}
	ptdt.Arrival, err = parseDateTime(data.Arrival)
	if err != nil {
		return gen.err(err, "Arrival", "arrival_date_time", data.Arrival, "parseDateTime failed")
	}

	return nil
}
